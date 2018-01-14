package common

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/speps/go-hashids"
    "net/http"
    "time"
    "encoding/json"
    "strconv"
    "crypto/md5"
    "encoding/hex"
    "strings"
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

func GetTimestamp() int64 {
    now := time.Now()
    millis := now.UnixNano() / 1000000

    return millis
}

func CreateUniqueChars() string {
    millisStr := strconv.FormatInt(GetTimestamp(), 10)
    last3Char := millisStr[len(millisStr)-3:]
    one, _ := strconv.Atoi(last3Char[0:0])
    two, _ := strconv.Atoi(last3Char[0:1])
    three, _ := strconv.Atoi(last3Char[0:3])

    hasher := md5.New()
    hasher.Write([] byte(millisStr))

    hd := hashids.NewData()
    hd.Salt = hex.EncodeToString(hasher.Sum(nil))
    h, _ := hashids.NewWithData(hd)
    id, _ := h.Encode([] int { one, two, three })

    return strings.ToUpper(id)
}