package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	serverName := os.Getenv("SERVER_NAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlResponse := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
			    <meta charset="UTF-8">
			    <meta name="viewport" content="width=device-width, initial-scale=1.0">
			    <title>Document</title>
			</head>
			<body>
			    <h1>Hello from %s </h1>
			</body>
			</html>
	`, serverName)
		w.Header().Set("Content-type", "text-html; charset=utf-8")
		fmt.Fprint(w, htmlResponse)
	})
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
