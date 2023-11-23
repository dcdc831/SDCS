package main

import (
	"SDCS/httpagent"
	"os"
)

func main() {
	args := os.Args[1:]

	agent := httpagent.NewHttpAgent(0, args[0], args[1])
	agent.StartHttpAgent()
}
