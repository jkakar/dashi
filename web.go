package dashi

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type SearchResults struct {
	Dashboards []*SearchResult `json:"dashboards"`
}

type SearchHandler struct {
	manifest *Manifest
}

func NewSearchHandler(manifest *Manifest) *SearchHandler {
	return &SearchHandler{manifest: manifest}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := r.Header.Get("Accept")
	switch {
	case strings.Contains(v, "application/json"):
		h.searchJSON(w, r)
	case strings.Contains(v, "text/html"):
		h.searchHTML(w, r)
	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func (h *SearchHandler) searchJSON(w http.ResponseWriter, r *http.Request) {
	service, dashboard := parseQuery(r.URL.Path)
	dashboards := h.manifest.Search(service, dashboard)
	res := &SearchResults{Dashboards: dashboards}
	buf, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", buf)
}

func (h *SearchHandler) searchHTML(w http.ResponseWriter, r *http.Request) {
	service, dashboard := parseQuery(r.URL.Path)
	dashboards := h.manifest.Search(service, dashboard)
	if len(dashboards) == 1 {
		http.Redirect(w, r, dashboards[0].URL, http.StatusFound)
		return
	}

	results := h.manifest
	if len(dashboards) >= 1 {
		results = processResults(dashboards)
	}

	tmpl, err := template.New("index.html").ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseQuery(query string) (service string, dashboard string) {
	query = strings.TrimPrefix(query, "/")
	parts := strings.Split(query, " ")
	if len(parts) > 0 {
		service = parts[0]
	}
	if len(parts) > 1 {
		dashboard = parts[1]
	}
	return
}

func processResults(results []*SearchResult) *Manifest {
	manifest := &Manifest{
		Teams: []*Team{},
	}
	teams := map[string]*Team{}
	services := map[string]*Service{}

	for _, r := range results {
		team, ok := teams[r.Team]
		if !ok {
			team = &Team{
				Name:     r.Team,
				Services: []*Service{},
			}
			teams[r.Team] = team
			manifest.Teams = append(manifest.Teams, team)
		}

		service, ok := services[r.Service]
		if !ok {
			service = &Service{
				Name:       r.Service,
				Dashboards: []*Dashboard{},
			}
			services[r.Service] = service
			team.Services = append(team.Services, service)
		}

		dashboard := &Dashboard{
			Name: r.Name,
			URL:  r.URL,
		}
		service.Dashboards = append(service.Dashboards, dashboard)
	}
	return manifest
}
