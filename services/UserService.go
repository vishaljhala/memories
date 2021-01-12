package services

import (
	"encoding/json"
	"memories/application"
	"memories/models"
	"net/http"

	"github.com/gorilla/mux"
)

func AddUser(response http.ResponseWriter, request *http.Request) {
	db := application.Storage()
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		application.Logger().Info(err.Error())
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]bool{"ok": false})
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		application.Logger().Info(result.Error.Error())
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]bool{"ok": false})
		return
	}
	json.NewEncoder(response).Encode(map[string]uint{"id": user.ID})
	application.Logger().Info("Add User Successful")
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	db := application.Storage()
	vars := mux.Vars(request)
	result := db.Delete(&models.User{}, vars["id"])
	if result.Error != nil {
		application.Logger().Info(result.Error.Error())
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]bool{"ok": false})
		return
	}
	application.Logger().Info("Delete User Successful")
	response.WriteHeader(204)
}

func ListUsers(response http.ResponseWriter, request *http.Request) {
	db := application.Storage()
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		application.Logger().Info(result.Error.Error())
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]bool{"ok": false})
		return
	}
	json.NewEncoder(response).Encode(users)
}
