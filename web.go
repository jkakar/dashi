package dashi

import (
	"encoding/json"
	"fmt"
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
	switch r.Header.Get("Content-Type") {
	case "application/json":
		h.searchJSON(w, r)
	case "text/html":
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
	switch len(dashboards) {
	case 0:
		w.WriteHeader(http.StatusNotFound)
	case 1:
		http.Redirect(w, r, dashboards[0].URL, http.StatusFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
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
