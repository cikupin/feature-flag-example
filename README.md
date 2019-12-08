# Feature flag example
A simple feature flag example using flagr https://github.com/checkr/flagr

## Requirement

- [flagr](https://github.com/checkr/flagr)

## Commands

Data seed

```bash
$ go run main.go seed
```

HTTP rest api

```bash
$ go run main.go http
```

## Endpoint

1. test toggle feature `http://localhost:8080/toggle-feature`
2. test toggle provider `http://localhost:8080/toggle-provider`
