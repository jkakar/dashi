package dashi

import (
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Dashboard struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
	URL  string `yaml:"url"`
}

type Service struct {
	Name       string       `yaml:"name"`
	Dashboards []*Dashboard `yaml:"dashboards"`
}

type Team struct {
	Name     string     `yaml:"name"`
	Services []*Service `yaml:"services"`
}

type Manifest struct {
	Teams []*Team `yaml:"teams"`
}

// Parse decodes a YAML document into a manifest.
func Unmarshal(in []byte, manifest *Manifest) error {
	if err := yaml.Unmarshal(in, &manifest); err != nil {
		return err
	}
	return nil
}

type SearchResult struct {
	Team    string
	Service string
	Name    string
	Env     string
	URL     string
}

// Search returns dashboards that match the service and deploy query
// values. Services and deploys match if the query values are either partial
// or exact matches to their names.
func (m *Manifest) Search(service, deploy string) []*SearchResult {
	result := []*SearchResult{}
	for _, t := range m.Teams {
		for _, s := range t.Services {
			if !strings.HasPrefix(s.Name, service) {
				continue
			}
			for _, d := range s.Dashboards {
				if !strings.HasPrefix(d.Name, deploy) {
					continue
				}
				match := &SearchResult{
					Team:    t.Name,
					Service: s.Name,
					Name:    d.Name,
					Env:     d.Env,
					URL:     d.URL,
				}
				result = append(result, match)
			}
		}
	}
	return result
}
