package main
 
import (
  "github.com/WRWPhillips/go-pic2text-site/internal/conf"
  "github.com/WRWPhillips/go-pic2text-site/internal/server"
  "github.com/WRWPhillips/go-pic2text-site/internal/cli"
)
func main() {
  cli.Parse()
  server.Start(conf.NewConfig())
}