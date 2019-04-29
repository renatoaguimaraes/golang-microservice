package util

import (
	"log"

	"github.com/gocql/gocql"
)

var session *gocql.Session

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "users"
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		log.Fatal(err)
	}
}

// Session return cassandra connection
func Session() *gocql.Session {
	return session
}

// Close cassandra connection
func Close() {
	session.Close()
}
