package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get Expected result, wanted %q got %q \n", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := Depart("Gopher")
	expect := "Goodbye, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get Expected result, wanted %q got %q \n", expect, got)
	}
}

// TABLE dRIVEN Test
func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "hi", expect: "Hello, hi!\n"},
		{input: "", expect: "Hello, !\n"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("not get same as i/p input:%v expected: %q got: %q ", s.input, s.expect, got)
		}
	}

}

// func TestFailureRypes(t *testing.T) {
// 	t.Error("Errorrrrrr")
// 	t.Fatal("Fatal Immediate FAilure")
// 	t.Error("this will not print because of fatal")
// }
