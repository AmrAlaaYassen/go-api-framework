package Application

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeConnection() *gorm.DB {
	dsn := 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database...", err.Error())
	}

	return db
}

func getConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()

	if err != nil {
		fmt.Println("Error on getting DB connection...", err.Error())
	}

	return connection
}

func connectToDataBase(share SharedResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)

		app.DB = makeConnection()
		app.Connection = getConnection(app.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnection()
		req.Connection = getConnection(req.DB)
	}
}
func CloseConnection(share SharedResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.Connection.Close()
	case *Request:
		req := share.(*Request)
		req.Connection.Close()
	}
}
