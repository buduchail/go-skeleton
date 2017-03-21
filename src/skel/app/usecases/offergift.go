package usecases

import (
	"errors"
	"skel/domain"
)

func OfferGift(shrineId int, level int, gift domain.Gift, repo domain.DayOfTheDeadRepository) (*domain.Shrine, error) {

	shrine := repo.GetShrineByID(shrineId)
	if shrine == nil {
		return nil, errors.New("That shrine does not exist")
	}

	offer := repo.GetGiftById(gift.ID)
	if offer == nil {
		return shrine, errors.New("That gift does not exist")
	}

	err := shrine.AddGift(offer, level)
	if err != nil {
		return shrine, err
	}

	repo.SaveShrine(shrine)

	return shrine, nil
}
