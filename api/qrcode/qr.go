package zeit_qr_code

import (
	// "fmt"
	"net/http"
	"strconv"

	qc "github.com/skip2/go-qrcode"
)

// Handler asd
func QrHandler(w http.ResponseWriter, r *http.Request) {
	var image []byte
	image, err := qc.Encode("https://example.org", qc.Medium, 256)

	if err != nil {
		_ = image
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(image)))
	if _, err := w.Write(image); err != nil {
		http.Error(w, "failed to write image", http.StatusPartialContent)
	}

	// if _, err := w.Write(buffer.Bytes()); err != nil {
	// log.Println("unable to write image.")
	// }
	// fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
