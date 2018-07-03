package example

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Client is a simple demo wrapper
type Client struct {
	client    string
	machines  []string
	password  string
	separator string
}

// iterate through 'machines' trying to connect to each in turn.
// Assumes that 'machines' is non-empty
func tryConnect(machines []string, password string, timeout bool) (client, int, error) {
	var err error
	for _, address := range machines {
		conn := "example.org"
		var db int

		idx := strings.Index(address, "/")
		if idx != -1 {
			db, err = strconv.Atoi(address[idx+1:])
			if err == nil {
				address = address[:idx]
			}
		}

		network := "tcp"
		if _, err = os.Stat(address); err == nil {
			network = "unix"
		}
		fmt.Println("Trying to connect to the node %s", address)
	}
	return nil, 0, err
}
