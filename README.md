Dashi is a redirector.

## Overview

You write YAML files containing the configuration of your dashboards, with
services grouped together with the teams that own them. These are loaded by
the application and used to serve results when requests are made.

A search endpoint matches queries to dashboards. It issues redirect to the
dashboard when a single match is found, displays a list of dashboards if more
than one matches the query, or shows all dashboards if none matched the query:

```
GET /:query
```

The query must match the form `<service> <dashboard>`. Dashi can be setup as a
search bookmark in Chrome so that you can type `d service production` in the
omnibox and get redirected to the relevant dashboard. Partial matches against
both the service and dashboard are accepted, so `d ser prod` would be
equivalent.

If the `Accept` header contains `application/json` the service will return a
JSON object for the matching dashboards. If the `Accept` header contains
`text/html` the service will respond with an HTTP 404 if no match is found,
perform a redirect if exactly one match is found, or a web page listing all
matching dashboards.

### CLI

Add the following to your `~/.bash_profile` to create a simple CLI that will
allow you to open dashboards directly from your terminal:

```bash
dashi() {
    QUERY=$(python -c "import urllib;print urllib.quote(raw_input())" <<< "$@")
    open https://dashi-api.herokuapp.com/$QUERY
}
alias ds=dashi
```

With that you can run `ds ser prod`, for example, to open a dashboard.

## Manifest format

```
teams:
    - name: team name
      services:
          - name: service name
            dashboards:
                - name: dashboard name
                  env: location
                  url: dashboard url
```

## Examples

Start the API:

```
go run cmd/dashi-api/main.go
```

And make some requests to see what the results look like:

```
curl --silent -H 'Accept: application/json' localhost:8080 | jq .
curl --silent -H 'Accept: application/json' localhost:8080/domain | jq .
curl --silent -H 'Accept: application/json' localhost:8080/domain%20ie | jq .
```
