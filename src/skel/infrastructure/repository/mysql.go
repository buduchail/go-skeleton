package repository

import (
	"errors"
	"database/sql"
	"encoding/json"
	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/crud"

	"skel/domain"
)

type (
	DayOfTheDeadMySqlRepository struct {
		gifts   catrina.CRUD
		shrines catrina.CRUD
		dead    catrina.CRUD
		stats   map[string]int
	}
)

func NewDayOfTheDeadMySqlRepository(dsn string) (*DayOfTheDeadMySqlRepository, error) {
	var (
		gifts,
		shrines,
		dead catrina.CRUD
		err error
	)

	gifts, err = crud.NewMySqlCRUD(dsn, "ofrendas", []string{"id", "tipo", "nombre"}, hydrateGift)
	if err != nil {
		return nil, err
	}

	shrines, err = crud.NewMySqlCRUD(dsn, "altares", []string{"id", "id_difunto", "dedicatoria", "niveles", "repisas"}, hydrateShrine)
	if err != nil {
		return nil, err
	}

	dead, err = crud.NewMySqlCRUD(dsn, "difuntos", []string{"id", "nombre", "nacido", "fallecido", "campo"}, hydrateDead)
	if err != nil {
		return nil, err
	}

	m := &DayOfTheDeadMySqlRepository{
		gifts:   gifts,
		shrines: shrines,
		dead:    dead,
	}

	m.stats = map[string]int{
		"ofrendas": 0,
		"altares":  0,
		"difuntos": 0,
	}

	return m, nil
}

// Helper methods

func hydrateGift(rows sql.Rows) (interface{}, error) {

	gift := domain.Gift{}

	err := rows.Scan(&gift.ID, &gift.Type, &gift.Name)
	if err != nil {
		return nil, err
	}

	return &gift, nil
}

func hydrateShrine(rows sql.Rows) (interface{}, error) {
	var (
		shelves string = ""
		err     error
	)

	shrine := domain.Shrine{}

	err = rows.Scan(&shrine.ID, &shrine.MexicanID, &shrine.Tribute, &shrine.Levels, &shelves)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(shelves), &shrine.Shelves)
	if err != nil {
		return nil, err
	}

	return &shrine, nil
}

func hydrateDead(rows sql.Rows) (interface{}, error) {

	dead := domain.Dead{}

	err := rows.Scan(&dead.ID, &dead.Name, &dead.Born, &dead.Deceased, &dead.Field)
	if err != nil {
		return nil, err
	}

	return &dead, nil
}

func insertValues(crud catrina.CRUD, values []catrina.Value) (int64, error) {

	lastID, err := crud.Insert(values)
	if err != nil {
		return 0, err
	}

	id, ok := lastID.(int64)
	if !ok {
		return 0, errors.New("Last ID has not a correct type (int64)")
	}

	return id, nil
}

// Public interface

func (r DayOfTheDeadMySqlRepository) FindGiftsByType(t string) []*domain.Gift {

	gifts := make([]*domain.Gift, 0)

	rows, err := r.gifts.SelectWhereFields([]string{"tipo"}, []catrina.Value{t})
	if err != nil {
		// silently ignore this error
		return gifts
	}

	for row := range rows {
		gift, ok := row.Result.(*domain.Gift)
		if !ok {
			// silently ignore this row
		}
		gifts = append(gifts, gift)
	}

	return gifts
}

func (r DayOfTheDeadMySqlRepository) GetGiftById(id int) *domain.Gift {

	obj, err := r.gifts.Select(id)
	if err != nil {
		return nil
	}

	gift, ok := obj.(*domain.Gift)
	if !ok {
		return nil
	}

	return gift
}

func (r DayOfTheDeadMySqlRepository) SaveShrine(shrine *domain.Shrine) error {
	var (
		shelves []byte
		err     error
	)

	shelves, err = json.Marshal(shrine.Shelves)
	if err != nil {
		return err
	}

	values := []catrina.Value{
		shrine.MexicanID,
		shrine.Tribute,
		shrine.Levels,
		string(shelves),
	}

	if shrine.ID > 0 {
		err = r.shrines.Update(shrine.ID, values)
	} else {
		shrine.ID, err = insertValues(r.shrines, values)
	}

	return err
}

func (r DayOfTheDeadMySqlRepository) GetAllShrines() []*domain.Shrine {

	shrines := make([]*domain.Shrine, 0)

	rows, err := r.shrines.SelectWhereExpression("1", []catrina.Value{})
	if err != nil {
		// silently ignore this error
		return shrines
	}

	for row := range rows {
		shrine, ok := row.Result.(*domain.Shrine)
		if !ok {
			// silently ignoring this row
			continue
		}
		shrines = append(shrines, shrine)
	}

	return shrines
}

func (r DayOfTheDeadMySqlRepository) GetShrineByID(id int) *domain.Shrine {

	obj, err := r.shrines.Select(id)
	if err != nil {
		return nil
	}

	shrine, ok := obj.(*domain.Shrine)
	if !ok {
		return nil
	}

	return shrine
}

func (r DayOfTheDeadMySqlRepository) DeleteShrine(id int) error {

	return r.shrines.Delete(id)
}

func (r DayOfTheDeadMySqlRepository) GetAllDeadPeople() []*domain.Dead {

	people := make([]*domain.Dead, 0)

	rows, err := r.dead.SelectWhereExpression("1", []catrina.Value{})
	if err != nil {
		// silently ignore this error
		return people
	}

	for row := range rows {
		person, ok := row.Result.(*domain.Dead)
		if !ok {
			// silently ignoring this row
			continue
		}
		people = append(people, person)
	}

	return people
}

func (r DayOfTheDeadMySqlRepository) GetDeadPersonById(id int) *domain.Dead {

	obj, err := r.dead.Select(id)
	if err != nil {
		panic(err)
		return nil
	}

	dead, ok := obj.(*domain.Dead)
	if !ok {
		panic(err)
		return nil
	}

	return dead
}

func (r DayOfTheDeadMySqlRepository) GetStats() map[string]int {
	return r.stats
}
