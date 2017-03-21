package domain

func CreateShrine(template Shrine) (*Shrine, error) {

	shrine := Shrine{
		MexicanID: template.MexicanID,
		Tribute:   template.Tribute,
	}

	err := shrine.SetLevels(template.Levels)

	if err != nil {
		return nil, err
	}

	return &shrine, nil
}
