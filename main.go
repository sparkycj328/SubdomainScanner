package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

// create a struct to hold the DNS server responses
type response struct {
	IPAddress string
	HostName  string
}

func main() {

	// Declare flag variables to be read in when running the program
	var (
		flDomain      = flag.String("domain", "", "The domain to be scanned")
		flWordList    = flag.String("wordlist", "", "The wordlist to be used for scanning")
		flWorkerCount = flag.Int("count", 100, "The amount of workers to be used for scanning")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server address to be used")
	)
	// parse the flag variables
	flag.Parse()

	// ensure that the non default variables contain a string
	if *flDomain == "" || *flWordList == "" {
		fmt.Println("-domain and -wordlist are required")
		os.Exit(1)
	}
	fmt.Println(flWorkerCount, " ", flServerAddr)
}

// queryARecords, as name implies, will be used to query for A records from the DNS server
func queryARecords(fqdn, serverAddr string) ([]string, error) {
	// Declare a new dns message
	var (
		m   dns.Msg
		ips []string
	)

	// Set the dns question for A type records
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
}
