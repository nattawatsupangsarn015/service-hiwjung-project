package test

import (
	"example/service-hiwjung-project/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterHomeWithSuccess(t *testing.T, NODE_ENV string) {
	r := api.Router(NODE_ENV)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}
