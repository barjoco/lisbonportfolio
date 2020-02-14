package main

import (
	"fmt"
	"net/http"
	"path"
	"utils"
)

const (
	port = ":8084"
)

var (
	thisDir   = utils.GetThisDir()
	staticDir = thisDir + "/data/static"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqFile := path.Join(staticDir, r.URL.Path)
		http.ServeFile(w, r, reqFile)
	})

	fmt.Printf("Site running on port: %s...\n", port)
	http.ListenAndServe(port, nil)
}
