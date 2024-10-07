package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

type LoadBalancer struct {
	backends []*Backend
	current  uint64
}

func NewLoadBalancer(backends []string) *LoadBalancer {
	var servers []*Backend

	for _, backend := range backends {
		url, err := url.Parse(backend)
		if err != nil {
			log.Fatalf("Error parsing backend URL: %v", err)
		}
		servers = append(servers, &Backend{
			URL:          url,
			Alive:        true,
			ReverseProxy: httputil.NewSingleHostReverseProxy(url),
		})
	}

	return &LoadBalancer{
		backends: servers,
	}
}

func (lb *LoadBalancer) GetNextBackend() *Backend {
	next := atomic.AddUint64(&lb.current, 1)
	return lb.backends[next%uint64(len(lb.backends))]
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.GetNextBackend()
	log.Printf("Fowarding request to: %s", backend.URL.String())
	backend.ReverseProxy.ServeHTTP(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've hit the port on ....")
}


func main() {
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}
	lb := NewLoadBalancer(backends)
	log.Println("Load balancer started on port ....")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", lb)
}
