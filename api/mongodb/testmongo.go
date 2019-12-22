package zeit_qr_code

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestConnectMongo(w http.ResponseWriter, r *http.Request) {

	mongodb_uri := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create new client: %v", err), http.StatusBadRequest)

	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to connect to DB: %v", err), http.StatusBadRequest)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to ping DB: %v", err), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Successful DB connection\n"))

}
