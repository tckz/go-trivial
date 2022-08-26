package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	// goのnamingだとsplit_wordsの区切り判定にそぐわないところはある
	type SomeConfig struct {
		// SOME_API_URL_DOMAIN
		ApiURLDomain string `required:"true" split_words:"true"`
		// SOME_APIURL_DOMAIN
		APIURLDomain string `required:"true" split_words:"true"`
	}

	type Config struct {
		Some SomeConfig `required:"true"`
	}

	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Printf("err=%v", err)
		return
	}

	fmt.Fprintf(os.Stderr, "%+v\n", c)
}
