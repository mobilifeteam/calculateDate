package tngpdevgo

import (
	"testing"
)

func Test_DateFormatProcess_Should_Be_Monday7March1979(t *testing.T) {
	expected := "Monday, 7 March 1977"
	startDate := "Monday, 7 March 1977"
	actual := DateFormatProcess(startDate)

	if expected != actual {
		t.Errorf("expected Monday, 7 March 1977 but get %s", actual)
	}

}
