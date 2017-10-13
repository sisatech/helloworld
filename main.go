package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func page(colour string) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
 <meta charset="UTF-8">
 <title>Hello World</title>
 <meta name="description" content="Description">
 <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
 <!-- Compiled and minified CSS -->
 <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/css/materialize.min.css">
 <!-- Compiled and minified JavaScript -->
 <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/js/materialize.min.js"></script>
 <style>
 body {
	 background-color: ` + colour + `;
 }
 h5 {
	 color: #2396d8;
 }
 </style>
</head>
<body  class="valign-wrapper" style="height:100vh;">
<div class="row">
<div class="center-align">
<img src="data:image/png;base64,` + picture + `" alt="Vorteil">
<h5>WELCOME TO VORTEIL</h5>
</div>
</div>
</body>
</html>`
}

const white = "#FFFFFF"

func main() {
	colour := os.Getenv("BACKGROUND")
	if colour == "" {
		log.Printf("No background color set in BACKGROUND environment variable\n")
		colour = white
	}

	colour = "#" + strings.TrimPrefix(strings.TrimPrefix(colour, "0x"), "#")
	colour = strings.ToUpper(colour)

	if len(colour) != 7 {
		log.Printf("Invalid BACKGROUND color: must be six characters of hexadecimal (like '0xFFFFFF')\n")
		colour = white
	}

	valid := true
	for i := 1; i < len(colour); i++ {
		c := colour[i]
		if c < '0' || c > 'F' || (c > '9' && c < 'A') {
			valid = false
		}
	}
	if !valid {
		log.Printf("Invalid BACKGROUND color: non-hexadecimal characters detected\n")
		colour = white
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SERVING REQUEST")
		rdr := strings.NewReader(page(colour))
		_, err := io.Copy(w, rdr)
		if err != nil {
			log.Printf("Connection error: %v\n", err.Error())
		}
	})
	port := os.Getenv("BIND")
	if port == "" {
		port = "8888"
	}
	log.Printf("Binding port: %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
