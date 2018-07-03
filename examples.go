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
// this may not work perfectly, learning some stuff here :)
func tryConnect(machines []string, password string, timeout bool) (int, int, error) {
	var err error
	for _, address := range machines {
		var db int

		idx := strings.Index(address, "/")
		if idx != -1 {
			// the db value is provided
			db, err = strconv.Atoi(address[idx+1:])
			if err == nil {
				address = address[:idx]
				fmt.Println(db) // \0/
			}
		}

		network := "tcp"
		if _, err = os.Stat(address); err == nil {
			network = "unix"
			fmt.Println(network) // \0/
		}
		fmt.Println("Trying to connect to the node %s", address)
	}
	return 0, 0, err
}
