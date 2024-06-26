package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("true booleans have different hash keys")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("false booleans have different hash keys")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("true and false booleans have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	integer1 := &Integer{Value: 1}
	integer2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 2}
	diff2 := &Integer{Value: 2}

	if integer1.HashKey() != integer2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if integer1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
