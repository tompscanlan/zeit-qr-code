package zeitqrcode

import (
	"fmt"
	"log"

	"github.com/gorilla/sessions"

	"common"
	"net/http"
)

// LoginHandler uses auth0 for auth
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch new store.

	session := common.GetSession(r)

	// Add a value.
	session.Values["foo"] = "bar"

	// Save.
	if err := sessions.Save(r, w); err != nil {
		log.Printf("Error saving session: %v", err)
	}

	fmt.Fprintln(w, "ok")
}
