package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Login user handler
/*
-> username string - name of user
*/
func (h *handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get role of user
	_, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(fmt.Errorf("invalid token: %s", err))
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// If not Admin
	if role != "Admin" {
		log.Println(errors.New(" attempt to get user list: no valid token"))
		fmt.Println("")
		http.Error(w, "Internal Server Error", http.StatusForbidden)
		return
	}

	// Get num of page
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Bad request", 400)
	}

	// Get rows count in ine page
	per_page, err := strconv.Atoi(r.FormValue("per_page"))
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Bad request", 400)
	}

	fmt.Println("POST user-list: part ", page)

	// How much to skip
	offset := per_page * (page - 1)

	// Get rows with users data
	users, err := h.userService.GetAll(offset, per_page)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}

	last_page := page
	if len(users) == per_page {
		last_page += 1
	}

	// Write user data to response body
	data := map[string]interface{}{
		"total":        offset + len(users),
		"per_page":     per_page,
		"current_page": page,
		"last_page":    page + 1,
		"from":         offset,
		"to":           offset + len(users),
		"data":         users,
	}

	json.NewEncoder(w).Encode(data)
}