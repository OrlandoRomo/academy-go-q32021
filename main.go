package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/api"
	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/router"
	"github.com/OrlandoRomo/academy-go-q32021/registry"
)

func main() {
	var (
		urbanDictionaryApiKey = flag.String("urban-dictionary-api-key", envString("URBAN_DICTIONARY_API_KEY", ""), "Urban Dictionary api key")
		urbanDictionaryPort   = flag.String("urban-dictionary-port", envString("URBAN_DICTIONARY_PORT", "8080"), "Urban Dictionary port to listen to")
	)
	flag.Parse()

	urbanDictionaryClient := api.NewUrbanDictionary(*urbanDictionaryApiKey)

	r := registry.NewRegistry(urbanDictionaryClient)

	router := router.NewRouter(r.NewAppController())
	log.Printf("listening on http://localhost:%s", *urbanDictionaryPort)
	http.ListenAndServe(fmt.Sprintf(":%s", *urbanDictionaryPort), router)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
