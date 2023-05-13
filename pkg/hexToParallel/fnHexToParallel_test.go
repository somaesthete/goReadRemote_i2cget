package hexToParallel

import (
	"reflect"
	"testing"
)

func Test__HexToParallelOneByte__SHOULD_PASS(t *testing.T) {
	got := HexToParallelOneByte("0xffff")
	want := []bool{true, true, true, true, true, true, true, true}
	if reflect.DeepEqual(got, want) {
		t.Logf("Result was correct,\ngot: %v\nwant: %v.", got, want)
	} else {
		t.Errorf("Result was incorrect,\ngot: %v\nwant: %v.", got, want)
	}
}

func Test__HexToParallelOneByte__SHOULD_FAIL(t *testing.T) {
	got := HexToParallelOneByte("0xfbff")
	want := []bool{true, true, true, true, true, true, true, false}
	if reflect.DeepEqual(got, want) {
		t.Logf("Result was correct,\ngot: %v\nwant: %v.", got, want)
	} else {
		t.Errorf("Result was incorrect,\ngot: %v\nwant: %v.", got, want)
	}
}

func Test__HexWordParse__SHOULD_PASS(t *testing.T) {
	got := HexToParallelSanitizeWord("0xff 0x12 0x23 0xba")

	want := []bool{
		true, true, true, true, true, false, true, true,
		true, true, true, true, true, false, true, false,
		false, false, true, false, false, false, true, true,
		true, true, true, true, true, true, true, true,
	}
	if reflect.DeepEqual(got, want) {
		t.Logf("Result was correct,\ngot: %v\nwant: %v.", got, want)
	} else {
		t.Errorf("Result was incorrect,\ngot: %v\nwant: %v.", got, want)
	}
}
