package main

import "github.com/lilleyz7/gocialMedium/internal/server"

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
