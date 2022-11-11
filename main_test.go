package main

import "testing"

func TestContainsAny(t *testing.T) {
	s := "this is some random string"
	ch := "a"

	result := containsAny(s, ch)
	if result != true {
		t.Errorf("containsAny('%v','%v') function FAILED, expected %v, got %v", s, ch, true, result)
	} else {
		t.Logf("containsAny('%v','%v') function PASSED, expected %v, got %v", s, ch, true, result)
	}
}
