package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/undiabler/golang-whois"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		target string
		err    error
	)

	// prior command-line arguments
	if len(os.Args) > 1 {
		target = os.Args[1]
	} else {
		// secondly get from clipboard
		target, err = clipboard.ReadAll()
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	// if target is ip address set domain-name to target.
	if isIp(target) {
		target = getDomain(target)
	}

	whoIs(target)
}

func isIp(target string) bool {
	// This is not Ip-address validation
	r := regexp.MustCompile(`([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})`)
	return r.MatchString(target)
}

func getDomain(ipAddr string) string {
	result, err := net.LookupAddr(ipAddr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return result[0]
}

func whoIs(domain string) {
	// Eliminate last dot
	length := len(domain)
	if (domain[length-1 : length]) == "." {
		domain = domain[0 : length-1]
	}

	result, err := whois.GetWhois(domain)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("Your request %s result is ...\n", domain)
	fmt.Printf("%s", result)

	if len(strings.Split(domain, ".")) <= 2 {
		return
	}

	var input string
	fmt.Print("Are you satisfy?[y/n] ")
	fmt.Scanln(&input)

	if input == "n" {
		whoIs(domain[strings.Index(domain, ".")+1 : len(domain)])
	}

}
