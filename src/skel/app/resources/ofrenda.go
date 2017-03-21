package resources

import (
	"errors"
	"strconv"
	"net/http"
	"encoding/json"

	"skel/app"
	"skel/app/usecases"
	"skel/domain"
)

type (
	OfrendaHandler struct {
		repo domain.DayOfTheDeadRepository
	}
)

func NewOfrendaHandler(repo domain.DayOfTheDeadRepository) *OfrendaHandler {
	return &OfrendaHandler{repo: repo}
}

func (o OfrendaHandler) Post(parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (o OfrendaHandler) Get(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {

	i, err := strconv.Atoi(string(id))

	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, errors.New("ID must be an integer")
	}

	// We don't create a use case for this simple query

	gift := o.repo.GetGiftById(i)

	if gift == nil {
		return http.StatusNotFound, app.EmptyBody, errors.New("There is no such a gift")
	}

	str, _ := json.MarshalIndent(gift, "", "    ")

	return http.StatusOK, str, nil
}

func (o OfrendaHandler) GetMany(parentIds []app.ResourceID, params app.QueryParameters) (code int, body app.Payload, err error) {

	types, exists := params["type"]
	if ! exists || len(types) == 0 {
		return http.StatusBadRequest, app.EmptyBody, errors.New("You must provide a gift type (food, drink or flower)")
	}

	// Another simple case, but we wrap it in a use case just to compare

	gifts, err := usecases.ListGiftsByType(types[0], o.repo)

	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, err
	}

	str, _ := json.MarshalIndent(gifts, "", "    ")

	return http.StatusOK, str, nil
}

func (o OfrendaHandler) Put(id app.ResourceID, parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (o OfrendaHandler) Delete(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}
