// cmd/kernelnftprotocolnetworknext/main.go
package main

import (
"flag"
"log"
"os"

"kernelnftprotocolnetworknext/internal/kernelnftprotocolnetworknext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := kernelnftprotocolnetworknext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
