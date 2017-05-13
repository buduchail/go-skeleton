package resources

import (
	"errors"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/rest"

	"skel/app/usecases"
	"skel/domain"
)

type (
	OfrendaHandler struct {
		rest.ResourceHandler
		repo domain.DayOfTheDeadRepository
	}
)

func NewOfrendaHandler(repo domain.DayOfTheDeadRepository) *OfrendaHandler {
	return &OfrendaHandler{repo: repo}
}

func (o OfrendaHandler) Get(id string, parentIds []string) (code int, body catrina.Payload, err error) {

	i, err := strconv.Atoi(string(id))

	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, errors.New("ID must be an integer")
	}

	// We don't create a use case for this simple query

	gift := o.repo.GetGiftById(i)

	if gift == nil {
		return http.StatusNotFound, catrina.EmptyBody, errors.New("There is no such a gift")
	}

	str, _ := json.MarshalIndent(gift, "", "    ")

	return http.StatusOK, str, nil
}

func (o OfrendaHandler) GetMany(parentIds []string, params catrina.QueryParameters) (code int, body catrina.Payload, err error) {

	types, exists := params["type"]
	if ! exists || len(types) == 0 {
		return http.StatusBadRequest, catrina.EmptyBody, errors.New("You must provide a gift type (food, drink or flower)")
	}

	// Another simple case, but we wrap it in a use case just to compare

	gifts, err := usecases.ListGiftsByType(types[0], o.repo)

	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, err
	}

	str, _ := json.MarshalIndent(gifts, "", "    ")

	return http.StatusOK, str, nil
}
