package database
 
import (
  "github.com/WRWPhillips/go-pic2text-site/internal/conf"

  "github.com/go-pg/pg/v10"
)
 
func NewDBOptions(config conf.Config) *pg.Options {
  return &pg.Options{
    Addr:     config.DbHost + ":" + config.DbPort,
    Database: config.DbName,
    User:     config.DbUser,
    Password: config.DbPassword,
  }
}