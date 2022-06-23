package main
 
import (
  "github.com/WRWPhillips/go-pic2text-site/internal/conf"
  "github.com/WRWPhillips/go-pic2text-site/internal/server"
)
func main() {
  server.Start(conf.NewConfig())
}