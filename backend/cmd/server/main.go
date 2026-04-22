package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status" : "ok}`)
	})
	slog.Info("server starting", "port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
