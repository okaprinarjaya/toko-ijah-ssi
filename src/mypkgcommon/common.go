package common

import (
    "database/sql"
    "net/http"
    "encoding/json"
)

func DbConnect() *sql.DB {
    db, err := sql.Open("sqlite3", "./tokoijah.db")
    CheckErr(err)

    return db
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    resp, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(resp)
}