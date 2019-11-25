package handler

import (
	"testing"
)

func compareTestStrings(t *testing.T, extractedText string, expectedText string) {
	if extractedText != expectedText {
		t.Errorf("Must be %s but was %s", expectedText, extractedText)
	}
}

func compareTestBooleans(t *testing.T, calculatedBoolean bool, expectedBoolean bool) {
	if calculatedBoolean != expectedBoolean {
		t.Errorf("Must be\n%t but was\n%t", calculatedBoolean, expectedBoolean)
	}
}
