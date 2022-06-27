package server

import (
  "github.com/WRWPhillips/go-pic2text-site/internal/database"
  "github.com/WRWPhillips/go-pic2text-site/internal/store"
  "github.com/WRWPhillips/go-pic2text-site/internal/conf"
)
 
func Start(config conf.Config) {
  jwtSetup(config)
  
  store.SetDBConnection(database.NewDBOptions(config))
  
  router := setRouter()
 
  router.Run(":9000")
}