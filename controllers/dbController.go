package controllers

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"../datamodel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

var dbinstance *dbr.Connection

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
func getDBInstance() (*dbr.Connection, error) {
	var err error
	if dbinstance == nil {
		conn, err := sql.Open(datamodel.DBTYPE, datamodel.PATH+datamodel.DATABASE_NAME)
		if err != nil {
			fmt.Errorf("mysql: could not get a connection: %v", err)
			os.Exit(1)
		}
		defer conn.Close()

		// Check the connection.
		if conn.Ping() == driver.ErrBadConn {
			fmt.Errorf("mysql: could not connect to the database. " +
				"could be bad address, or this address is not whitelisted for access.")
			os.Exit(1)

		}

		dbinstance, err = dbr.Open(datamodel.DBTYPE, datamodel.PATH+datamodel.DATABASE_NAME, nil)
		if err != nil {
			fmt.Println("Success!DBConnection!")
		}
	}

	return dbinstance, err
}

//rootにアクセスした時に呼ばれる。id=1だけ呼び出すお試しのもの
func ConnectDB(c echo.Context) error {
	conn, err := getDBInstance()
	if err != nil {
		os.Exit(-1)
	}
	sess := conn.NewSession(nil)

	var dbpacket datamodel.DBPacket
	sess.Select("*").From(datamodel.PACKET_TABLENAME).Where("id = ?", 1).Load(&dbpacket)

	return c.JSON(http.StatusCreated, dbpacket)
}

//id=?に応じたのを返す
func SelectPacketData(c echo.Context) error {
	conn, err := getDBInstance()
	if err != nil {
		os.Exit(-1)
	}
	sess := conn.NewSession(nil)

	id, _ := strconv.Atoi(c.QueryParam("id"))

	var packet datamodel.DBPacket
	sess.Select("*").From(datamodel.PACKET_TABLENAME).Where("id = ?", id).Load(&packet)
	return c.JSON(http.StatusCreated, packet)
}

//最新のidの情報を返す
func NewPacketData(c echo.Context) error {
	conn, err := getDBInstance()
	if err != nil {
		os.Exit(-1)
	}
	sess := conn.NewSession(nil)

	var packet datamodel.DBPacket

	sess.Select("*").From(datamodel.PACKET_TABLENAME).Where("id = (SELECT MAX(id) FROM " + datamodel.PACKET_TABLENAME + ")").Load(&packet)

	return c.JSON(http.StatusCreated, packet)
}
