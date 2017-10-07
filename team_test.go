package dashi

import (
	"reflect"
	"testing"
)

var teamData = []byte(`
services:
    - name: service name
      deploys:
          - name: app name
            runtime: location
            url: dashboard url
`)

func TestParseTeam(t *testing.T) {
	deploy := &Deploy{
		Name:    "app name",
		Runtime: "location",
		URL:     "dashboard url",
	}
	service := &Service{
		Name:    "service name",
		Deploys: []*Deploy{deploy},
	}
	want := &Team{
		Services: []*Service{service},
	}
	got, err := ParseTeam(teamData)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
