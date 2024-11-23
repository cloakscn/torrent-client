package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloakscn/bit-torrent-client/torrentfile"
)

func main() {
	fmt.Println("Hello bit-torrent-client!")

	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}	
}
