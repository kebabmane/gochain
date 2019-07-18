package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kebabmane/gochain/chain"
	"github.com/kebabmane/gochain/consensus/pow"
	"github.com/kebabmane/gochain/rpc"
	"github.com/kebabmane/gochain/storage"
)

func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	db, err := storage.NewDb("./db/test")
	if err != nil {
		return
	}

	blockchain := chain.NewBlockChain(db)

	consensus := pow.NewPowServer(blockchain, nil)
	consensus.Start()

	rpcServer := rpc.NewServer(blockchain)
	go rpcServer.Start()

	//So we need to capture our application exiting so we can clean up some stuff...
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		// need to revisit this but I need to be able to save the current block with a prefix as the current block for when the app boots back up again
		// err := blockchain.SaveCurrentBlock(blockchain.CurrentBlock())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Wait for 2 second to finish processing")
		db.Close()
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()

	select {}

	// Close and commit the datbase
	//db.Close()
}
