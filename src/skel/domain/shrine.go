package domain

import (
	"errors"
	"strconv"
)

const (
	// This is an arbitrary limit, but let's consider it
	// a valid business rule
	MAX_GIFTS = 5
)

// Niveles del altar
//
// Los niveles en el altar de muertos representan la cosmovisión, regularmente representando
// el mundo material y el inmaterial o los cuatro elementos, en cada uno de ellos se colocan
// diferentes objetos simbólicos para la cultura, religión o la persona a la que se le dedica
// el altar.
//
// Altares de dos niveles: son una representación de la división del cielo y la tierra
// Altares de tres niveles: representan el cielo, la tierra y el inframundo
// Altares de siete niveles: son el tipo de altar más convencional, representan los siete
// niveles que debe atravesar el alma para poder llegar al descanso o paz espiritual
func (s *Shrine) SetLevels(levels int) error {

	if levels != 2 && levels != 3 && levels != 7 {
		return errors.New("Shrines can only have 2, 3 or 7 levels")
	}

	s.Levels = levels
	s.Shelves = make([][]*Gift, levels)

	return nil
}

func (s *Shrine) AddGift(gift *Gift, level int) error {

	if level < 1 || level > s.Levels {
		return errors.New("This shrine only has " + strconv.Itoa(s.Levels) + " shelves")
	}

	if len(s.Shelves[level-1]) >= MAX_GIFTS {
		return errors.New("Sorry, but a shelf can only hold up to " + strconv.Itoa(MAX_GIFTS) + " gifts")
	}

	for _, g := range s.Shelves[level-1] {
		if g.Type != gift.Type {
			return errors.New("All gifts on a shelf must be of the same type")
		}
	}

	s.Shelves[level-1] = append(s.Shelves[level-1], gift)

	return nil
}

func (s *Shrine) GetGifts(level int) []*Gift {
	if level < 1 || level >= len(s.Shelves) {
		return nil
	}
	return s.Shelves[level-1]
}
