package mbedder

import (
	"log"
	"net/http"
)

func init() {

	log.SetFlags(log.Ldate | log.Ltime)

	home, err := newResponse("./private/home.html")
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/", home)
}
