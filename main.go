package main

import (
  "flag"
  "go/build"
  "log"
  "net/http"
  "path/filepath"
  "text/template"
)

var (
  addr =   flag.String("addr", ":8080", "http service address")
  assets = flag.String("assets", defaultAssetPath(), "path to assets")
  homeTpl  *template.Template
)

func defaultAssetPath() string {
  p, err := build.Default.Import("websocket_chat", "", build.FindOnly)

  if err != nil {
    return "."
  }
  return p.Dir
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
  homeTpl.Execute(c, req.Host)
}

func main() {
  flag.Parse()

  homeTpl = template.Must(template.ParseFiles(filepath.Join(*assets, "home.html")))

  go h.run()

  http.HandleFunc("/",   homeHandler)
  http.HandleFunc("/ws", wsHandler)

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }

}
