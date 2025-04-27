package util

import (
	"context"
	"encoding/json"
)

func GetId(r context.Context) int32 {
	id := r.Value("userId")
	idInt64, err := id.(json.Number).Int64()
	if err != nil {
		return 0
	}
	return int32(idInt64)
}
