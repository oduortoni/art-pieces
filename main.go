// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	fmt.Printf("Server listening on %s\n", port)
// 	http.ListenAndServe(":9000", nil)
// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {

// 	}
// }

package main

import (
	"fmt"

	log_t "github.com/oduortoni/art-pieces/log"
)

func main() {
	// Example of writing logs
	log_t.LogW("ALL", "Starting the application", nil)
	log_t.LogW("Main", "Something went wrong", fmt.Errorf("an example error"))

	// Example of reading logs for a specific function
	logs := log_t.LogR("MainFunction")
	for _, log := range logs {
		fmt.Println(log)
	}

	// Example of reading all logs
	allLogs := log_t.LogR("Main")
	for _, log := range allLogs {
		fmt.Println(log)
	}
}
