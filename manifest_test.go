package dashi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatchEmptyManifest(t *testing.T) {
	m := NewManifest(map[string]*Team{})
	got := m.Match("service", "deploy")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}

func TestMatchWithoutServiceMatch(t *testing.T) {
	team, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	manifest := map[string]*Team{"team": team}
	m := NewManifest(manifest)
	got := m.Match("unknown", "deploy")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}

func TestMatchWithoutDeployMatch(t *testing.T) {
	team, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	manifest := map[string]*Team{"team": team}
	m := NewManifest(manifest)
	got := m.Match("service name", "unknown")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}

func TestMatch(t *testing.T) {
	team, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	manifest := map[string]*Team{"team": team}
	m := NewManifest(manifest)
	got := m.Match("service name", "app name")
	if len(got) != 1 {
		t.Fatalf("got %#v, want 1-element slice", got)
	}
	want := &Dashboard{
		Team: "team",
		Name: "app name",
		URL:  "dashboard url",
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got %#v, want %#v", got[0], want)
	}
}

func TestMatchMultiple(t *testing.T) {
	team, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	manifest := map[string]*Team{"team1": team, "team2": team}
	m := NewManifest(manifest)
	res := m.Match("service name", "app name")
	if len(res) != 2 {
		t.Fatalf("got %#v, want 2-element slice", res)
	}
	for i, got := range res {
		want := &Dashboard{
			Team: fmt.Sprintf("team%d", i+1),
			Name: "app name",
			URL:  "dashboard url",
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %#v, want %#v", i, got, want)
		}
	}
}

func TestMatchPartialServiceMatch(t *testing.T) {
	team, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	manifest := map[string]*Team{"team": team}
	m := NewManifest(manifest)
	got := m.Match("serv", "app name")
	if len(got) != 1 {
		t.Fatalf("got %#v, want 1-element slice", got)
	}
	want := &Dashboard{
		Team: "team",
		Name: "app name",
		URL:  "dashboard url",
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got %#v, want %#v", got[0], want)
	}
}

func TestMatchPartialDeployMatch(t *testing.T) {
	team, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	manifest := map[string]*Team{"team": team}
	m := NewManifest(manifest)
	got := m.Match("service name", "app")
	if len(got) != 1 {
		t.Fatalf("got %#v, want 1-element slice", got)
	}
	want := &Dashboard{
		Team: "team",
		Name: "app name",
		URL:  "dashboard url",
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got %#v, want %#v", got[0], want)
	}
}
