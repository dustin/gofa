package main

type MemoryDatabase struct {
	name string
	docs map[string]Document
}

func (d *MemoryDatabase) GetDocument(id string) (Document, error) {
	doc, ok := d.docs[id]
	if !ok {
		return doc, NoDocument
	}
	return doc, nil
}

func (d *MemoryDatabase) CreateDocument(doc Document) (Revision, error) {
	d.docs[doc.Id] = doc
	return "xxx-new-xxx", nil
}

func (d *MemoryDatabase) GetInfo() (DBInfo, error) {
	return DBInfo{Name: d.name}, nil
}
