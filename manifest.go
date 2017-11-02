package dashi

import (
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Dashboard struct {
	Name string `yaml:"name"`
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
	r := &Manifest{}
	if err := yaml.Unmarshal(in, r); err != nil {
		return err
	}
	manifest.Teams = append(manifest.Teams, r.Teams...)
	return nil
}

type SearchResult struct {
	Team    string `json:"team"`
	Service string `json:"service"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

// Search returns dashboards that match the service and dashboard query
// values. Services and dashboards match if the query values are either partial
// or exact matches to their names.
func (m *Manifest) Search(service, dashboard string) []*SearchResult {
	result := []*SearchResult{}
	for _, t := range m.Teams {
		for _, s := range t.Services {
			if !strings.HasPrefix(s.Name, service) {
				continue
			}
			for _, d := range s.Dashboards {
				if !strings.HasPrefix(d.Name, dashboard) {
					continue
				}
				match := &SearchResult{
					Team:    t.Name,
					Service: s.Name,
					Name:    d.Name,
					URL:     d.URL,
				}
				result = append(result, match)
			}
		}
	}
	return result
}
