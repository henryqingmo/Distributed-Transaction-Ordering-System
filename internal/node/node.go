package node

import (
	"bufio"
	"cs425_mp1/internal/config"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	identifier string
	parsed     config.Parsed
}

func New(identifier string, parsed config.Parsed) *Node {
	return &Node{
		identifier: identifier,
		parsed:     parsed,
	}
}

func node(identifier string, parsed config.Parsed) {

}

func (n *Node) Run() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 1024), 1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		// Expect: "<Action> <account> <amount>"
		if len(strings.Fields(line)) < 3 {
			log.Printf("skipping malformed line: %q", line)
			continue
		}
		fields := strings.Fields(line)

		action := fields[0]

		switch action {
		case "DEPOSIT":
			account := fields[1]
			amount := fields[2]
			fmt.Printf("Action: %s, Account: %s, Amount: %s\n", action, account, amount)
		case "TRANSFER":
			account1 := fields[1]
			account2 := fields[3]
			amount := fields[4]
			fmt.Printf("Action: %s, Account1: %s, Account2: %s, Amount: %s\n", action, account1, account2, amount)
		}
	}
}
