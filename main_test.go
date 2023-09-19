package main

import (
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
