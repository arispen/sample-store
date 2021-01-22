package main
 
import (
	"log"
	"net/http"
	"fmt"
)
 
func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte(`
			<body><marquee> witaj świecie ( ͡° ͜ʖ ͡°) </marquee></body>
		`))
	})
	fmt.Print("listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
