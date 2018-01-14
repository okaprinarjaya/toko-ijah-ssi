package models

import (
    "database/sql"
    "errors"
    "mypkgcommon"
)

type Item struct {
    Sku string `json:"sku"`
    ItemName string `json:"item_name"`
    QuantityOnHand int `json:"quantity_on_hand"`
    Notes string `json:"notes"`
}

func (item *Item) generateTable(db *sql.DB)  {
    sqlTable := `
    CREATE TABLE IF NOT EXISTS items (
        sku VARCHAR(32) NOT NULL PRIMARY KEY
        item_name VARCHAR(128) NOT NULL
        quantity_on_hand NUMERIC 
        notes TEXT
    );
    `

    _, err := db.Exec(sqlTable)
    common.CheckErr(err)
}

func (item *Item) create(db *sql.DB) error {
    return errors.New("Not implemented")
}
