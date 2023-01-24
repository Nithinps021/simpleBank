package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/nithinps021/simplebank/util"
)

var testQueries *Queries
var db *sql.DB


func TestMain(m *testing.M) {
	var err error
	config,err:= util.LoadConfig("../../")
	if err!=nil{
		log.Fatal("Not able to load config file")
	}
	db, err = sql.Open(config.DB_DRIVE,config.DB_SOURCE)
	if err!=nil {
		log.Fatal("Cannot connect to db ",err)
	}
	testQueries = New(db)
	os.Exit(m.Run())
}
