package dashi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandlerWithEmptyQuery(t *testing.T) {
	manifest := &Manifest{}
	if err := Unmarshal(teamData, manifest); err != nil {
		t.Fatal(err)
	}
	handler := NewSearchHandler(manifest)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, &bytes.Buffer{})
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	payload := map[string][]*Dashboard{}
	if err := json.Unmarshal(buf, &payload); err != nil {
		t.Fatal(err)
	}
	if len(payload) != 1 {
		t.Fatal("got %d, want 1", len(payload), 1)
	}
}

func TestParseQuery(t *testing.T) {
	cases := []struct {
		query   string
		service string
		deploy  string
	}{
		{"/", "", ""},
		{"/service", "service", ""},
		{"/service deploy", "service", "deploy"},
	}
	for _, tt := range cases {
		service, deploy := parseQuery(tt.query)
		if service != tt.service {
			t.Fatalf("%s: got %s, want %s", tt.query, service, tt.service)
		}
		if deploy != tt.deploy {
			t.Fatalf("%s: got %s, want %s", tt.query, deploy, tt.deploy)
		}
	}
}
