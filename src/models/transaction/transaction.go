package transaction

import (
    "database/sql"
    "mypkgcommon"
    "strings"
)

type Transaction struct {
    TrxId string `json:"trx_id"`
    TrxType string `json:"trx_type"`
    Notes string `json:"notes"`
    Created string `json:"created"`
    Updated string `json:"updated"`
    TrxDetail [] TransactionDetail `json:"trx_detail"`
}

type TransactionDetail struct {
    Id int `json:"id"`
    TrxId string `json:"trx_id"`
    Sku string `json:"sku"`
    Quantity int `json:"quantity"`
    Notes string `json:"notes"`
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
    
    CREATE TABLE IF NOT EXISTS transactions_details (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        trx_id VARCHAR(32) NOT NULL,
        sku VARCHAR(32) NOT NULL,
        quantity INTEGER NOT NULL,
        notes TEXT NULL
    );
    `
    _, err := db.Exec(sqlTable)
    common.CheckErr(err)
}

func (trx *Transaction) Create(db *sql.DB) (sql.Result, error) {
    stmt, err := db.Prepare("INSERT INTO transactions (trx_id, trx_type, notes) VALUES (?,?,?)")
    if err != nil {
        return nil, err
    }

    trxCode := trx.TrxType + "-" + common.CreateUniqueChars()
    _, err = stmt.Exec(trxCode, trx.TrxType, trx.Notes)

    if err != nil {
        return nil, err
    }

    result, err := createDetails(db, trxCode, trx.TrxDetail)

    if err != nil {
        return nil, err
    }

    return result, nil
}

func createDetails(db *sql.DB, trxCode string, trxDetail [] TransactionDetail) (sql.Result, error) {
    sqlStr := "INSERT INTO transactions_details (trx_id, sku, quantity, notes) VALUES "
    var values [] interface{}

    for _, trxItem := range trxDetail {
        sqlStr += "(?,?,?,?),"
        values = append(values, trxCode, trxItem.Sku, trxItem.Quantity, trxItem.Notes)
    }

    sqlStr = strings.TrimSuffix(sqlStr, ",")
    stmt, err := db.Prepare(sqlStr)

    if err != nil {
        return nil, err
    }

    result, err := stmt.Exec(values...)

    if err != nil {
        return nil, err
    }

    return result, nil
}

func (trx *Transaction) GetTransactions(db *sql.DB, trxType string) ([] Transaction, error)  {
    rows, err := db.Query(
        "SELECT trx_id, trx_type, notes, created, updated FROM transactions WHERE trx_type = $1",
        trxType)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var transactions [] Transaction
    for rows.Next() {
        var trx Transaction
        err := rows.Scan(&trx.TrxId, &trx.TrxType, &trx.Notes, &trx.Created, &trx.Updated)
        if err != nil {
            return nil, err
        }

        transactions = append(transactions, trx)
    }

    return transactions, nil
}

func (trxDetails *TransactionDetail) GetTransactionDetails(db *sql.DB, trxCode string) ([] TransactionDetail, error)  {
    rows, err := db.Query("SELECT id, trx_id, sku, quantity, notes FROM transactions_details " +
        "WHERE trx_id = $1",
        trxCode)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var trxDetailItems [] TransactionDetail
    for rows.Next() {
        var trxItem TransactionDetail
        err := rows.Scan(&trxItem.Id, &trxItem.TrxId, &trxItem.Sku, &trxItem.Quantity, &trxItem.Notes)
        if err != nil {
            return nil, err
        }

        trxDetailItems = append(trxDetailItems, trxItem)
    }

    return trxDetailItems, nil
}
