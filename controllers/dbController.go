package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"../datamodel"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var dbinstance *dbr.Connection

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
func getDBInstance() (*dbr.Connection, error) {
	var err error
	if dbinstance == nil {
		if !fileExists(datamodel.LOCALPATH) {
			fmt.Println("DB notfound,exit :(")
			os.Exit(-1)
		}
		dbinstance, err = dbr.Open(datamodel.DBTYPE, datamodel.LOCALPATH, nil)
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
	//	sess.SelectBySql("SELECT * FROM " + datamodel.TABLENAME + " WHERE id = (SELECT MAX(id) FROM " + datamodel.TABLENAME + ")").Load(&packet)

	sess.Select("*").From(datamodel.PACKET_TABLENAME).Where("id = (SELECT MAX(id) FROM " + datamodel.PACKET_TABLENAME + ")").Load(&packet)

	return c.JSON(http.StatusCreated, packet)
}

func InsertDB(c echo.Context) error {
	conn, err := getDBInstance()
	if err != nil {
		os.Exit(-1)
	}
	u := new(datamodel.DistPacket)
	if err := c.Bind(u); err != nil {
		return err
	}
	sess := conn.NewSession(nil)

	sess.InsertInto(datamodel.DISTANCE_TABLENAME).Columns("id", "mac", "pwr", "distance").Values(u.ID, u.MACaddr, u.Pwr, u.Distance).Exec()
	return c.NoContent(http.StatusOK)
}
