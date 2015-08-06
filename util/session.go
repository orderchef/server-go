package util

import (
	"encoding/json"
)

type Session struct {
	UserID        *int    `json:"user_id"`
	AccessLevelID *string `json:"access_level_id"`
}

func DecodeSession(sess string) (Session, error) {
	var session Session
	err := json.Unmarshal([]byte(sess), &session)

	return session, err
}
