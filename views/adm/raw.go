package adm

import (
	"ainosite/models"
	"encoding/json"
	"time"
)

func RawMood(moods []models.Mood) ([]byte, error) {
	dto := make(map[string]int)

	for _, mood := range moods {
		dto[mood.Date.Format(time.DateOnly)] = mood.Value
	}

	return json.Marshal(dto)
}
