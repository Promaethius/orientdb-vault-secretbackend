package main

import (
	"log"
	"os"

	"github.com/hashicorp/vault/helper/pluginutil"
	"github.com/Promaethius/orientdb-vault-secretbackend"
)

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	err := orientdb.Run(apiClientMeta.GetTLSConfig())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
