package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nithinps021/simplebank/api"
	db "github.com/nithinps021/simplebank/db/sqlc"
	"github.com/nithinps021/simplebank/util"
)

func main() {
	config,err:=util.LoadConfig(".")
	if err!=nil{
		log.Fatal("Cannot load config",err)
	}
	conn,err := sql.Open(config.DB_DRIVE, config.DB_SOURCE)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	store:=db.NewStore(conn)
	server:=api.NewServer(store)

	err=server.Start(config.SERVER_ADDRESS)
	if(err!=nil){
		log.Fatal("Cannot start server ", err)
	}

}
