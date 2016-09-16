package main

import (
	"log"
	"strconv"
	"time"

	"github.com/miekg/dns"
	flag "github.com/spf13/pflag"
)

var (
	target       = flag.String("target", "google.com", "target for dns name resolution.")
	server       = flag.String("server", "127.0.1.1", "dns server ip address.")
	lookupPeriod = flag.Int("period", 500, "The time, in milliseconds, to execute a dns name resolution.")
	port         = flag.Int("port", 53, "Port number of the dns server.")
)

func main() {

	flag.Parse()

	client := dns.Client{}
	msg := dns.Msg{}
	msg.SetQuestion(*target+".", dns.TypeA)

	ticker := time.Tick(time.Duration(*lookupPeriod) * time.Millisecond)
	stopCh := make(chan struct{})

	for {
		select {
		case <-ticker:
			r, t, err := client.Exchange(&msg, *server+":"+strconv.Itoa(*port))
			if err != nil {
				log.Printf("err: %v\n", err)
				continue
			}
			log.Printf("Resolution took %v", t)
			if len(r.Answer) == 0 {
				log.Printf("err: zero results\n")
				continue
			}
			// for _, ans := range r.Answer {
			//  Arecord := ans.(*dns.A)
			//  log.Printf("%s", Arecord.A)
			// }
		case <-stopCh:
			return
		}
	}
}
