package domain

import (
	"testing"
	"reflect"
)

const (
	TEST_LEVELS = 3
)

var (
	food *Gift = &Gift{
		ID:   1,
		Type: "food",
		Name: "Chocolate",
	}
	flower *Gift = &Gift{
		ID:   10,
		Type: "flower",
		Name: "Cempasúchil",
	}
)

func newShrine() *Shrine {
	shrine := Shrine{
		ID:        1,
		MexicanID: 1,
		Tribute:   "Tribute to Frida Kahlo",
	}

	shrine.SetLevels(TEST_LEVELS)
	return &shrine
}

func TestShrine_SetLevels(t *testing.T) {

	levels := []struct {
		levels int
		valid  bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, false},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
	}

	for _, tt := range levels {
		shrine := Shrine{}
		shrine.Levels = tt.levels
		err := shrine.SetLevels(tt.levels)
		if (err == nil) != tt.valid {
			t.Errorf("A shrine can not have this number of levels: %d", tt.levels)
		}
	}
}

func TestShrine_AddGift(t *testing.T) {

	shrine := newShrine()

	err := shrine.AddGift(flower, 1)
	if err != nil {
		t.Errorf("Error adding a gift to a shrine (%s)", err)
	}
}

func TestShrine_AddGift_WrongLevel(t *testing.T) {

	levels := []int{0, TEST_LEVELS + 1}

	shrine := newShrine()

	for _, level := range levels {
		err := shrine.AddGift(flower, level)
		if err == nil {
			t.Errorf("Shrine accepted gift in wrong levels (%d)", level)
		}
	}
}

func TestShrine_AddGift_TooMany(t *testing.T) {

	shrine := newShrine()

	for i := 0; i < MAX_GIFTS; i++ {
		shrine.AddGift(flower, 1)
	}

	err := shrine.AddGift(flower, 1)
	if err == nil {
		t.Errorf("Shrine can only accept up to %d gifts in each levels", MAX_GIFTS)
	}
}

func TestShrine_AddGift_WrongType(t *testing.T) {

	shrine := newShrine()

	shrine.AddGift(flower, 1)

	err := shrine.AddGift(food, 1)
	if err == nil {
		t.Errorf("Shrine levels can only accept gifts of the same type")
	}
}

func TestShrine_GetGifts(t *testing.T) {

	shrine := newShrine()

	shrine.AddGift(flower, 1)

	level := shrine.GetGifts(1)

	if !reflect.DeepEqual(flower, level[0]) {
		t.Errorf("Level must contain added gift")
	}
}

func TestShrine_GetGifts_WrongLevel(t *testing.T) {

	levels := []int{0, TEST_LEVELS + 1}

	shrine := newShrine()

	for _, level := range levels {
		gifts := shrine.GetGifts(level)
		if gifts != nil {
			t.Errorf("Shrine returned gifts from wrong levels (%d)", level)
		}
	}
}
