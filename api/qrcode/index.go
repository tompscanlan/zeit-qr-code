package zeitqrcode

import (
	"fmt"
	"net/http"
)

// Handler just says Hi
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
