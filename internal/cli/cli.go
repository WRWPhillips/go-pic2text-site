package cli
 
import (
  "flag"
  "fmt"
  "os"

  "github.com/WRWPhillips/go-pic2text-site/internal/logging"
)
 
func usage() {
  fmt.Print(`This program runs go-pic2text-site backend server.
 
Usage:
 
pic2text [arguments]
 
Supported arguments:
 
`)
  flag.PrintDefaults()
  os.Exit(1)
}
 
func Parse() {
  flag.Usage = usage
  env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
  flag.Parse()
  logging.ConfigureLogger(*env)
  if *env == "prod" {
    logging.SetGinLogToFile()
  }
}