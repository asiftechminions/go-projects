package controller

import (
	"encoding/json"
	"go-mvc-app/go-mvc-app/model"
	"go-mvc-app/go-mvc-app/util"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var va = validator.New()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err = va.Struct(user); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = util.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	util.DB.Find(&users)
	response, err := json.Marshal(&users)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "applocation/json")
	w.Write(response)
}
