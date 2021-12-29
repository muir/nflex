# nflex - common interface to parsed config files

[![GoDoc](https://godoc.org/github.com/muir/nflex?status.png)](https://pkg.go.dev/github.com/muir/nflex)
![unit tests](https://github.com/muir/nflex/actions/workflows/go.yml/badge.svg)
[![report card](https://goreportcard.com/badge/github.com/muir/nflex)](https://goreportcard.com/report/github.com/muir/nflex)
[![codecov](https://codecov.io/gh/muir/nflex/branch/main/graph/badge.svg)](https://codecov.io/gh/muir/nflex)

Install:

	go get github.com/muir/nflex

---

Nflex is a common wrapper around multiple configuration file unpackers.  It is lower-level than
just unpacking into an `map[string]interface{}` and thus avoids adding incorrect type information
to YAML files that are parsed that way.

It currently supports:

- YAML
- JSON

It supports merging data from multiple files.

