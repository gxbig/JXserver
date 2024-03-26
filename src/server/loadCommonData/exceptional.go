package loadCommonData

import (
	"encoding/json"
)

// 接收
type requestResults struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}
