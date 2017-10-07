package dashi

import yaml "gopkg.in/yaml.v2"

// Deploy represents a live service running on a runtime.
type Deploy struct {
	Name    string `yaml:"name"`
	Runtime string `yaml:"runtime"`
	URL     string `yaml:"url"`
}

// Service represents a service.
type Service struct {
	Name    string    `yaml:"name"`
	Deploys []*Deploy `yaml:"deploys"`
}

type Manifest struct {
	Services []*Service `yaml:"services"`
}

// Parse decodes a YAML document into a manifest.
func ParseManifest(in []byte) (*Manifest, error) {
	m := &Manifest{}
	if err := yaml.Unmarshal(in, &m); err != nil {
		return nil, err
	}
	return m, nil
}
