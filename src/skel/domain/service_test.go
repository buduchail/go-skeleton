package domain

import "testing"

func newTemplate() Shrine {
	return Shrine{
		MexicanID: 1,
		Tribute:   "Tribute to Frida Kahlo",
		Levels:    3,
	}
}

func TestCreateShrine(t *testing.T) {

	template := newTemplate()

	shrine, err := CreateShrine(template)
	if shrine == nil || err != nil {
		t.Errorf("Failed creating shrine (%v)", err)
	}
}

func TestCreateShrine_Values(t *testing.T) {

	template := newTemplate()

	shrine, err := CreateShrine(template)

	if err != nil ||
		shrine.MexicanID != template.MexicanID ||
		shrine.Tribute != template.Tribute ||
		len(shrine.Shelves) != template.Levels {
		t.Errorf("Shrine created with wrong values (expected %v, got %v)", template, shrine)
	}
}

func TestCreateShrine_WrongLevel(t *testing.T) {

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

		template := newTemplate()
		template.Levels = tt.levels

		_, err := CreateShrine(template)

		if (err == nil) != tt.valid {
			t.Errorf("A shrine can not be created with this number of levels: %d", tt.levels)
		}
	}
}
