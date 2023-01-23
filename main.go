package main

import (
	"database/sql"
	"log"
	"github.com/nithinps021/simplebank/api"
	db "github.com/nithinps021/simplebank/db/sqlc"
	_ "github.com/lib/pq"

)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable"
	serverAddress ="0.0.0.0:8080"
)

func main() {
	conn,err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	store:=db.NewStore(conn)
	server:=api.NewServer(store)

	err=server.Start(serverAddress)
	if(err!=nil){
		log.Fatal("Cannot start server ", err)
	}

}