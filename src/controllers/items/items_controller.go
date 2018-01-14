package items

import (
    "net/http"
    "mypkgcommon"
    "models/item"
)

func GetItems(w http.ResponseWriter, r *http.Request)  {
    conn := common.DbConnect()
    item := item.Item{}
    rowsItem, err := item.GetItems(conn)
    conn.Close()

    if err != nil {
        common.RespondWithJSON(w, http.StatusInternalServerError, map[string] string { "error": err.Error() })
        return
    }

    common.RespondWithJSON(w, http.StatusOK, rowsItem)
}