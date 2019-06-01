package server

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	myVar2 := 6
	fmt.Fprintf(w, `<h1>LTW Server is running</h1>
	<h2>Lesson 2 Variables and Constant</h2>
	<p>%d</p>`, myVar2)
}
func Serve() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
