package main

import "database/sql"

var uploadDir string

type FC struct {
	db *sql.DB
}

func NewFileController(db *sql.DB) *FC {
	return &FC{db: db}
}

type cloud struct {
	locateC string
	typeC   string
	nameC   string
}

type locateFile struct {
	Name          string `json:"name"`
	TypePoint     string `json:"typePoint"`
	Locate        string `json:"locate"`
	InCloud       string `json:"inCloud"`
	LocateInCloud string `json:"locateInCloud"`
}

type dFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}
