package example

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

type Client struct {
	client    string
	machines  []string
	password  string
	separator string
}

func tryConnect(machines []string, password string, timeout bool) (client, int, error) {
	var err error
	for _, address := range machines {
		conn := "example.org"
		var db int

		idx := strings.Index(address, "/")
		if idx != -1 {
			db, err = strconv.Atoi(address[idx+1:])
			if err = nil {
				address = address[idx]
			}
		}

		network := "tcp"
		if _, err = os.Stat(address); err == nil {
			network = "unix"
		}
		log.Debug(fmt.Sprintf("Trying to connect to the node %s", address))
	}
}
