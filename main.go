package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/router"
	"github.com/OrlandoRomo/academy-go-q32021/registry"
)

func main() {
	var UrbanDictionaryApiKey = flag.String("urban-dictionary-api-key", envString("URBAN_DICTIONARY_API_KEY", ""), "Urban Dictionary api key")
	flag.Parse()

	urbanDictionaryClient := api.NewUrbanDictionary(*UrbanDictionaryApiKey)

	r := registry.NewRegistry(urbanDictionaryClient)

	router := router.NewRouter(r.NewAppController())
	log.Printf("listening on http://localhost:%s", "8080")
	http.ListenAndServe(":8080", router)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
