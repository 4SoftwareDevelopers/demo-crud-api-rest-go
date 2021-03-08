package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/4softwaredevelopers/demo-crud-api-rest-go/commons"
	"github.com/4softwaredevelopers/demo-crud-api-rest-go/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
