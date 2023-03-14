package trimes

import (
	"testing"
)

func TestHeros(t *testing.T) {
	to := "Ben10"
	expect := "Hello, hero Ben10\n"

	if Heros(to) != expect {
		t.Errorf("Expected %v got %v", expect, to)
	}
}

func TestVillains(t *testing.T) {
	to := "Oblivion"
	get := "Hello, Villain Oblivion\n"

	if Villains(to) != get {
		t.Errorf("expectd %v get %v", get, to)
	}
}

func TestHerosTable(t *testing.T) {
	k := []struct {
		input string
		get   string
	}{
		{"Gwen", "Hello, hero Gwen\n"},
		{"Kevin", "Hello, hero Kevin\n"},
	}
	for _, s := range k {
		x := Heros(s.input)
		if x != s.get {
			t.Errorf("expected %v and get %v", s.get, x)
		}
	}
}

func TestVillainsTable(t *testing.T) {
	k := []struct {
		input string
		get   string
	}{
		{"obm", "Hello, Villain obm\n"},
		{"tomo", "Hello, Villain tomo\n"},
	}
	for _, s := range k {
		x := Villains(s.input)
		if x != s.get {
			t.Errorf("expected %v and get %v", s.get, x)
		}
	}
}
