package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/nazarnovak/tretton37/backend/api"
	"github.com/nazarnovak/tretton37/backend/pkg/fetcher"
)

func main() {
	// This could be moved into an environment variable or a config
	port := ":8080"

	srv := &http.Server{
		Addr:    port,
		Handler: router(),
	}

	fmt.Println("Running server on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(fmt.Errorf("Error starting server: %s", err.Error()))
	}
}

func router() *web.Mux {
	mux := web.New()

	mux.Get("/api/employees", api.EmployeesHandler(fetcher.Tretton37{}))

	return mux
}
