package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReplaceInStream(t *testing.T) {
	input := "{\"test\": \"123,456,789,789,123\"}"
	var inBuffer bytes.Buffer
	var outBuffer bytes.Buffer

	readWriter := bufio.NewReadWriter(bufio.NewReader(&inBuffer), bufio.NewWriter(&outBuffer))

	inBuffer.WriteString(input)
	replaceInStream("test", readWriter)

	scanner := bufio.NewScanner(&outBuffer)

	if !scanner.Scan() {
		t.Error("nothing written to outBuffer")
	}

	actual := scanner.Text()
	expected := "{\"test\": \"123,456,789\"}"

	if expected != actual {
		t.Errorf("got %s expected %s", expected, actual)
	}
}

func TestReplace(t *testing.T) {
	fieldName := "test"
	input := "{\"test\": \"123,456,789,789,123\"}"

	expected := "{\"test\": \"123,456,789\"}"
	actual := replace(fieldName, input)

	if expected != actual {
		t.Errorf("got %s expected %s", expected, actual)
	}
}

func TestDistinct(t *testing.T) {
	values := [...]string{"123", "456", "789", "789", "123"}

	expected := [...]string{"123", "456", "789"}
	actual := distinct(values[:])

	if expected != [3]string(actual) {
		t.Errorf("got %v expected %v", expected, actual)
	}
}
