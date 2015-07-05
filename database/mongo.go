package database

import (
	// "fmt"
	"github.com/shymega/shelves/models"
	log "gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

var (
	LogMongoDRV, LogMongoTX = models.ReturnDBLoggers()
)

// MongoConnect takes a hostname variable of string type, connects to the MongoDB server,
// and returns a type of *mgo.Session and a nil error value, if successful.
func MongoConnect(hostname string) (*mgo.Session, error) {
	session, err := mgo.Dial(hostname)
	if err != nil {
		LogMongoDRV.Crit("Error creating a session to MongoDB.", log.Ctx{"Hostname": hostname}, log.Ctx{"Error message": err})
		return session, err
	}
	defer session.Close()

	return session, nil
}
