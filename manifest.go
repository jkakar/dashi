package dashi

type Match struct {
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

func (m *Manifest) Match(service, deploy string) []*Match {
	res := []*Match{}
	for name, t := range m.teams {
		res = append(res, m.matchServices(name, service, t.Services)...)
	}
	return res
}

func (m *Manifest) matchServices(team, name string, services []*Service) []*Match {
	res := []*Match{}
	for _, s := range services {
		if s.Name == name {
			res = append(res, m.matchDeploys(team, s.Name, s.Deploys)...)
		}
	}
	return res
}

func (m *Manifest) matchDeploys(team, name string, deploys []*Deploy) []*Match {
	res := []*Match{}
	for _, d := range deploys {
		if d.Name == name {
			match := &Match{
				Team: team,
				Name: d.Name,
				URL:  d.URL,
			}
			res = append(res, match)
		}
	}
	return res
}
