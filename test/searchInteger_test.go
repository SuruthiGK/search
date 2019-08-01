package test

import (
	"reflect"
	"search/internal/search/api"
	"testing"
)

//TestSearch_1
//positive case when string is provided with integers
func TestSearch_1(t *testing.T) {
	expectedResult := []string{"3", "4"}
	result := api.Search("St3ar4")
	isValid := reflect.DeepEqual(result, expectedResult)
	if !isValid {
		t.Errorf("result is incorrect, got: %v, want %v", result, expectedResult)
	}
}

//TestSearch_2
//negative case when string is provided without integers
//lengths of the result is validated
func TestSearch_2(t *testing.T) {
	expectedResult := 0
	result := api.Search("star")
	isValid := (len(result) == expectedResult)
	if !isValid {
		t.Errorf("result is incorrect, got: %d, want %d", len(result), expectedResult)
	}
}

//TestSearch_3
//negative case when string is provided with only special characters
//lengths of the result is validated
func TestSearch_3(t *testing.T) {
	expectedResult := 0
	result := api.Search("@#%$^&")
	isValid := (len(result) == expectedResult)
	if !isValid {
		t.Errorf("result is incorrect, got: %d, want %d", len(result), expectedResult)
	}
}
