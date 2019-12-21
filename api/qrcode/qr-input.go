package zeit_qr_code

import (
	// "fmt"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	qc "github.com/skip2/go-qrcode"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		vals :=	r.URL.Query()

		if vals.Get("q") == "" {
			http.Error(w, "Must pass 'q' param", http.StatusBadRequest)
		}

		var image []byte
		image, err := qc.Encode(vals.Get("q"), qc.Medium, 256)

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(image)))
		if _, err := w.Write(image); err != nil {
			http.Error(w, "failed to write image", http.StatusPartialContent)
		}
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
func QrHandler(w http.ResponseWriter, r *http.Request) {

	// if _, err := w.Write(buffer.Bytes()); err != nil {
	// log.Println("unable to write image.")
	// }
	// fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}