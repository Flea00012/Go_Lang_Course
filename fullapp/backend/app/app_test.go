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

func Test404(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 405 {
		t.Errorf("Status should be 405, but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""
	if respString != expected {
		t.Errorf("Response should be %s, got %s instead", expected, respString)
	}
}

func TestRequest(t *testing.T) {
	r := newRouter()
	mockTLSServer := httptest.NewTLSServer(r)
	req := httptest.NewRequest("POST", mockTLSServer.URL, nil)
	fmt.Printf("req: %v", req.Body)
}

func TestResponse(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	addr := mockServer.Listener.Addr()
	fmt.Println("address is: ", addr)
	
}
