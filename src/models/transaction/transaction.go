package transaction

import (
    "database/sql"
    "mypkgcommon"
)

type Transaction struct {
    TrxId string `json:"trx_id"`
    TrxType string `json:"trx_type"`
    Notes string `json:"notes"`
    Created string `json:"created"`
    Updated string `json:"updated"`
}

func (trx *Transaction) GenerateTable(db *sql.DB) {
    sqlTable := `
    CREATE TABLE IF NOT EXISTS transactions (
        trx_id VARCHAR(32) NOT NULL PRIMARY KEY,
        trx_type VARCHAR(6) NOT NULL,
        notes TEXT NULL,
        created DATETIME NOT NULL DEFAULT current_timestamp,
        updated DATETIME NOT NULL DEFAULT current_timestamp
    );
    `
    _, err := db.Exec(sqlTable)
    common.CheckErr(err)
}

func (trx *Transaction) Create(db *sql.DB, trxType string, notes string) (sql.Result, error) {
    stmt, err := db.Prepare("INSERT INTO transactions (trx_id, trx_type, notes) VALUES (?,?,?)")
    if err != nil {
        return nil, err
    }

    result, err := stmt.Exec("789", trxType, notes)
    if err != nil {
        return nil, err
    }

    return result, nil
}
