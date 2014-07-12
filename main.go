package main

import (
  "flag"
  "go/build"
  "log"
  "net/http"
  "path/filepath"
  "text/template"
  "os"
)

var (
  addr =   flag.String("addr", ":" + os.Getenv("PORT"), "http service address")
  assets = flag.String("assets", defaultAssetPath(), "path to assets")
  homeTpl  *template.Template
)

func defaultAssetPath() string {
  p, err := build.Default.Import("go_chat", "", build.FindOnly)

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

  homeTpl = template.Must(template.ParseFiles(filepath.Join(*assets, "client.html")))

  go h.run()

  http.HandleFunc("/",   homeHandler)
  http.HandleFunc("/ws", wsHandler)

  log.Println("Starting server...")

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
