package main

import "testing"

type testWebRequest struct {
}

func (t testWebRequest) FetchBytes(typeFort int) ([]byte, error) {
	return []byte(`{"content":"this is test"}`), nil
}

func TestGetContent(t *testing.T) {
	want := "this is test"
	got, err := getFortune(testWebRequest{}, 1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("Content want: %s, got: %s", want, got)
	}
}
