package domain

type (
	DayOfTheDeadRepository interface {
		FindGiftsByType(t string) []*Gift
		GetGiftById(id int) *Gift

		SaveShrine(shrine *Shrine) error
		GetAlShrines() []*Shrine
		GetShrineByID(id int) *Shrine
		DeleteShrine(id int) error

		GetAllDeadPeople() []*Dead
		GetDeadPersonById(id int) *Dead

		// Dirty hack to get object stats in status endpoint
		GetStats() map[string]int
	}

	Gift struct {
		ID   int
		Type string
		Name string
	}

	Shrine struct {
		ID        int
		MexicanID int
		Tribute   string
		Levels    int
		Shelves   [][]*Gift
	}

	Dead struct {
		ID       int
		Name     string
		Born     string
		Deceased string
		Field    string
	}
)
