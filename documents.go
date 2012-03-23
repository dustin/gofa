package main

import (
	"encoding/json"
	"net/http"
)

func getRequestRev(req *http.Request) Revision {
	req.ParseForm()
	revstr := req.Form.Get("rev")
	if revstr == "" {
		revstr = req.Header.Get("If-Match")
		if revstr != "" {
			revstr = revstr[1:-1]
		}
	}
	return Revision(revstr)
}

func getDocumentRev(j map[string]interface{}) (rv Revision) {
	if rraw, ok := j["_rev"]; ok {
		if rstr, ok := rraw.(string); ok {
			rv = Revision(rstr)
		}
	}
	return
}

func putDocument(args []string, w http.ResponseWriter, req *http.Request) {
	dbname := args[0]
	doc := Document{Id: args[1]}
	w.Header().Set("Content-Type", "application/json")
	d := json.NewDecoder(req.Body)
	err := d.Decode(&doc.Body)
	if err != nil {
		emitError(400, w, "decode", "Error decoding: "+err.Error())
		return
	}

	doc.Rev = getRequestRev(req)
	if doc.Rev == "" {
		doc.Rev = getDocumentRev(doc.Body)
	}

	db := getDatabase(dbname)
	if db == nil {
		emitError(400, w, "no_db", "No such DB: "+dbname)
		return
	}
	rev, err := db.CreateDocument(&doc)
	if err != nil {
		emitError(400, w, "generic", err.Error())
		return
	}
	mustEncode(200, w, map[string]interface{}{"ok": true,
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
		emitError(400, w, "no_db", "No such DB: "+dbname)
		return
	}
	doc, err := db.GetDocument(docid)
	if err != nil {
		emitError(400, w, "generic", err.Error())
		return
	}
	mustEncode(200, w, &doc)
}

func rmDocument(args []string, w http.ResponseWriter, req *http.Request) {
	dbname := args[0]
	docid := args[1]
	db := getDatabase(dbname)
	w.Header().Set("Content-Type", "application/json")
	if db == nil {
		emitError(400, w, "no_db", "No such DB: "+dbname)
		return
	}

	rev := getRequestRev(req)

	err := db.DeleteDocument(docid, rev)
	if err != nil {
		emitError(400, w, "generic", err.Error())
		return
	}
	mustEncode(200, w, map[string]interface{}{
		"ok": true, "rev": string(rev),
	})
}
