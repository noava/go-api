package pollen

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type PollenData struct {
	Hourly struct {
		AlderPollen  	[]float64 `json:"alder_pollen"`
		BirchPollen  	[]float64 `json:"birch_pollen"`
		GrassPollen  	[]float64 `json:"grass_pollen"`
		MugwortPollen []float64 `json:"mugwort_pollen"`
		OlivePollen  	[]float64 `json:"olive_pollen"`
		RagweedPollen []float64 `json:"ragweed_pollen"`
	} `json:"hourly"`
}

// Pollen level at highest severity time
func getHighestPollenLevel(pollenLevels []float64) float64 {
	if len(pollenLevels) == 0 {
		return 0
	}

	highest := pollenLevels[0]
	for _, level := range pollenLevels {
		if level > highest {
			highest = level
		}
	}
	return highest
}

// Classify pollen level
func classifyPollen(pollenLevels []float64) string {
	if len(pollenLevels) == 0 {
		return "No data"
	}

	worstLevel := getHighestPollenLevel(pollenLevels)

	switch {
	case worstLevel < 2:
		return "Low"
	case worstLevel < 6:
		return "Moderate"
	case worstLevel < 10:
		return "High"
	default:
		return "Very High"
	}
}


// Messages about pollen levels
func interpretPollenLevels(pollenLevels []float64) string {
	if len(pollenLevels) == 0 {
		return "Oops! No data available for this pollen type."
	}

	pollenAt16 := getHighestPollenLevel(pollenLevels)
	classification := classifyPollen([]float64{pollenAt16})

	switch classification {
	case "Low":
		return "It's a good day to be outside! Enjoy the fresh air."
	case "Moderate":
		return "Pollen levels are moderate. You may want to take precautions if you're sensitive."
	case "High":
		return "High levels of pollen detected. It's best to limit outdoor activities."
	case "Very High":
		return "Very high pollen levels. Consider staying indoors to avoid allergy symptoms."
	default:
		return "No data available on pollen levels."
	}
}



// Fetch and process pollen data
func SeverityHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	types := r.URL.Query().Get("types")

	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon parameters", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("https://air-quality-api.open-meteo.com/v1/air-quality?latitude=%s&longitude=%s", lat, lon)

	var pollenTypes []string
	if types != "" {
		// Add "_pollen" to each requested type
		for _, pollenType := range strings.Split(types, ",") {
			pollenTypes = append(pollenTypes, pollenType+"_pollen")
		}
	}

	// Default to all pollen types
	if len(pollenTypes) == 0 {
		pollenTypes = []string{"alder_pollen", "birch_pollen", "grass_pollen", "mugwort_pollen", "olive_pollen", "ragweed_pollen"}
	}

	pollenQuery := fmt.Sprintf("hourly=%s", strings.Join(pollenTypes, ","))

	// Completed API request URL
	apiURL = fmt.Sprintf("%s&%s", apiURL, pollenQuery)

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	var pollenData PollenData
	if err := json.Unmarshal(body, &pollenData); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}

	// Get all pollen severities for worst pollen level
	pollenSummary := map[string]string{
		"alder_pollen":  	classifyPollen(pollenData.Hourly.AlderPollen),
		"birch_pollen":  	classifyPollen(pollenData.Hourly.BirchPollen),
		"grass_pollen":  	classifyPollen(pollenData.Hourly.GrassPollen),
		"mugwort_pollen": classifyPollen(pollenData.Hourly.MugwortPollen),
		"olive_pollen":  	classifyPollen(pollenData.Hourly.OlivePollen),
		"ragweed_pollen": classifyPollen(pollenData.Hourly.RagweedPollen),
	}

	// Filter summary based on the requested pollen types
	filteredSummary := make(map[string]string)
	for _, pollenType := range pollenTypes {
		if val, exists := pollenSummary[pollenType]; exists {
			filteredSummary[pollenType] = val
		}
	}

	// Add interpretation for all requested pollen types
	interpretation := make(map[string]string)
	for _, pollenType := range pollenTypes {
		switch pollenType {
		case "alder_pollen":
			interpretation["alder"] = interpretPollenLevels(pollenData.Hourly.AlderPollen)
		case "birch_pollen":
			interpretation["birch"] = interpretPollenLevels(pollenData.Hourly.BirchPollen)
		case "grass_pollen":
			interpretation["grass"] = interpretPollenLevels(pollenData.Hourly.GrassPollen)
		case "mugwort_pollen":
			interpretation["mugwort"] = interpretPollenLevels(pollenData.Hourly.MugwortPollen)
		case "olive_pollen":
			interpretation["olive"] = interpretPollenLevels(pollenData.Hourly.OlivePollen)
		case "ragweed_pollen":
			interpretation["ragweed"] = interpretPollenLevels(pollenData.Hourly.RagweedPollen)
		}
	}

	output := struct {
		Summary      		map[string]string `json:"summary"`
		Interpretation 	map[string]string `json:"interpretation"`
	}{
		Summary:      	filteredSummary,
		Interpretation: interpretation,
	}

	// Logging user requests
	log.Printf("Getting data for: Lat %s & Long %s", lat, lon)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
