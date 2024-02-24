package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func logRequest(r *http.Request, err error) {
  log.Println(r.RemoteAddr, r.Method, r.URL.Path)
}


func main() {
    var address, prefix, webdavDir string

    flag.StringVar(&address, "address", "0.0.0.0:8080", "Address to listen to.")
    flag.StringVar(&prefix, "prefix", "", "Prefix is the URL path prefix to strip from WebDAV resource paths.")
    flag.StringVar(&webdavDir, "webdavDir", ".", "Path to serve")
    flag.Parse()

    handler := &webdav.Handler{
        Prefix:     prefix,
        FileSystem: webdav.Dir(webdavDir),
        LockSystem: webdav.NewMemLS(),
        Logger:     logRequest,
    }

    log.Println("serving", webdavDir, "on", address)
    http.ListenAndServe(address, handler)
}

