package testdemos

import "testing"

/*

t.FallNow() t.Fail()
t.Fatal(...) t.Error(...)
t.Fatalf(,...) t.Errorf(,...)

go test
go test -v
go test -run <Test>Greet

go test -cover
go test -coverprofile cover.out
go tool cover -func cover.out
go tool cover -html cover.out
go test -coverprofile count.out -covermode count


Log/Logf
Helper
Skip, Skipf
Run
Parallel

*/

func TestGreet(t *testing.T) {
	got := Greet("_name_")
	expected := "hello, _name_\n"

	if got != expected {
		t.Errorf("Did not get expected result! expected: %v ,got: %v", expected, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("_name_")
	expected := "Goodbye, _name_\n"

	if got != expected {
		t.Errorf("Did not get expected result! expected: %v ,got: %v", expected, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	testsSc := []struct {
		input  string
		expect string
	}{
		{input: "_name_", expect: "hello, _name_\n"},
		{input: "john", expect: "hello, john\n"},
	}

	for i, s := range testsSc {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("[%d] Did not get expected result\n      input: %v\n   expected: %v\n        got: %v",
				i, s.input, s.expect, got,
			)
		}
	}
}
