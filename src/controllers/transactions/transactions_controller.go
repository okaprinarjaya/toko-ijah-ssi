package transactions

import (
    "net/http"
    "mypkgcommon"
    "models/transaction"
    "encoding/json"
    "fmt"
)

const TRX_TYPE_IN string = "TRX-IN"
const TRX_TYPE_OUT string = "TRX-OUT"

func CreateTransaction(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    trx := transaction.Transaction{}
    err := decoder.Decode(&trx)
    common.CheckErr(err)

    fmt.Println(trx)
    fmt.Println(trx.TrxId)
    fmt.Println(trx.TrxType)
    fmt.Println(trx.Notes)

    conn := common.DbConnect()
    _, err = trx.Create(conn, TRX_TYPE_IN, "This is a notes for this TRX IN")
    conn.Close()

    if err != nil {
        common.RespondWithJSON(w, http.StatusInternalServerError, map[string] string { "error": err.Error() })
        return
    }

    common.RespondWithJSON(w, http.StatusCreated, map[string] string { "status": "created" })
}
