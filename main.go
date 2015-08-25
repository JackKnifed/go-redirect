package main

import (
			"net/http"
		)

func redir(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "https://"+req.Host+req.RequestURI, http.StatusMovedPermanently)
}

func main() {
		http.ListenAndServe(":8080", http.HandlerFunc(redir))
}
