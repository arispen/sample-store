package main
 
import (
	"log"
	"net/http"
	"fmt"
	"database/sql"
    _ "github.com/jackc/pgx/v4/stdlib"
	"os"
	"html/template"
)

var (
	dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
	dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
	instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
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
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
			socketDir = "/cloudsql"
	}
	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	dbPool, err := sql.Open("pgx", dbURI)
	checkError(err)

	rows, err := dbPool.Query(`SELECT "id", "title", "downloads" FROM "downloads";`)
	checkError(err)
	defer rows.Close()
	var packs []packInfo;
	for rows.Next() {
		var id int
		var title string
		var downloads int
		err = rows.Scan(&id, &title, &downloads)
		checkError(err)
		packs = append(packs, packInfo{id: id, title: title, downloads: downloads})
	}

	fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		packOneDownloads := packs[0].downloads
		tmpl := template.Must(template.ParseFiles("template.html"))
		data := templateData{PackOneDownloads: packOneDownloads}
		tmpl.Execute(w, data)
	})
	fmt.Print("listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func getDownloads(title string) (downloads int) {
	return 0
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
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