package moon

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expect := 4
	if got != expect {
		t.Errorf("Did not get expected result. Got: '%v',wanted '%v'", got, expect)
	}
}

// func TestSubs(t *testing.T) {
// 	got := 10 - 5
// 	expec := 4
// 	if got != expec {
// 		t.Errorf("Did not get expected result. Got: '%v',wanted '%v'", got, expect)
// 	}
// }
