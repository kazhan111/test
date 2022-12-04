package test

import (
	"fmt"
	"testing"
)

func ErrorString(list []string, want string, got string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

func TestOneElem(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if want != got {
		t.Error(ErrorString(list, got, want))
	}

}
func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if want != got {
		t.Error(ErrorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if want != got {
		t.Error(ErrorString(list, got, want))
	}
}
