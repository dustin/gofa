package main

import (
	"errors"
)

var NoDocument = errors.New("No document")

type Revision string

type Attachment struct {
	ContentType string
	Body        []byte
}

type Document struct {
	Id          string
	Rev         Revision
	Body        map[string]interface{}
	Attachments map[string]Attachment
}

type DBInfo struct {
	Name          string `json:"db_name"`
	Compacting    bool   `json:"compact_running"`
	FormatVersion int    `json:"disk_format_version"`
	Size          int    `json:"disk_size"`
	DocCount      int    `json:"doc_count"`
	DelCount      int    `json:"doc_del_count"`
	StartTime     uint64 `json:"instance_start_time"`
	PurgeSeq      int    `json:"purge_seq"`
	UpdateSeq     int    `json:"update_seq"`
}

type Database interface {
	GetDocument(id string) (Document, error)
	CreateDocument(doc Document) (Revision, error)
	GetInfo() (DBInfo, error)
}
