package main

import "github.com/miekg/dns"

func main() {
	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		
	})
}
