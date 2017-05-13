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
	AltarHandler struct {
		rest.ResourceHandler
		repo domain.DayOfTheDeadRepository
	}
)

func NewAltarHandler(repo domain.DayOfTheDeadRepository) *AltarHandler {
	return &AltarHandler{repo: repo}
}

func (a AltarHandler) Post(parentIds []string, payload catrina.Payload) (code int, body catrina.Payload, err error) {
	var template, shrine *domain.Shrine

	template = &domain.Shrine{}

	err = json.Unmarshal(payload, template)

	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, err
	}

	shrine, err = usecases.CreateShrine(*template, a.repo)

	if err != nil {
		return http.StatusInternalServerError, catrina.EmptyBody, err
	}

	str, _ := json.MarshalIndent(shrine, "", "    ")

	return http.StatusOK, str, nil
}

func (a AltarHandler) Get(id string, parentIds []string) (code int, body catrina.Payload, err error) {
	var shrine *domain.Shrine

	shrineId, err := strconv.Atoi(string(id))

	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, errors.New("ID must be an integer")
	}

	shrine = a.repo.GetShrineByID(shrineId)

	if shrine == nil {
		return http.StatusNotFound, catrina.EmptyBody, errors.New("That shrine does not exist yet")
	}

	str, _ := json.MarshalIndent(shrine, "", "    ")

	return http.StatusOK, str, nil
}

func (a AltarHandler) Delete(id string, parentIds []string) (code int, body catrina.Payload, err error) {
	var shrineId int

	shrineId, err = strconv.Atoi(string(id))

	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, errors.New("ID must be an integer")
	}
	err = a.repo.DeleteShrine(shrineId)

	if err != nil {
		return http.StatusInternalServerError, catrina.EmptyBody, err
	}

	return http.StatusOK, catrina.EmptyBody, nil
}
