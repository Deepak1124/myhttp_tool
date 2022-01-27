package main

import (
	"reflect"
	"testing"
)

func TestIsCommandValid(t *testing.T) {

	t.Run("checkIsCommandValidWithParallelFlagAsZero_ValidationShouldFail", func(t *testing.T) {
		invalidArg := 0
		urls := []string{"google.com", "http://gmail.com", "http://www.facebook.com"}

		actual := isCommandValid(invalidArg, urls)
		expected := false
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("checkIsCommandValidWithParallelFlagAsNEGATIVENumber_ValidationShouldFail", func(t *testing.T) {
		invalidArg := -1
		urls := []string{"google.com", "http://gmail.com", "http://www.facebook.com"}

		actual := isCommandValid(invalidArg, urls)
		expected := false
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("checkIsCommandValidWithParallelFLAGAsPositiveNumber_ValidationShouldPass", func(t *testing.T) {
		invalidArg := 5
		urls := []string{"google.com", "http://gmail.com", "http://www.facebook.com"}

		actual := isCommandValid(invalidArg, urls)
		expected := true
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("checkIsCommandValidWithZeroArguments_ValidationShouldFail", func(t *testing.T) {
		invalidArg := 5
		urls := []string{}

		actual := isCommandValid(invalidArg, urls)
		expected := false
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}

func TestValidateURL(t *testing.T) {

	t.Run("checkValidateURLWithInvalidUrl_ShouldReturnFalse", func(t *testing.T) {
		invalidUrl := "google.com"
		actual := validateURL(invalidUrl)
		expected := false
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("checkValidateURLWithValidUrl_ShouldReturnTrue", func(t *testing.T) {
		invalidUrl := "http://google.com"
		actual := validateURL(invalidUrl)
		expected := true
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}

func TestAddHttpToDomains(t *testing.T) {

	t.Run("checkAddHttpToDomainsWithInvalidArgs_ShouldReturnCorrectUrlsOnly", func(t *testing.T) {
		args := []string{"##|", "http://gmail.com", "http://www.facebook.com"}
		actual := addHttpToDomains(args)

		expected := []string{"http://gmail.com", "http://www.facebook.com"}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("checkAddHttpToDomainsWithAllValidArgs_ShouldReturnAllUrls", func(t *testing.T) {
		args := []string{"google.com", "http://gmail.com", "http://www.facebook.com"}
		actual := addHttpToDomains(args)

		expected := []string{"http://google.com", "http://gmail.com", "http://www.facebook.com"}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}
