package transactions

import (
    "net/http"
    "mypkgcommon"
    "models/transaction"
    "encoding/json"
)

func CreateTransaction(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    trx := transaction.Transaction{}
    err := decoder.Decode(&trx)
    common.CheckErr(err)

    conn := common.DbConnect()
    _, err = trx.Create(conn)
    common.CheckErr(err)
    conn.Close()

    if err != nil {
        common.RespondWithJSON(w, http.StatusInternalServerError, map[string] string { "error": err.Error() })
        return
    }

    common.RespondWithJSON(w, http.StatusCreated, map[string] string { "status": "created" })
}
