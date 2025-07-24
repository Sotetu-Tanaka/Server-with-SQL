package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "id パラメータが必要です。", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "idは数字です", http.StatusBadRequest)
			return
		}

		row := db.QueryRow("SELECT id, name FROM users WHERE id =?", id)

		var uid int
		var name string
		err = row.Scan(&uid, &name)

		if err == sql.ErrNoRows {
			http.Error(w, "データが見つかりません", http.StatusNotFound)
			return
		} else if err != nil {
			fmt.Println(err)
			http.Error(w, "データベースエラー", http.StatusInternalServerError)
			return
		}

		user := map[string]interface{}{
			"id":   uid,
			"name": name,
		}

		w.Header().Set("Content_Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
