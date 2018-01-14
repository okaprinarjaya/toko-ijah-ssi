package test

import (
    "testing"
    "models"
    "mypkgcommon"
)

func TestCreateTable(t *testing.T) {
    conn := common.DbConnect();
    item := models.Item{}
    item.GenerateTable(conn)
    conn.Close()
}
