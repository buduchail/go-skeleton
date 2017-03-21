package domain

import "testing"

var shrineLevels = []struct {
	level int
	valid bool
}{
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

func TestShrineSetLevels(t *testing.T) {
	for _, tt := range shrineLevels {
		shrine := Shrine{}
		shrine.Levels = tt.level
		err := shrine.SetLevels(tt.level)
		if (err == nil) != tt.valid {
			t.Errorf("A shrine can not have this number of levels: %d", tt.level)
		}
	}
}
