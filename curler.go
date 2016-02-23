package curler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// New creates a new curler
func New(orig http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		parts := []string{"curl", "-v", "-X", r.Method}

		for k, v := range r.Header {
			parts = append(parts, "-H", fmt.Sprintf("%s: %s", k, strings.Join(v, ",")))
		}

		if r.Method == "POST" || r.Method == "PATCH" || r.Method == "PUT" {
			cts := strings.Split(r.Header.Get("Content-Type"), ";")
			var ct string
			if len(cts) > 0 {
				ct = cts[0]
			}
			if strings.HasSuffix(ct, "json") {
				b, _ := ioutil.ReadAll(r.Body)
				r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
				parts = append(parts, "-d", string(b))
			}
		}

		parts = append(parts, r.URL.String())
		fmt.Println(strings.Join(parts, " "))
		orig.ServeHTTP(rw, r)
	})
}
