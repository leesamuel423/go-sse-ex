package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/events", events)
	http.ListenAndServe(*addr, nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	fmt.Fprintln(w, "hello from events")
	tokens := []string{"this", "is", "a", "test", "of", "live", "server", "event", "streaming", "..."}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// add wait time
		time.Sleep(time.Millisecond * 420)
	}
}

/* Sam Notes
- flag.String() takes three params
  - "addr": name of flag
  - ":8080": default value if no flag provided
  - "http service address": description of flag for help text
- returns a pointer to string (*string) that can be used to access flag

- flag.Parse() -> activate flags
- http.HandleFunc -> register function to handle HTTP request at specific path
- http.ListenAndServe -> start web server
  - first argument: address to listen on ... * for *addr because addr is pointer from flag.String()
  - second argument: router (nil uses default router)
- http.ResponseWriter -> how you send responses back to client
- http.Request -> contains info about incoming request

- w.Header().Set("Content-Type", "text/event-stream")
  - lets client know the type is an event-stream and should respond in real-time

- need to flush buffer earlier so that content is written immediately instead of accumulating in buffer and sending at once

*/
