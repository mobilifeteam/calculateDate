package tngpdevgo

import "testing"

func Test_CalculateHours_Should_Be_362492(t *testing.T) {
	expected := 362472
	actual := CalculateHours(15103)

	if actual != expected {
		t.Errorf("expected 362492 but get %d", actual)
	}
}
