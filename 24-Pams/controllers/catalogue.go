package controllers

import (
	"encoding/json"
	"net/http"
	model "olowoleru/pams/models"

	"github.com/cmatthias/mapstructure"
)

func CreateCatalogue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var raw map[string]interface{}               // a dict of any data type
	decoder := json.NewDecoder(r.Body)           // decoding the body parameter of teh object
	if err := decoder.Decode(&raw); err != nil { // modifying the raw data with a pointer by passing the address location
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data_catalogue interface{}
	if raw["catalogue_name"] == "vehicle" {
		var vehicle model.VehicleCatalogue
		if err := mapstructure.Decode(raw, &vehicle); err != nil { // if an error is returned
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data_catalogue = vehicle
	}

	// switch raw["catalogue_name"] {
	// case "vehicle":
	// 	if err := mapstructure.Decode(raw, &vehicle); err != nil { // if an error is returned
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	data_catalogue = vehicle

	// case "building":
	// 	var building model.BuildingCatalogue
	// 	if err := mapstructure.Decode(raw, &building); err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	data_catalogue = building

	// case "equipment":
	// 	var equipment model.VehicleCatalogue
	// 	if err := mapstructure.Decode(raw, &equipment); err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	data_catalogue = equipment

	// default:
	// 	http.Error(w, "Unknown catalogue type", http.StatusBadRequest)
	// 	return
	// }

}
