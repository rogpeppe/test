package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync/atomic"

	"cuelang.org/go/cue/load"
)

func main() {
	var userAgent atomic.Pointer[string]
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		agent := req.Header.Get("User-Agent")
		userAgent.Store(&agent)
	}))
	u, _ := url.Parse(srv.URL)
	os.Setenv("CUE_REGISTRY", u.Host)
	os.Setenv("CUE_EXPERIMENT", "modules")
	inst := load.Instances([]string{"."}, nil)[0]
	if agent := userAgent.Load(); agent != nil {
		fmt.Printf("user agent: %s\n", *agent)
	} else {
		fmt.Printf("error: %v\n", inst.Err)
	}
}
