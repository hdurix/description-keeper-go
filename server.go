package description_keeper

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
