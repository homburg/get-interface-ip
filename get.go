package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"log"
	"net"
	"strings"
)

func main() {
	usage := `get-ip
	
	Get my ip by interface name.

	Usage:
	  get-ip <interface>

	`

	arguments, err := docopt.Parse(usage, nil, true, "get-ip 1.0", false)

	if nil != err {
		log.Fatal(err)
	}

	interfaceName := arguments["<interface>"].(string)

	ifi, err := net.InterfaceByName(interfaceName)

	if err != nil {
		log.Fatal(err)
	}

	addrs, err := ifi.Addrs()
	if err != nil {
		log.Fatal(err)
	}

	if len(addrs) == 0 {
		log.Fatalf("Could not get any addresses for %s\n", interfaceName)
	}

	firstAddress := addrs[0].String()
	firstAddress = strings.Split(firstAddress, "/")[0]

	fmt.Println(firstAddress)
}
