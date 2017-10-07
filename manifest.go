package dashi

import (
	"sort"
	"strings"
)

type Dashboard struct {
	Team string
	Name string
	URL  string
}

type Manifest struct {
	teams map[string]*Team
}

func NewManifest(teams map[string]*Team) *Manifest {
	return &Manifest{teams: teams}
}

func (m *Manifest) Match(service, deploy string) []*Dashboard {
	names := []string{}
	for n, _ := range m.teams {
		names = append(names, n)
	}
	sort.Strings(names)

	res := []*Dashboard{}
	for _, name := range names {
		t := m.teams[name]
		res = append(res, m.matchServices(name, service, deploy, t.Services)...)
	}
	return res
}

func (m *Manifest) matchServices(team, name, deploy string, services []*Service) []*Dashboard {
	res := []*Dashboard{}
	for _, s := range services {
		if strings.HasPrefix(s.Name, name) {
			res = append(res, m.matchDeploys(team, deploy, s.Deploys)...)
		}
	}
	return res
}

func (m *Manifest) matchDeploys(team, name string, deploys []*Deploy) []*Dashboard {
	res := []*Dashboard{}
	for _, d := range deploys {
		if strings.HasPrefix(d.Name, name) {
			match := &Dashboard{
				Team: team,
				Name: d.Name,
				URL:  d.URL,
			}
			res = append(res, match)
		}
	}
	return res
}
