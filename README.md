Dashi is a redirector.

## Format

```
services:
    - name: service name
      deploys:
          - name: app name
            runtime: location
            url: dashboard url
```

## User experience

You write YAML files containing the configuration of your dashboards with a
file per team containing that teams service definitions. These are
automatically compiled into the application and served when requests are made.

A search endpoint matches queries to dashboards and returns a redirect on
success or an HTTP 404 if a match isn't found:

```
GET /:query
```

The query must match the form `<service> <deploy>`. Dashi can be setup as a
search bookmark in Chrome so that you can type `dashi myservice production` in
the omnibox and get redirected to the relevant dashboard.

If the query is empty or only contains a service, and more than one dashboard
matches it, the list of relevant dashboards will be returned. If the
`Content-Type` is `text/plain` the service will return a textual listing of
the matching dashboards. If the `Content-Type` is `text/html` the service will
return a web page listing the matching dashboards.
