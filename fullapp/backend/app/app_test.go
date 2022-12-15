package fullapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	fmt.Println("testing the app")
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status not ok, got: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World!"
	if respString != expected {
		t.Errorf("Response should be %s, but got %s ", expected, respString)
	}
}