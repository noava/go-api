package gardening

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/noava/go-api/db"
)

type PlantResponse struct {
	Name 			string
	StartDay  uint
	EndDay    uint
	Type 			string
}

// Get plants between two dates. If the start day is greater than the end day, it means the plant can be planted in the previous year.
func GetPlantsByDay(day uint) ([]PlantResponse, error) {
	var plants []PlantResponse
	result := db.DB.Model(&db.Plant{}).
									Select("name, start_day, end_day, type").
									Where("(? BETWEEN start_day AND end_day) OR (start_day > end_day AND (? >= start_day OR ? <= end_day))", 
											day, day, day).									
									Find(&plants)
	if result.Error != nil {
		return nil, result.Error
	}
	return plants, nil
}

func WhenToPlantHandler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	if date == "" {
		date = time.Now().Format("02-Jan")
	}
	
	// Convert date string to day of year uint
	t, err := time.Parse("02-Jan", date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	dayOfYear := uint(t.YearDay())

	plants, err := GetPlantsByDay(dayOfYear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Logging user requests
	log.Printf("Getting data for: %s", date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plants)
}

