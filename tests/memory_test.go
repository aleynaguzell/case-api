package tests

import (
	"case-api/storage/cache"
	"testing"
)

var InMemory *cache.InMemory //declare

func init() {
	InMemory = cache.New() //assign
}

func TestStore(t *testing.T) {

	nilResult := InMemory.Set("active-tabs", "getir")
	if nilResult != nil {
		t.Fail()
	}
	if value, err := InMemory.Get("active-tabs"); err != nil || value != "getir" {
		t.Fail()
	}
}

func TestGet(t *testing.T) {

	nilResult := InMemory.Set("active-tabs", "getir")
	if nilResult != nil {
		t.Fail()
	}

	value, err := InMemory.Get("active-tabs")
	if err != nil || value != "getir" {
		t.Fail()
	}
}