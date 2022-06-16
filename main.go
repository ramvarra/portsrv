package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetExternalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ipAddr := conn.LocalAddr().String()
	return strings.Split(ipAddr, ":")[0]
}

func GetHostNameIP() (string, string) {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Fatal(err)
	}
	return name, strings.Join(addrs, ",")
}
type handler struct {
	Port string
}

var MyName, MyIP = GetHostNameIP()
var MyExtIP = GetExternalIP()

func (m *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("Hello: MyName: %s MyIP: %s ExtIP: %s Port %s Host: %s RemoteAddr %s", 
		MyName, MyIP, MyExtIP, m.Port, req.Host, req.RemoteAddr)
	fmt.Println(msg)
	fmt.Fprintln(resp, msg)
}

func isValidPort(port string) bool {
	n, err := strconv.Atoi(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad port '%s' - %v\n", port, err)
		return false
	}
	if n < 1 || n > 65535 {
		fmt.Fprintf(os.Stderr, "Bad Port. %s must be between 1, 65535\n", port)
		return false
	}
	return true
}

func startHTTPServer(port string, background bool) {
	fmt.Printf("Creating server on port '%s' - background %v\n", port, background)

	m := &handler{Port: port}
	if background {
		go func() {
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), m))
		}()
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), m))
	}
}
func main() {
	// fmt.Printf("ArgC: %d ARGV: %q ARGT: %T\n", len(os.Args), os.Args, os.Args)
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: portsrv port1 port2 ...\n")
		os.Exit(1)
	}
	log.Printf("MyName: %s MyIP: %s MyExtIP: %s Ports %v\n", MyName, MyIP, MyExtIP, os.Args[1:])

	for i, port := range os.Args[1:] {
		if !isValidPort(port) {
			os.Exit(1)
		}
		startHTTPServer(port, i < len(os.Args)-2)
	}

}
