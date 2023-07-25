package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/count", readiness)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	log.Println("server listening on", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server: %v", err)
	}
}

func readiness(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-REQUEST-ID")
	log.Println("start", requestID)
	defer log.Println("done", requestID)

	time.Sleep(5 * time.Second)

	response := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		panic(err)
	}
}

// graceful shutdown

// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	router := http.NewServeMux()

// 	router.HandleFunc("/count", readiness)

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: router,
// 	}

// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, syscall.SIGTERM)

// 	go func() {
// 		log.Println("server listening on", server.Addr)
// 		if err := server.ListenAndServe(); err != http.ErrServerClosed {
// 			log.Fatalf("server: %v", err)
// 		}
// 	}()

// 	log.Printf("server is running on :%s", server.Addr)

// 	<-stop

// 	log.Printf("server shutting down")

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	if err := server.Shutdown(ctx); err != nil {
// 		log.Fatalf("server Shutdown Failed:%+v", err)
// 	}

// 	log.Printf("process clean up...")
// }

// func readiness(w http.ResponseWriter, r *http.Request) {
// 	requestID := r.Header.Get("X-REQUEST-ID")

// 	log.Println("start", requestID)
// 	defer log.Println("done", requestID)

// 	time.Sleep(5 * time.Second)

// 	response := struct {
// 		Status string `json:"status"`
// 	}{
// 		Status: "OK",
// 	}

// 	err := json.NewEncoder(w).Encode(&response)
// 	if err != nil {
// 		panic(err)
// 	}
// }
