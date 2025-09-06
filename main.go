package main

import (
	"DeSheans/test-customapp/pkg"
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type Config struct {
	RTP float64
}

func ParseFlags() *Config {
	cfg := &Config{}

	flag.Float64Var(&cfg.RTP, "rtp", -1.0, "percentage of prizes that will be returned to a player depending on funds deposited during the game initially")

	flag.Parse()

	if cfg.RTP == -1.0 {
		log.Println("rtp flag wasn's provided, set default value - 1.0")
		cfg.RTP = 1.0
	}
	if cfg.RTP <= 0 || cfg.RTP > 1 {
		log.Fatalf("Invalid RTP value: %.2f. Must be (0, 1]", cfg.RTP)
	}

	return cfg
}

func main() {
	cfg := ParseFlags()

	m := http.NewServeMux()

	m.HandleFunc("/get", getHandlerFunc(cfg.RTP))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}

func getHandlerFunc(rtp float64) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		mult := pkg.Mult(rtp)

		enc := json.NewEncoder(w)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = enc.Encode(struct {
			Result float64 `json:"result"`
		}{
			Result: mult,
		})
	}
}
