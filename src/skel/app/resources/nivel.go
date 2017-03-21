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
	NivelHandler struct {
		repo domain.DayOfTheDeadRepository
	}
)

func NewNivelHandler(repo domain.DayOfTheDeadRepository) *NivelHandler {
	return &NivelHandler{repo: repo}
}

func (n NivelHandler) Post(parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (n NivelHandler) Get(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	var shrineId, level int

	shrineId, err = strconv.Atoi(string(parentIds[0]))
	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, err
	}

	shrine := n.repo.GetShrineByID(shrineId)
	if shrine == nil {
		return http.StatusNotFound, app.EmptyBody, nil
	}

	level, err = strconv.Atoi(string(id))
	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, err
	}

	str, _ := json.MarshalIndent(shrine.GetLevel(level), "", "    ")

	return http.StatusOK, str, nil
}

func (n NivelHandler) GetMany(parentIds []app.ResourceID, params app.QueryParameters) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (n NivelHandler) Put(id app.ResourceID, parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	var (
		shrine            *domain.Shrine
		shrineId, shelfId int
	)

	shelfId, err = strconv.Atoi(string(id))
	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, errors.New("Shelf ID must be an integer")
	}

	shrineId, err = strconv.Atoi(string(parentIds[0]))
	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, errors.New("Shrine ID must be an integer")
	}

	gift := domain.Gift{}

	err = json.Unmarshal(payload, &gift)
	if err != nil {
		return http.StatusBadRequest, app.EmptyBody, err
	}

	shrine, err = usecases.OfferGift(shrineId, shelfId, gift, n.repo)

	if err != nil {
		return http.StatusInternalServerError, app.EmptyBody, err
	}

	str, _ := json.MarshalIndent(shrine, "", "    ")

	return http.StatusOK, str, nil
}

func (n NivelHandler) Delete(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}
