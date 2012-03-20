package main

import (
	"fmt"
	"testing"
)

type sample struct {
	Method  string
	Path    string
	Handler routeHandler
	Args    []string
}

var routeExemplars = []sample{
	sample{"GET", "/", serverInfo, []string{}},
	sample{"GET", "/_all_dbs", listDatabases, []string{}},
	sample{"GET", "/_busy", reservedHandler, []string{"busy"}},
	sample{"GET", "/mydatabase", dbInfo, []string{"mydatabase"}},
	sample{"PUT", "/mydatabase", createDB, []string{"mydatabase"}},
	sample{"DELETE", "/mydatabase", deleteDB, []string{"mydatabase"}},
}

func stringyEquals(a, b interface{}) bool {
	return fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b)
}

func TestVariousRouting(t *testing.T) {
	for _, s := range routeExemplars {
		re, args := findHandler(s.Method, s.Path)
		if !stringyEquals(re.Handler, s.Handler) {
			t.Fatalf("Returned the incorrect handler for %v:%v - %#v",
				s.Method, s.Path, re)
		}
		if !stringyEquals(args, s.Args) {
			t.Fatalf("on %v:%v - Expected args %#v, got %#v",
				s.Method, s.Path, s.Args, args)
		}
	}
}
