package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/stream", streamFileHandler)

	http.ListenAndServe(":3000", r)
}

func streamFileHandler(w http.ResponseWriter, r *http.Request) {
	// Open the file
	file, err := os.Open("largefile.txt")
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set headers for streaming
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	// Create a flusher to flush the buffer
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Buffer to hold file chunks
	buf := make([]byte, 32) // 32 bytes buffer - this is small for testing purposes

	// Stream file in chunks
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				http.Error(w, "Error reading file", http.StatusInternalServerError)
			}
			break // Exit the loop if an error occurs or EOF is reached
		}
		if n == 0 {
			break // Exit the loop if no more data is read
		}

		// Write chunk to response
		if _, err := w.Write(buf[:n]); err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
		flusher.Flush() // Flush the buffer to the client

		// Simulate delay - note: would not be used in production code
		time.Sleep(1 * time.Second)
	}
}
