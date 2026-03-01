package main

import (
	"cs425_mp1/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "./mp1_node <identifier> <configuration file>")
		os.Exit(1)
	}

	//name := os.Args[1]
	filePath := os.Args[2]

	parsed, err := config.ParseConfig(filePath)
	if err != nil {
		log.Fatalf("parse config: %v", err)
	}

	for _, n := range parsed.Nodes {
		fmt.Printf("ID: %s, Host: %s, Port: %s\n", n.ID, n.Host, n.Port)
	}
	identifier := os.Args[1]

	node, err := config.ParseIdentifier(parsed, identifier)

	n := node.NewNode(node, parsed)
	n.Run()

}
