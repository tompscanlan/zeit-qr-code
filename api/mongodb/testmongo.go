package zeitqrcode

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ShortLink represents the record that will be stored
type ShortLink struct {
	Name string
	Age  int
	City string
}

func mongoConnect() (*mongo.Client, error) {

	mongodbURI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("unable to create new client: %v", err)
	}

	return client, nil
}

// TestConnectMongo pings the mongodb server at the env var
// $MONGODB_URI and responds with the result
func TestConnectMongo(w http.ResponseWriter, r *http.Request) {
	client, err := mongoConnect()
	if err != nil {
		http.Error(w, fmt.Sprintf("mongoConnect error: %v", err), http.StatusBadRequest)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to ping DB: %v", err), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<blink>Successful DB connection</blink>\n"))
}
