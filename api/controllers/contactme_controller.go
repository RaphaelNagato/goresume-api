package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RaphaelNagato/goresume-api/api/models"
	"github.com/RaphaelNagato/goresume-api/api/responses"
)

func (server *Server) CreateMessage(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	message := models.ContactMe{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	newMessage, err := message.SaveMessage(server.DB)

	if err != nil {

		formattedError := errors.New("cannot save to database")

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, newMessage.ID))
	responses.JSON(w, http.StatusCreated, newMessage)
}

func (server *Server) GetMessage(w http.ResponseWriter, r *http.Request) {

	message := models.ContactMe{}

	messages, err := message.FindAllMessages(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, messages)
}
