package resources

import (
	"errors"
	"net/http"
	"encoding/json"
	"skel/app"
	"skel/app/usecases"
	"skel/domain"
)

type (
	DifuntoHandler struct {
		repo domain.DayOfTheDeadRepository
	}
)

func NewDifuntoHandler(repo domain.DayOfTheDeadRepository) *DifuntoHandler {
	return &DifuntoHandler{repo: repo}
}

func (d DifuntoHandler) Post(parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, errors.New("Don't mess with the dead")
}

func (d DifuntoHandler) Get(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	var dead *domain.Dead

	dead, err = usecases.FindDeadPerson(string(id), d.repo)

	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, err
	}

	if dead == nil {
		return http.StatusNotFound, app.EmptyBody, errors.New("Who are you looking for?")
	}

	str, _ := json.MarshalIndent(dead, "", "    ")

	return http.StatusOK, str, nil
}

func (d DifuntoHandler) GetMany(parentIds []app.ResourceID, params app.QueryParameters) (code int, body app.Payload, err error) {

	str, _ := json.MarshalIndent(d.repo.GetAllDeadPeople(), "", "    ")

	return http.StatusOK, str, nil
}

func (d DifuntoHandler) Put(id app.ResourceID, parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, errors.New("Don't mess with the dead")
}

func (d DifuntoHandler) Delete(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, errors.New("Don't mess with the dead")
}
