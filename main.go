package main
 
import (
	"log"
	"net/http"
	"fmt"
    _ "github.com/jackc/pgx/v4/stdlib"
	"os"
	"html/template"
)

type packInfo struct {
	id int
	title string
	downloads int
}

type templateData struct {
	PackOneDownloads int
}
 
func main() {
	fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		// TODO: move db init out of handlers
		db, err := InitializeDatabaseConnection()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT id, title, downloads FROM downloads;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var packs []packInfo;
		for rows.Next() {
			var id int
			var title string
			var downloads int
			err = rows.Scan(&id, &title, &downloads)
			if err != nil {
				log.Fatal(err)
			}
			packs = append(packs, packInfo{id: id, title: title, downloads: downloads})
		}
		packOneDownloads := packs[0].downloads
		tmpl := template.Must(template.ParseFiles("template.html"))
		data := templateData{PackOneDownloads: packOneDownloads}
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/download", DownloadHandler)

	fmt.Print("listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func getDownloads(title string) (downloads int) {
	return 0
}

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}