package main

import (
	"github.com/gocql/gocql"
	"log"
)

// DBSession stores the current Cassandra session
type DBSession struct {
	Session *gocql.Session
}

// NewSession starts a new Cassandra session and stores it in the DBSession struct
// This function call should be followed by `defer db.Session.Close()` where appropriate
func NewSession(hosts []string, keyspace string, consistency gocql.Consistency) *DBSession {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = consistency

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	return &DBSession{
		Session: session,
	}
}

func (db *DBSession) getUserByUsername(username string) User {
	var usernameField, passwordField string
	if err := db.Session.Query(`SELECT username, password FROM users WHERE username=? LIMIT 1`, username).
		Scan(&usernameField, &passwordField); err != nil {
		log.Fatal(err)
	}

	return User{
		Username: usernameField,
		Password: passwordField,
	}
}
