package main

import (
	"log"
	"net/http"
)

func dbInfo(args []string, w http.ResponseWriter, req *http.Request) {
	dbname := args[0]
	if db, ok := databases[dbname]; ok {
		info, err := db.GetInfo()
		if err != nil {
			log.Fatalf("Error getting DB info: %v", err)
		}
		mustEncode(w, info)
	} else {
		w.WriteHeader(404)
		emitError("not_found", "no_db_file", w)
	}
}

func checkDB(args []string, w http.ResponseWriter, req *http.Request) {
	dbname := args[0]
	if _, ok := databases[dbname]; ok {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
}

func createDB(args []string, w http.ResponseWriter, req *http.Request) {
	err := makeDatabase(args[0])
	if err == nil {
		w.WriteHeader(201)
	} else {
		w.WriteHeader(412)
		emitError("file_exists", err.Error(), w)
	}
}

func deleteDB(args []string, w http.ResponseWriter, req *http.Request) {
	err := destroyDatabase(args[0])
	if err == nil {
		mustEncode(w, map[string]interface{}{"ok": true})
	} else {
		w.WriteHeader(412)
		emitError("not_found", err.Error(), w)
	}
}

func dbChanges(args []string, w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
	emitError("not_implemented", "Not supporting changes yet.", w)
}
