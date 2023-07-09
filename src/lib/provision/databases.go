package provision

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func getDB() (*sql.DB, error) {
	connectionString := "root:password@tcp(127.0.0.1:3306)/"
	utils.DebugLog("Waiting for DB to be ready: " + connectionString)
	err := WaitForDB(connectionString)
	if err != nil {
		return nil, err
	}
	utils.DebugLog("DB is ready: " + connectionString)
	return sql.Open("mysql", connectionString)
}

type nullLogger struct {
}

func (d nullLogger) Print(args ...interface{}) {
}

func WaitForDB(conString string) error {
	mysql.SetLogger(nullLogger{})
	for {
		db, err := sql.Open("mysql", conString)
		if err != nil {
			return err
		}
		err = db.Ping()
		if err == nil {
			db.Close()
			break
		}
		time.Sleep(1 * time.Second)
		continue
	}
	return nil
}

func CreateDBAndUser(dbName string) error {
	utils.DebugLog("CreateDBAndUser started: " + dbName)
	err := CreateDatabase(dbName)
	if err != nil {
		utils.DebugLog("CreateDatabase error: " + err.Error())
		return err
	}
	err = CreateDBUserWithPrivileges(dbName)
	if err != nil {
		utils.DebugLog("CreateDBUserWithPrivileges error: " + err.Error())
		return err
	}
	utils.DebugLog("CreateDBAndUser completed: " + dbName)
	return nil
}

func CreateDatabase(dbName string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sDB, err := utils.GetSafeDBString(dbName)
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + sDB)
	if err != nil {
		return err
	}

	return nil
}

func CreateDBUserWithPrivileges(dbName string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sDB, err := utils.GetSafeDBString(dbName)
	if err != nil {
		return err
	}

	// ignore error, user may already exist
	db.Exec("CREATE USER '" + sDB + "'@'%' IDENTIFIED BY '" + sDB + "'")

	db.Exec("GRANT ALL PRIVILEGES ON " + sDB + ".* TO '" + sDB + "'@'%'")

	_, err = db.Exec("FLUSH PRIVILEGES")
	return err
}
