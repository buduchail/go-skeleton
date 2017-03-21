package usecases

import (
	"errors"
	"skel/domain"
)

func CreateShrine(template domain.Shrine, repo domain.DayOfTheDeadRepository) (shrine *domain.Shrine, err error) {

	if template.ID > 0 {
		return nil, errors.New("New shrine object must have an empty ID")
	}

	mexican := repo.GetDeadPersonById(template.MexicanID)
	if mexican == nil {
		return nil, errors.New("A shrine has to be dedicated to an illustrious mexican, none was found by that ID")
	}

	if template.Tribute == "" {
		template.Tribute = "Tribute to " + mexican.Name
	}

	shrine, err = domain.CreateShrine(template)

	if err != nil {
		shrine = nil
		return
	}

	err = repo.SaveShrine(shrine)

	if err != nil {
		shrine = nil
	}

	return
}
