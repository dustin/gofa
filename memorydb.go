package main

import (
	"errors"
)

type MemoryDatabase struct {
	name string
	docs map[string]*Document
}

func (d *MemoryDatabase) GetDocument(id string) (*Document, error) {
	doc, ok := d.docs[id]
	if !ok {
		return doc, NoDocument
	}
	return doc, nil
}

func (d *MemoryDatabase) DeleteDocument(id string, rev Revision) error {
	doc, ok := d.docs[id]
	if !ok {
		return NoDocument
	}
	if doc.Rev != rev {
		return errors.New("Invalid revision.")
	}
	return nil
}

func (d *MemoryDatabase) CreateDocument(doc *Document) (Revision, error) {
	if existing, ok := d.docs[doc.Id]; ok {
		if existing.Id != doc.Id {
			return existing.Rev, errors.New("Invalid revision.")
		}
	} else {
		if doc.Rev != "" {
			return Revision(""), errors.New("Invalid revision.")
		}
	}
	d.docs[doc.Id] = doc
	doc.Rev = "xxx-new-xxx"
	return doc.Rev, nil
}

func (d *MemoryDatabase) GetInfo() (DBInfo, error) {
	return DBInfo{Name: d.name}, nil
}
