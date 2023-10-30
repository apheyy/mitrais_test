package config

import (
	"encoding/json"
	"middle-developer-test/common/constants"
	"middle-developer-test/model"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, resp *model.ApiResponse) {
	w.Header().Set("Content-Type", constants.APP_JSON)
	json.NewEncoder(w).Encode(resp)
}
