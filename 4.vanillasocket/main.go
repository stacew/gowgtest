package main

import (
	"log"
	"net/http"

	"github.com/stacew/gowgtest/4.vanillasocket/snsp"
)

//complete
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s, err := snsp.NewServerAndBaseHandler()
	defer s.Close()
	if err != nil {
		log.Fatalln(err)
	}

	go s.Serve()
	http.Handle("/socket.io/", s)
	http.Handle("/", http.FileServer(http.Dir("./public")))

	err = http.ListenAndServe(":3000", nil) //호출 후에 handler 등록 더 못함
	if err != nil {
		log.Fatalln(err)
	}
}
