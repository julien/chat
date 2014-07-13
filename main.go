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

func homeHandler(w http.ResponseWriter, r *http.Request) {
  homeTpl.Execute(w, r.Host)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
  flag.Parse()

  homeTpl = template.Must(template.ParseFiles(filepath.Join(*assets, "client.html")))

  go h.Run()

  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/assets/", fileHandler)
  http.HandleFunc("/ws", wsHandler)

  log.Println("Starting server...")

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
