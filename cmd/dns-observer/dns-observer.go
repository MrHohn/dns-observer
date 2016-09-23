package main

import (
	"log"
	"strconv"
	"time"

	"github.com/miekg/dns"
	flag "github.com/spf13/pflag"
)

var (
	target   = flag.String("target", "kubernetes.default.svc.cluster.local", "Target for dns name resolution.")
	server   = flag.String("server", "10.0.0.10", "Dns server ip address.")
	period   = flag.Int("period", 250, "The period, in milliseconds, to execute a dns name resolution.")
	port     = flag.Int("port", 53, "Port number of the dns server.")
	protocol = flag.String("protocol", "udp", "Protocol to use (udp ot tcp), default to udp")
	exit     = flag.Bool("exit", false, "Exit when fail.")
)

func main() {

	flag.Parse()

	ticker := time.Tick(time.Duration(*period) * time.Millisecond)
	stopCh := make(chan struct{})

	for {
		client := dns.Client{Net: *protocol}
		msg := dns.Msg{}
		msg.SetQuestion(*target+".", dns.TypeA)

		select {
		case <-ticker:
			r, t, err := client.Exchange(&msg, *server+":"+strconv.Itoa(*port))
			if err != nil {
				log.Printf("err: %v\n", err)
				if *exit {
					return
				}
				continue
			}
			log.Printf("Resolution took %v", t)
			if len(r.Answer) == 0 {
				log.Printf("err: zero results\n")
				if *exit {
					return
				}
				continue
			}
			// for _, ans := range r.Answer {
			// 	Arecord := ans.(*dns.A)
			// 	log.Printf("%s", Arecord.A)
			// }
		case <-stopCh:
			return
		}
	}
}
