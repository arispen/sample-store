package main
 
import (
	"log"
	"net/http"
	"fmt"
	"database/sql"
    _ "github.com/jackc/pgx/v4/stdlib"
	"os"
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
 
func main() {
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
			socketDir = "/cloudsql"
	}
	
	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	
	// dbPool is the pool of database connections.
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
	
		packs = append(packs, packInfo{id, title, downloads})
		//fmt.Println(name, downloads)
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {

		var html []byte

		for _, pack := range(packs) {
			html = append(html, []byte("hello world ")...)
			html = append(html, []byte(pack.title)...)
		}

		w.Write(html)
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