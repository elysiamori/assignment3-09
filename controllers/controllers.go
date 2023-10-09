package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/elysiamori/assignment3-09/models"
	"github.com/elysiamori/assignment3-09/service"
	"gorm.io/gorm"
)

func UpdaterData(db *gorm.DB) {
	for {

		water := service.RandomizeData()
		wind := service.RandomizeData()

		dataToUpdate := models.Weather{
			Water: water,
			Wind:  wind,
		}
		apiURL := "http://localhost:3000/weather"
		payload, _ := json.Marshal(dataToUpdate)
		req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(payload))

		if err != nil {
			fmt.Println("Error Request:", err)
			time.Sleep(15 * time.Second)
			continue
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Errorrequest:", err)
			time.Sleep(15 * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("API returned non-OK status:", resp.Status)
			time.Sleep(15 * time.Second)
			continue
		}

		waterStat := service.UpdateStatusWater(water)
		windStat := service.UpdateStatusWind(wind)

		logData := models.Weather{
			Water: water,
			Wind:  wind,
		}

		logJSON, _ := json.MarshalIndent(logData, "", "  ")
		fmt.Printf("%s\nstatus water: %s\nstatus wind: %s\n", string(logJSON), waterStat, windStat)

		time.Sleep(15 * time.Second)
	}
}
