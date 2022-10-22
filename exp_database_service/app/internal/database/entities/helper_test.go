package entities

import (
	"testing"
)

func TestStringToSortType(t *testing.T) {
	if asc := StringToSortType(""); asc != Asc {
		t.Errorf("can't get sort type")
	}
	if des := StringToSortType("DESC"); des != Desc {
		t.Errorf("can't get sort type")
	}
}
