package pollen

import (
	"encoding/json"
	"log"
	"net/http"
)

type PollenInfo struct {
	WhatIs 					map[string]string `json:"what_pollen_is"`
	Symptoms 				[]string 					`json:"symptoms"`
	WhoIsAtRisk 		[]string 					`json:"who_is_at_risk"`
	ManagementTips 	[]string 					`json:"management_tips"`
	SeasonalPollen 	map[string]string `json:"seasonal_pollen"`
}

func PollenInfoHandler(w http.ResponseWriter, r *http.Request) {
	info := PollenInfo{
		// "What is" questions
		WhatIs: map[string]string{
			"Pollen": "A fine powder released by plants as part of their reproductive cycle.",
			"PollenAllergy": "A pollen allergy occurs when the immune system overreacts to pollen from trees, grasses, or weeds. This reaction leads to inflammation in the nasal passages, eyes, and sometimes lungs.",
		},

		// Symptoms of pollen allergies
		Symptoms: []string{
			"Sneezing",
			"Runny or stuffy nose",
			"Itchy or watery eyes",
			"Itchy throat or ears",
			"Coughing Harshly",
			"Fatigue (from poor sleep due to symptoms)",
			"Wheezing or shortness of breath (in severe cases)",
		},
		
		// Who is at risk for pollen allergies
		WhoIsAtRisk: []string{
			"Individuals with a family history of allergies",
			"Those with asthma or other allergic conditions",
			"Live in areas with high pollen counts",
			"Work outdoors or spend a lot of time outside",
		},

		// Management tips for pollen allergies
		ManagementTips: []string{
			"Check pollen forecasts daily and stay indoors on high pollen days.",
			"Keep windows closed and use air purifiers to reduce pollen indoors.",
			"Shower and change clothes after being outside to remove pollen.",
			"Wear sunglasses and a mask when going outside during high pollen seasons.",
			"Use antihistamines or prescribed allergy medications.",
		},

		// Common seasonal pollen types
		SeasonalPollen: map[string]string{
			"Spring": "Tree pollen (Alder, Birch, Maple, Oak)",
			"Summer": "Grass pollen",
			"Fall":   "Weed pollen (Ragweed, Mugwort)",
		},		
	}

	// Logging user requests
	log.Printf("Showing info about pollen")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
