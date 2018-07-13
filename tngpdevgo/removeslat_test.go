package tngpdevgo

import (
	"testing"
)

func Test_RemoveSlatFromDateString_Input_7_Should_be_7(t *testing.T) {
	expectedDay, expectedMonth, expectedYear := "7", "3", "1977"

	actualDay, actualMonth, actualYear := changeRemoveSlat("7/3/1977")
	if actualDay != expectedDay {
		t.Errorf("expected %s but go %s", expectedDay, actualDay)
	}
	if actualMonth != expectedMonth {
		t.Errorf("expected %s but go %s", expectedMonth, actualMonth)
	}
	if actualYear != expectedYear {
		t.Errorf("expected %s but go %s", expectedYear, actualYear)
	}
}
