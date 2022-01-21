package order

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) Complite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	if role != "Worker" {
		log.Println(fmt.Errorf("fatal attempt to create order without permission by %s", username))
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	uuid := r.FormValue("order")

	order, err := h.orderService.Get(uuid)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if username != order.Worker {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if err := h.orderService.Complite(uuid); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
