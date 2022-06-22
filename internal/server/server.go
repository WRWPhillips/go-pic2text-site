package server

import (
  "github.com/WRWPhillips/go-pic2text-site/internal/database"
  "github.com/WRWPhillips/go-pic2text-site/internal/store"
)
 
func Start() {
  store.SetDBConnection(database.NewDBOptions())
  router := setRouter()
 
  router.Run(":9000")
}