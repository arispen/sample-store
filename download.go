package main

import (
	"log"
	"net/http"
)
// TODO:
// 1. increment downloads count by 1
// 2. respond with .zip

// DownloadHandler handles download route
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: move db init out of handlers
	db, err := InitializeDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := db.Exec("UPDATE downloads SET downloads=downloads+1 WHERE id=1;"); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=hiphop1.zip")
	http.ServeFile(w, r, "./packs/hiphop1.zip")
}