package test

import (
	"reflect"
	"search/internal/search/api"
	"testing"
)

func TestSearch(t *testing.T) {
	// t.Fatal("not implemented")
	result := []string{}
	expectedResult := []string{"3", "4"}
	result = api.Search("Su3r4")
	isValid := reflect.DeepEqual(result, expectedResult)
	if !isValid {
		t.Errorf("result is incorrect, got: %v, want %v", result, expectedResult)
	}
}
