package zeitqrcode

import (
	"fmt"
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/sessions"
	"github.com/kidstuff/mongostore"

	"net/http"
)

// LoginHandler uses auth0 for auth
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch new store.
	dbsess, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	defer dbsess.Close()

	store := mongostore.NewMongoStore(dbsess.DB("test").C("test_session"), 3600, true,
		[]byte("secret-key"))

	// Get a session.
	session, err := store.Get(r, "session-key")
	if err != nil {
		log.Println(err.Error())
	}

	// Add a value.
	session.Values["foo"] = "bar"

	// Save.
	if err = sessions.Save(r, w); err != nil {
		log.Printf("Error saving session: %v", err)
	}

	fmt.Fprintln(w, "ok")
}
