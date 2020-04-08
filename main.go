package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	log.Println("Starting server :8080")

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,

		// setting custom timeouts
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

// resolveHosIP function
func resolveHostIP() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIP, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {

			ip := networkIP.IP.String()

			fmt.Println("Resolved Host IP: " + ip)

			return ip
		}
	}
	return ""
}

// Greet Function
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from "+resolveHostIP())
}
