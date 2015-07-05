package models

import (
	log15 "gopkg.in/inconshreveable/log15.v2"
)

// Database Loggers
var (
	// MongoDB
	LogMongoDRV = log15.New(log15.Ctx{"Module": "Database/MongoDB/Driver"}) // MongoDB Driver Logger
	LogMongoTX  = log15.New(log15.Ctx{"Module": "Database/MongoDB/TX"})     // MongoDB Transactions Logger (TX)
)

// DesktopGUI Loggers
var (
	LogDGUICore = log15.New(log15.Ctx{"Module": "DesktopGUI/Core"})
)

// Auth package Loggers
var (
	LogAuthFileConnect  = log15.New(log15.Ctx{"Module": "Auth/File/Connect"})
	LogAuthCryptGen     = log15.New(log15.Ctx{"Module": "Auth/Crypt/Gen"})
	LogAuthCryptCompare = log15.New(log15.Ctx{"Module": "Auth/Crypt/Compare"})
)

type Book struct {
	ID     int
	Title  string
	Author string
	ISBN   int
}

/** RESTful Server Handlers (Structures) **/

type IndexHandlerResponse struct {
	StatusCode int    `json:"statusCode"`
	Output     string `json:"output"`
	Doc        string `json:"doc"`
}

type ErrorHandlerResponse struct {
	StatusCode             int    `json:"statusCode"`
	BriefStatusCodeDescrip string `json:"briefStatusCodeDescrip"`
	Doc                    string `json:"doc"`
}

// ReturnDBLoggers returns two loggers, both of which are of 'log15.Logger' type for the
// MongoDB module.
func ReturnDBLoggers() (log15.Logger, log15.Logger) {
	return LogMongoDRV, LogMongoTX
}

// ReturnDGUILoggers returns one logger, which is of 'log15.Logger' type for the
// DesktopGUI module.
func ReturnDGUILoggers() log15.Logger {
	return LogDGUICore
}

// ReturnAuthLoggers returns three loggers, all of which are of 'log15.Logger' type for
// the auth package for Shelves.
func ReturnAuthLoggers() (log15.Logger, log15.Logger, log15.Logger) {
	return LogAuthFileConnect, LogAuthCryptGen, LogAuthCryptCompare
}
