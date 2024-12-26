package main

import (
	"flag"
	"fmt"

	metadataService "github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/chapter03/services/metadata"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "Port for metadata service")
	flag.Parse()
	fmt.Printf("Starting API at http://localhost:%d\n", port)
	metadataService.Run(port)
}
