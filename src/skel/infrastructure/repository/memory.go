package repository

import (
	"io/ioutil"
	"encoding/json"
	"skel/domain"
	"errors"
	"strconv"
)

type (
	DayOfTheDeadMemoryRepository struct {
		data struct {
			Ofrendas []*domain.Gift
			Altares  []*domain.Shrine
			Difuntos []*domain.Dead
		}
		index struct {
			GiftByType map[string][]*domain.Gift
			ShrineById map[int]*domain.Shrine
		}
		lastId int
		stats  map[string]int
	}
)

func NewDayOfTheDeadMemoryRepository() (m *DayOfTheDeadMemoryRepository) {
	m = &DayOfTheDeadMemoryRepository{}
	m.stats = map[string]int{
		"ofrendas": 0,
		"altares":  0,
		"difuntos": 0,
	}
	m.index.GiftByType = make(map[string][]*domain.Gift, 0)
	m.index.ShrineById = make(map[int]*domain.Shrine, 0)
	return m
}

func (m *DayOfTheDeadMemoryRepository) LoadData(path string) (err error) {

	str, err := ioutil.ReadFile(path);
	if err != nil {
		return err
	}

	err = json.Unmarshal(str, &m.data);
	if err != nil {
		return err
	}

	m.index.GiftByType = make(map[string][]*domain.Gift, 0)
	m.index.ShrineById = make(map[int]*domain.Shrine, 0)

	for _, gift := range m.data.Ofrendas {
		m.index.GiftByType[gift.Type] = append(m.index.GiftByType[gift.Type], gift)
	}

	for _, shrine := range m.data.Altares {
		m.index.ShrineById[shrine.ID] = shrine
	}

	return nil;
}

func (m *DayOfTheDeadMemoryRepository) FindGiftsByType(t string) []*domain.Gift {
	gifts := make([]*domain.Gift, 0, len(m.data.Ofrendas))

	for _, g := range m.data.Ofrendas {
		if g.Type == t {
			gifts = append(gifts, g)
		}
	}

	return gifts
}

func (m *DayOfTheDeadMemoryRepository) GetGiftById(id int) *domain.Gift {
	if id > 0 && id <= len(m.data.Ofrendas) {
		return m.data.Ofrendas[id-1]
	}
	return nil
}

func (m *DayOfTheDeadMemoryRepository) SaveShrine(shrine *domain.Shrine) error {
	if shrine.ID > 0 {
		old, exists := m.index.ShrineById[shrine.ID]
		if !exists {
			return errors.New("Shrine not found: " + strconv.Itoa(shrine.ID))
		}
		old.Levels = shrine.Levels
		old.Shelves = shrine.Shelves
	} else {
		m.lastId += 1
		shrine.ID = m.lastId
		m.index.ShrineById[shrine.ID] = shrine
	}
	return nil
}

func (m *DayOfTheDeadMemoryRepository) GetAlShrines() []*domain.Shrine {
	return nil
}

func (m *DayOfTheDeadMemoryRepository) GetShrineByID(id int) *domain.Shrine {
	shrine, _ := m.index.ShrineById[id]
	return shrine
}

func (m*DayOfTheDeadMemoryRepository) DeleteShrine(id int) error {

	_, exists := m.index.ShrineById[id]
	if !exists {
		return errors.New("That shrine does not exist")
	}

	delete(m.index.ShrineById, id)

	return nil
}

func (m *DayOfTheDeadMemoryRepository) GetAllDeadPeople() []*domain.Dead {
	return m.data.Difuntos
}

func (m *DayOfTheDeadMemoryRepository) GetDeadPersonById(id int) *domain.Dead {
	if id > 0 && id <= len(m.data.Difuntos) {
		return m.data.Difuntos[id-1]
	}
	return nil
}

func (m *DayOfTheDeadMemoryRepository) GetStats() map[string]int {
	m.stats["ofrendas"] = len(m.data.Ofrendas)
	m.stats["altares"] = len(m.index.ShrineById)
	m.stats["difuntos"] = len(m.data.Difuntos)
	return m.stats
}
