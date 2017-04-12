package resources

import (
	"errors"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/buduchail/catrina"

	"skel/app/usecases"
	"skel/domain"
)

type (
	NivelHandler struct {
		catrina.ResourceHandler
		repo domain.DayOfTheDeadRepository
	}
)

func NewNivelHandler(repo domain.DayOfTheDeadRepository) *NivelHandler {
	return &NivelHandler{repo: repo}
}

func (n NivelHandler) Get(id catrina.ResourceID, parentIds []catrina.ResourceID) (code int, body catrina.Payload, err error) {
	var shrineId, level int

	shrineId, err = strconv.Atoi(string(parentIds[0]))
	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, err
	}

	shrine := n.repo.GetShrineByID(shrineId)
	if shrine == nil {
		return http.StatusNotFound, catrina.EmptyBody, nil
	}

	level, err = strconv.Atoi(string(id))
	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, err
	}

	str, _ := json.MarshalIndent(shrine.GetLevel(level), "", "    ")

	return http.StatusOK, str, nil
}

func (n NivelHandler) Put(id catrina.ResourceID, parentIds []catrina.ResourceID, payload catrina.Payload) (code int, body catrina.Payload, err error) {
	var (
		shrine            *domain.Shrine
		shrineId, shelfId int
	)

	shelfId, err = strconv.Atoi(string(id))
	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, errors.New("Shelf ID must be an integer")
	}

	shrineId, err = strconv.Atoi(string(parentIds[0]))
	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, errors.New("Shrine ID must be an integer")
	}

	gift := domain.Gift{}

	err = json.Unmarshal(payload, &gift)
	if err != nil {
		return http.StatusBadRequest, catrina.EmptyBody, err
	}

	shrine, err = usecases.OfferGift(shrineId, shelfId, gift, n.repo)

	if err != nil {
		return http.StatusInternalServerError, catrina.EmptyBody, err
	}

	str, _ := json.MarshalIndent(shrine, "", "    ")

	return http.StatusOK, str, nil
}
