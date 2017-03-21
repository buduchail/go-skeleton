package usecases

import (
	"skel/domain"
)

func ListGiftsByType(type_ string, repository domain.DayOfTheDeadRepository) (gifts []*domain.Gift, err error) {

	return repository.FindGiftsByType(type_), nil
}
