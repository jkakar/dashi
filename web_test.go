package dashi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	manifest := &Manifest{}
	if err := Unmarshal(teamData, manifest); err != nil {
		t.Fatal(err)
	}
	handler := NewSearchHandler(manifest)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	cases := []struct {
		name string
		url  string
		want *SearchResults
	}{
		{
			name: "empty query",
			url:  srv.URL,
			want: &SearchResults{
				Dashboards: []*SearchResult{
					&SearchResult{
						Team:    "team name",
						Service: "service name",
						Name:    "dashboard name",
						Env:     "location",
						URL:     "dashboard url",
					},
				},
			},
		},
		{
			name: "service query",
			url:  srv.URL + "/serv",
			want: &SearchResults{
				Dashboards: []*SearchResult{
					&SearchResult{
						Team:    "team name",
						Service: "service name",
						Name:    "dashboard name",
						Env:     "location",
						URL:     "dashboard url",
					},
				},
			},
		},
		{
			name: "service and dashboard query",
			url:  srv.URL + "/serv%20dash",
			want: &SearchResults{
				Dashboards: []*SearchResult{
					&SearchResult{
						Team:    "team name",
						Service: "service name",
						Name:    "dashboard name",
						Env:     "location",
						URL:     "dashboard url",
					},
				},
			},
		},
		{
			name: "unmatched query",
			url:  srv.URL + "/unmatched",
			want: &SearchResults{Dashboards: []*SearchResult{}},
		},
	}

	for _, tt := range cases {
		req, err := http.NewRequest(http.MethodGet, tt.url, &bytes.Buffer{})
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
			t.Fatalf("%s: got %d, want %d", tt.name, resp.StatusCode, http.StatusOK)
		}
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := &SearchResults{}
		if err := json.Unmarshal(buf, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatal("%s: got %#v, want %#v", tt.name, got, tt.want)
		}
	}
}

func TestParseQuery(t *testing.T) {
	cases := []struct {
		query     string
		service   string
		dashboard string
	}{
		{"/", "", ""},
		{"/service", "service", ""},
		{"/service dashboard", "service", "dashboard"},
	}
	for _, tt := range cases {
		service, dashboard := parseQuery(tt.query)
		if service != tt.service {
			t.Fatalf("%s: got %s, want %s", tt.query, service, tt.service)
		}
		if dashboard != tt.dashboard {
			t.Fatalf("%s: got %s, want %s", tt.query, dashboard, tt.dashboard)
		}
	}
}
