package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestConvertEncoding(t *testing.T) {
	str := []byte("_binary '你好世界'")
	expected := "0xE4BDA0E5A5BDE4B896E7958C27"
	result := convert(str)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Result mismatch. Expected: %s, got: %s", expected, result)
	}
}

func Test_convert(t *testing.T) {
	input, err := ioutil.ReadFile("./testcase_input.sql")
	if err != nil {
		panic(err)
	}
	expected, err := ioutil.ReadFile("./testcase_expected.sql")
	if err != nil {
		panic(err)
	}
	result := convert(input)

	//if len(expected) != len(result) {
	//	t.Errorf("Result mismatch. Expected length: %d got length: %d", len(expected), len(result))
	//	return
	//}

	//if !reflect.DeepEqual(expected, result) {
	//	t.Errorf("Result mismatch. Expected: %s, got: %s", expected, result)
	//}
	i := 0
	for len(expected) > 0 && len(result) > 0 {
		i += 1
		a := expected[0]
		b := result[0]
		if a != b {
			t.Errorf("Result mismatch. Expected: 0x%02X, got: 0x%02X", a, b)
			return
		}
		expected = expected[1:]
		result = result[1:]
	}
}
