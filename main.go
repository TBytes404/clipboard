package main

import (
	"log"
	"os"

	"github.com/tbytes404/clipboard/db"
	"github.com/tbytes404/clipboard/handler"
	"github.com/tbytes404/clipboard/router"
	"github.com/tbytes404/clipboard/store"
)

func main() {
	addr := os.Getenv("CLPBRD_ADDR")
	if addr == "" {
		addr = ":8282"
	}

	dsn := os.Getenv("CLPBRD_DSN")
	if dsn == "" {
		dsn = "./db/test.db"
	}

	db, err := db.NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	bs := store.NewBlobsStore(db)
	bh := handler.NewBlobsHandler(bs)

	log.Fatal(router.Init(bh).Start(addr))
}
