package main

import (
	"encoding/json"
	"net/http"
)

func putDocument(args []string, w http.ResponseWriter, req *http.Request) {
	dbname := args[0]
	doc := Document{Id: args[1]}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Transfer-Encoding", "chunked")
	d := json.NewDecoder(req.Body)
	err := d.Decode(&doc.Body)
	if err != nil {
		w.WriteHeader(400)
		emitError("decode", "Error decoding: "+err.Error(), w)
		return
	}
	// TODO:  Check rev, attachments, etc...
	db := getDatabase(dbname)
	if db == nil {
		w.WriteHeader(400)
		emitError("no_db", "No such DB: "+dbname, w)
		return
	}
	rev, err := db.CreateDocument(&doc)
	if err != nil {
		w.WriteHeader(400)
		emitError("generic", err.Error(), w)
		return
	}
	mustEncode(w, map[string]interface{}{"ok": true,
		"id":  doc.Id,
		"rev": string(rev),
	})
}

func getDocument(args []string, w http.ResponseWriter, req *http.Request) {
	dbname := args[0]
	docid := args[1]
	db := getDatabase(dbname)
	w.Header().Set("Content-Type", "application/json")
	if db == nil {
		w.WriteHeader(400)
		emitError("no_db", "No such DB: "+dbname, w)
		return
	}
	doc, err := db.GetDocument(docid)
	if err != nil {
		w.WriteHeader(400)
		emitError("generic", err.Error(), w)
		return
	}
	mustEncode(w, &doc)
}
