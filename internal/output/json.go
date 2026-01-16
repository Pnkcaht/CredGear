package output

import (
	"encoding/json"

	"cred-sanitizer/internal/model"
)

func JSON(c []model.Credential) ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
