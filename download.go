package main

import "net/http"
// TODO:
// 1. increment downloads count by 1
// 2. respond with .zip

// DownloadHandler handles download route
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// if _, err := app.db.Exec(sqlInsert, team); err != nil {
	// 	fmt.Fprintf(w, "unable to save vote: %s", err)
	// 	return fmt.Errorf("DB.Exec: %v", err)
	// } 
	http.ServeFile(w, r, "./packs/hiphop1.zip")
	//w.Write([]byte(`hello download`))
}