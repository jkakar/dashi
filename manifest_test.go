package dashi

import "testing"

func TestMatchEmptyManifest(t *testing.T) {
	m := NewManifest(map[string]*Team{})
	got := m.Match("service", "deploy")
	if len(got) != 0 {
		t.Fatalf("got %#v, want empty slice", got)
	}
}
