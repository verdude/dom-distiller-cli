package main

import (
  "flag"
  "fmt"
  "time"

  distiller "github.com/markusmobius/go-domdistiller"
)

func main() {
  url := flag.String("url", "", "URL to distill")
  flag.Parse()

  if *url == "" {
    panic("URL is required")
  }

  // Start distiller
  result, err := distiller.ApplyForURL(*url, time.Minute, nil)
  if err != nil {
    panic(err)
  }

  fmt.Println("Title:", result.Text)
}
