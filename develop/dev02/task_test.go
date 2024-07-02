package main

import "testing"

func TestUnpackString1(t *testing.T) {
	input := "a4bc2d5e"
	expOut := "aaaabccddddde"
	funcOut := UnpackString(input)

	if funcOut != expOut {
		t.Errorf("unpackString() = %q, expOut %q", funcOut, expOut)
	}
}

func TestUnpackString2(t *testing.T) {
	input := "abcd"
	expOut := "abcd"
	funcOut := UnpackString(input)

	if funcOut != expOut {
		t.Errorf("unpackString() = %q, expOut %q", funcOut, expOut)
	}
}

func TestUnpackString3(t *testing.T) {
	input := "45"
	expOut := ""
	funcOut := UnpackString(input)

	if funcOut != expOut {
		t.Errorf("unpackString() = %q, expOut %q", funcOut, expOut)
	}
}

func TestUnpackString4(t *testing.T) {
	input := ""
	expOut := ""
	funcOut := UnpackString(input)

	if funcOut != expOut {
		t.Errorf("unpackString() = %q, expOut %q", funcOut, expOut)
	}
}
