package item

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
    Created string `json:"created"`
    Updated string `json:"updated"`
}

func (item *Item) GenerateTable(db *sql.DB) {
    sqlTable := `
    CREATE TABLE IF NOT EXISTS items (
        sku VARCHAR(32) NOT NULL PRIMARY KEY,
        item_name VARCHAR(128) NOT NULL,
        quantity_on_hand NUMERIC NOT NULL DEFAULT 0, 
        notes TEXT NULL,
        created DATETIME NOT NULL DEFAULT current_timestamp,
        updated DATETIME NOT NULL DEFAULT current_timestamp
    );
    `

    _, err := db.Exec(sqlTable)
    common.CheckErr(err)
}

func (item *Item) GetItems(db *sql.DB) ([] Item, error)  {
    rows, err := db.Query("SELECT sku, item_name, quantity_on_hand FROM items")
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var items [] Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.Sku, &item.ItemName, &item.QuantityOnHand); err != nil {
            return nil, err
        }
        items = append(items, item)
    }

    return items, nil
}

func (item *Item) Create(db *sql.DB) error {
    return errors.New("Not implemented")
}
