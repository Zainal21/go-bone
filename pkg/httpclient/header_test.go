// Package httpclient
package httpclient

import "testing"

func TestNormalize(t *testing.T) {

	str := "jakarta-medan"
	expected := "Jakarta-Medan"

	result := Normalize(str)

	if result == expected {
		t.Logf("expected %s, got %s", expected, result)
	} else {
		t.Errorf("expected %s, got %s", expected, result)
	}

}
