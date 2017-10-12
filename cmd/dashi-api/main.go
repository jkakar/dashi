package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jkakar/dashi"
	"github.com/joeshaw/envdecode"
)

type config struct {
	Port int `env:"PORT,default=8080"`
}

func main() {
	cfg := &config{}
	if err := envdecode.StrictDecode(cfg); err != nil {
		log.Fatalf("loading config failed: %q", err)
	}
	manifest := &dashi.Manifest{}
	if err := dashi.Unmarshal(data, manifest); err != nil {
		log.Fatal(err)
	}
	handler := dashi.NewSearchHandler(manifest)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}
	log.Printf("listening on port %d", cfg.Port)
	log.Fatalf("listen and serve failed: %q", s.ListenAndServe())
}

var data = []byte(`
teams:
    - name: Runtime
      services:
          - name: domain-funnel
            dashboards:
                - name: st-legacy
                  env: common-runtime-st
                  url: https://metrics.librato.com/s/spaces/351543?source=common-runtime-st-shared
                - name: ie-legacy
                  env: common-runtime-ie
                  url: https://metrics.librato.com/s/spaces/351543?source=common-runtime-ie-shared
                - name: va-legacy
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/351543?source=common-runtime-va-shared
          - name: release-manager
            dashboards:
                - name: st01
                  env: common-runtime-st
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-st01
                - name: ie01
                  env: common-runtime-ie
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-ie01
                - name: ie02
                  env: common-runtime-ie
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-ie02
                - name: va01
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-va01
                - name: va02
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-va02
                - name: va03
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-va03
                - name: va04
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-va04
                - name: va05
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-va05
                - name: va06
                  env: common-runtime-va
                  url: https://metrics.librato.com/s/spaces/429997?source=common-runtime-va06
`)
