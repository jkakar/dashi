package dashi

import (
	"fmt"
	"reflect"
	"testing"
)

var teamData = []byte(`
teams:
    - name: team name
      services:
          - name: service name
            dashboards:
                - name: dashboard name
                  url: https://example.com/dashboard
`)

var multiTeamData = []byte(`
teams:
    - name: team1
      services:
          - name: service name
            dashboards:
                - name: dashboard name
                  url: https://example.com/dashboard/1
    - name: team2
      services:
          - name: service name
            dashboards:
                - name: dashboard name
                  url: https://example.com/dashboard/2
`)

func TestUnmarshal(t *testing.T) {
	dashboard := &Dashboard{
		Name: "dashboard name",
		URL:  "https://example.com/dashboard",
	}
	service := &Service{
		Name:       "service name",
		Dashboards: []*Dashboard{dashboard},
	}
	team := &Team{
		Name:     "team name",
		Services: []*Service{service},
	}
	want := &Manifest{
		Teams: []*Team{team},
	}
	got := &Manifest{}
	if err := Unmarshal(teamData, got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

func TestUnmarshalAppends(t *testing.T) {
	manifest := &Manifest{}
	if err := Unmarshal(teamData, manifest); err != nil {
		t.Fatal(err)
	}
	if err := Unmarshal(multiTeamData, manifest); err != nil {
		t.Fatal(err)
	}
	if len(manifest.Teams) != 3 {
		t.Fatalf("got %d, want 3", len(manifest.Teams))
	}
}

func TestSearchEmptyManifest(t *testing.T) {
	m := &Manifest{}
	got := m.Search("service", "dashboard")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}

func TestSearchWithoutServiceMatch(t *testing.T) {
	m := &Manifest{}
	if err := Unmarshal(teamData, m); err != nil {
		t.Fatal(err)
	}
	got := m.Search("unknown", "dashboard")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}

func TestSearchWithoutDashboardMatch(t *testing.T) {
	m := &Manifest{}
	if err := Unmarshal(teamData, m); err != nil {
		t.Fatal(err)
	}
	got := m.Search("service name", "unknown")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}

func TestSearch(t *testing.T) {
	m := &Manifest{}
	if err := Unmarshal(teamData, m); err != nil {
		t.Fatal(err)
	}
	got := m.Search("service name", "dashboard name")
	if len(got) != 1 {
		t.Fatalf("got %#v, want 1-element slice", got)
	}
	want := &SearchResult{
		Team:    "team name",
		Service: "service name",
		Name:    "dashboard name",
		URL:     "https://example.com/dashboard",
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got %#v, want %#v", got[0], want)
	}
}

func TestSearchMultiple(t *testing.T) {
	m := &Manifest{}
	if err := Unmarshal(multiTeamData, m); err != nil {
		t.Fatal(err)
	}
	res := m.Search("service name", "dashboard name")
	if len(res) != 2 {
		t.Fatalf("got %#v, want 2-element slice", res)
	}
	for i, got := range res {
		want := &SearchResult{
			Team:    fmt.Sprintf("team%d", i+1),
			Service: "service name",
			Name:    "dashboard name",
			URL:     fmt.Sprintf("https://example.com/dashboard/%d", i+1),
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %#v, want %#v", i, got, want)
		}
	}
}

func TestSearchPartialServiceMatch(t *testing.T) {
	m := &Manifest{}
	if err := Unmarshal(teamData, m); err != nil {
		t.Fatal(err)
	}
	got := m.Search("serv", "dashboard name")
	if len(got) != 1 {
		t.Fatalf("got %#v, want 1-element slice", got)
	}
	want := &SearchResult{
		Team:    "team name",
		Service: "service name",
		Name:    "dashboard name",
		URL:     "https://example.com/dashboard",
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got %#v, want %#v", got[0], want)
	}
}

func TestSearchPartialDashboardMatch(t *testing.T) {
	m := &Manifest{}
	if err := Unmarshal(teamData, m); err != nil {
		t.Fatal(err)
	}
	got := m.Search("service name", "dashboard")
	if len(got) != 1 {
		t.Fatalf("got %#v, want 1-element slice", got)
	}
	want := &SearchResult{
		Team:    "team name",
		Service: "service name",
		Name:    "dashboard name",
		URL:     "https://example.com/dashboard",
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got %#v, want %#v", got[0], want)
	}
}
