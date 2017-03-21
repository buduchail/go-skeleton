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
	AltarHandler struct {
		repo domain.DayOfTheDeadRepository
	}
)

func NewAltarHandler(repo domain.DayOfTheDeadRepository) *AltarHandler {
	return &AltarHandler{repo: repo}
}

func (a AltarHandler) Post(parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	var template, shrine *domain.Shrine

	template = &domain.Shrine{}

	err = json.Unmarshal(payload, template)

	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, err
	}

	shrine, err = usecases.CreateShrine(*template, a.repo)

	if err != nil {
		return http.StatusInternalServerError, app.EmptyBody, err
	}

	str, _ := json.MarshalIndent(shrine, "", "    ")

	return http.StatusOK, str, nil
}

func (a AltarHandler) Get(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	var shrine *domain.Shrine

	shrineId, err := strconv.Atoi(string(id))

	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, errors.New("ID must be an integer")
	}

	shrine = a.repo.GetShrineByID(shrineId)

	if shrine == nil {
		return http.StatusNotFound, app.EmptyBody, errors.New("That shrine does not exist yet")
	}

	str, _ := json.MarshalIndent(shrine, "", "    ")

	return http.StatusOK, str, nil
}

func (a AltarHandler) GetMany(parentIds []app.ResourceID, params app.QueryParameters) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (a AltarHandler) Put(id app.ResourceID, parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (a AltarHandler) Delete(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	var shrineId int

	shrineId, err = strconv.Atoi(string(id))

	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, errors.New("ID must be an integer")
	}
	err = a.repo.DeleteShrine(shrineId)

	if err != nil {
		return http.StatusInternalServerError, app.EmptyBody, err
	}

	return http.StatusOK, app.EmptyBody, nil
}
