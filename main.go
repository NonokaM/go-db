package main
import(
	"io"
	"log"
	"net/http"
)

func main() {

	//  Define handler
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// Handler processing details
		io.WriteString(w , "Hello, World!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w , "Posting Article...\n")
	}

	// Register to use handler
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)

	log.Println("server start at port 8080")
	// Start Server(use ListenAndServe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
