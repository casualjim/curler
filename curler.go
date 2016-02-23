package curler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// New creates a new curler
func New(orig http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		parts := []string{"curl", "-v", "-X", r.Method}

		for k, v := range r.Header {
			ck := http.CanonicalHeaderKey(k)
			if ck != "Host" && ck != "User-Agent" {
				parts = append(parts, "-H", fmt.Sprintf("'%s: %s'", k, strings.Join(v, ",")))
			}
		}

		if r.Method == "POST" || r.Method == "PATCH" || r.Method == "PUT" {
			b, err := ioutil.ReadAll(r.Body)

			if err != nil {
				log.Println("curler:", err)
			}
			if err == nil && len(b) > 0 {
				r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
				parts = append(parts, "-d '", string(b)+"'")
			}
		}

		parts = append(parts, r.URL.String())
		fmt.Println(strings.Join(parts, " "))
		orig.ServeHTTP(rw, r)
	})
}
