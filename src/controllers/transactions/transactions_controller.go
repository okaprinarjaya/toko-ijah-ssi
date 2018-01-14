package transactions

import (
    "net/http"
    "mypkgcommon"
    "models/transaction"
    "encoding/json"
    "github.com/gorilla/mux"
)

func CreateTransaction(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    trx := transaction.Transaction{}
    err := decoder.Decode(&trx)
    common.CheckErr(err)

    conn := common.DbConnect()
    _, err = trx.Create(conn)
    defer conn.Close()
    common.CheckErr(err)

    if err != nil {
        common.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    common.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "created"})
}

func ListTransactions(w http.ResponseWriter, req *http.Request) {
    trxType := req.URL.Query().Get("trxType")
    conn := common.DbConnect()
    trx := transaction.Transaction{}

    rows, err := trx.GetTransactions(conn, trxType)
    defer conn.Close()
    common.CheckErr(err)

    if err != nil {
        common.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    common.RespondWithJSON(w, http.StatusCreated, rows)
}

func ListTransactionDetails(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    conn := common.DbConnect()
    trx := transaction.TransactionDetail{}
    rows, err := trx.GetTransactionDetails(conn, vars["trxCode"])

    defer conn.Close()

    if err != nil {
        common.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    common.RespondWithJSON(w, http.StatusCreated, rows)
}
