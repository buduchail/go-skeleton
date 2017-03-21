package usecases

import (
	"errors"
	"strconv"
	"skel/domain"
)

func FindDeadPerson(id string, repository domain.DayOfTheDeadRepository) (dead *domain.Dead, err error) {

	i, err := strconv.Atoi(id)

	if err != nil {
		return nil, errors.New("ID must be an integer")
	}

	dead = repository.GetDeadPersonById(i)

	if dead == nil {
		return nil, nil
	}

	return dead, nil
}
