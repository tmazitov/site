package order

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type updateParams struct {
	UUID    string
	Title   string
	Comment string
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get username from access token
	username, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if role != "Manager" && role != "Admin" {
		log.Println(fmt.Errorf("fatal attempt to update order without permission by %s", username))
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var params updateParams
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 102400))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if err = json.Unmarshal(body, &params); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	order, err := h.orderService.Get(params.UUID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if username != order.Writer {
		log.Println(err.Error())
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	newData := map[string]string{
		"Title":   params.Title,
		"Comment": params.Comment,
	}

	if err := h.orderService.Update(params.UUID, newData); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
