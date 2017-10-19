Dashi is a redirector.

## Format

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

## User experience

You write YAML files containing the configuration of your dashboards, with
services grouped together with the teams that own them. These are loaded by
the application and used to serve results when requests are made.

A search endpoint matches queries to dashboards and returns a redirect on
success or an HTTP 404 if a match isn't found:

```
GET /:query
```

The query must match the form `<service> <deploy>`. Dashi can be setup as a
search bookmark in Chrome so that you can type `dashi myservice production` in
the omnibox and get redirected to the relevant dashboard. Partial matches
against both the service and deploy are accepted, so `dashi mys prod` would be
equivalent.

If the query is empty or only contains a service, and more than one dashboard
matches it, the list of relevant dashboards will be returned. If the `Accept`
header contains `application/json` the service will return a JSON object for
the matching dashboards. If the `Accept` header contains `text/html` the
service will respond with an HTTP 404 if no match is found, perform a redirect
if exactly one match is found, or a web page listing all matching dashboards.

## Examples

Start the API:

```
go run cmd/dashi-api/main.go
```

And make some requests to see what the results look like:

```
curl --silent -H 'Content-Type: application/json' localhost:8080 | jq .
curl --silent -H 'Content-Type: application/json' localhost:8080/domain | jq .
curl --silent -H 'Content-Type: application/json' localhost:8080/domain%20ie | jq .
```
