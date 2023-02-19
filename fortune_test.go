package main

import (
	"context"
	"testing"
)

type testWebRequest struct {
}

func (t testWebRequest) FetchBytes(ctx context.Context, typeFort int) ([]byte, error) {
	return []byte(`{"content":"this is test"}`), nil
}

func TestGetContent(t *testing.T) {
	want := "this is test"
	got, err := getFortune(testWebRequest{}, 1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got.Content {
		t.Errorf("Content want: %s, got: %s", want, got)
	}
}

func TestSplitMultilineString(t *testing.T) {
	var fortune = FortuneResult{Content: "This is a long\nmultiline string that\nneeds to be split\ninto multiple lines.\n"}
	maxLength := 12
	expectedOutput := "This is a \nlong \nmultiline \nstring that \nneeds to be \nsplit \ninto \nmultiple lines."

	output := fortune.splitString(maxLength)
	if output != expectedOutput {
		t.Errorf("Expected output: %q\\nBut got: %q\\n", expectedOutput, output)
	}
}
