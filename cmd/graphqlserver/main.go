package main

import (
	"context"
	"fmt"
	"log"

	"github.com/voodoostack/fitstackapi/config"
	"github.com/voodoostack/fitstackapi/postgres"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("working")
}
