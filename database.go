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
		mustEncode(200, w, info)
	} else {
		emitError(404, w, "not_found", "no_db_file")
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
		emitError(412, w, "file_exists", err.Error())
	}
}

func deleteDB(args []string, w http.ResponseWriter, req *http.Request) {
	err := destroyDatabase(args[0])
	if err == nil {
		mustEncode(200, w, map[string]interface{}{"ok": true})
	} else {
		emitError(412, w, "not_found", err.Error())
	}
}

func dbChanges(args []string, w http.ResponseWriter, req *http.Request) {
	emitError(404, w, "not_implemented", "Not supporting changes yet.")
}
