package handler

import (
	"encoding/json"
	"net/http"
	"log"
)

func Para(w http.ResponseWriter, r *http.Request) {
	// レスポンスの Content-Type を設定
	w.Header().Set("Content-Type", "application/json")

	// クエリパラメータをマップとして取得
	queryParams := r.URL.Query()

	// ログにクエリパラメータを出力
	log.Println("Recieved query parameters:", queryParams)

	// レスポンス用のマップを作成し、starttime と endtime は配列ではなく単一の値として扱う
	responseMap := make(map[string]interface{})
	for key, values := range queryParams {
		if key == "starttime" || key == "endtime" {
			responseMap[key] = values[0]
		} else {
			responseMap[key] = values
		}
	}

	if err := json.NewEncoder(w).Encode(queryParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
