package common

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/sessions"
	"github.com/kidstuff/mongostore"

	"golang.org/x/oauth2"

	oidc "github.com/coreos/go-oidc"
)

// GetSession from the mongoDB
func GetSession(r *http.Request) *sessions.Session {
	// Fetch new store.
	dbsess, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	// defer dbsess.Close()

	store := mongostore.NewMongoStore(dbsess.DB("test").C("test_session"), 3600, true,
		[]byte("secret-key"))

	// Get a session.
	session, err := store.Get(r, "session-key")
	if err != nil {
		log.Println(err.Error())

	}
	return session
}

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

// NewAuthenticator asd
func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, os.Getenv("AUTH_DOMAIN"))
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
