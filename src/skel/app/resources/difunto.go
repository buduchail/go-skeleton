package resources

import (
	"errors"
	"net/http"
	"encoding/json"
	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/rest"

	"skel/app/usecases"
	"skel/domain"
)

type (
	DifuntoHandler struct {
		rest.ResourceHandler
		repo domain.DayOfTheDeadRepository
	}
)

func NewDifuntoHandler(repo domain.DayOfTheDeadRepository) *DifuntoHandler {
	return &DifuntoHandler{repo: repo}
}

func (d DifuntoHandler) Post(parentIds []catrina.ResourceID, payload catrina.Payload) (code int, body catrina.Payload, err error) {
	return http.StatusMethodNotAllowed, catrina.EmptyBody, errors.New("Don't mess with the dead")
}

func (d DifuntoHandler) Get(id catrina.ResourceID, parentIds []catrina.ResourceID) (code int, body catrina.Payload, err error) {
	var dead *domain.Dead

	dead, err = usecases.FindDeadPerson(string(id), d.repo)

	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, err
	}

	if dead == nil {
		return http.StatusNotFound, catrina.EmptyBody, errors.New("Who are you looking for?")
	}

	str, _ := json.MarshalIndent(dead, "", "    ")

	return http.StatusOK, str, nil
}

func (d DifuntoHandler) GetMany(parentIds []catrina.ResourceID, params catrina.QueryParameters) (code int, body catrina.Payload, err error) {

	str, _ := json.MarshalIndent(d.repo.GetAllDeadPeople(), "", "    ")

	return http.StatusOK, str, nil
}

func (d DifuntoHandler) Put(id catrina.ResourceID, parentIds []catrina.ResourceID, payload catrina.Payload) (code int, body catrina.Payload, err error) {
	return http.StatusMethodNotAllowed, catrina.EmptyBody, errors.New("Don't mess with the dead")
}

func (d DifuntoHandler) Delete(id catrina.ResourceID, parentIds []catrina.ResourceID) (code int, body catrina.Payload, err error) {
	return http.StatusMethodNotAllowed, catrina.EmptyBody, errors.New("Don't mess with the dead")
}
