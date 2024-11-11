package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CatalogueItemType interface {
	GetType() string
}

type BuildingCatalogue struct {
	ID            primitive.ObjectID `json:"_d,omitempty" bson:"_id, omitempty"`
	AssetType     string             `json:"asset_type" validate:"required,max=50"`
	CatalogueName string             `json:"catalogue_name"`
	AssetName     string             `json:"asset_name"`
	BuildingType  string             `json:"building_type"`
	RoomNumber    string             `json:"room_number"`
	BuildingSize  string             `json:"building_size"`
	Location      string             `json:"location"`
}

func (b BuildingCatalogue) GetType() string {
	return "Building"
}

type VehicleCatalogue struct {
	ID                  primitive.ObjectID `json:"_d,omitempty" bson:"_id, omitempty"`
	Asset_type          string
	CatalogueName       string `json:"catalogue_name"`
	Vehicle_type        string
	Make                string
	Model               string
	FuelType            string
	Manufacture_country string
	UsefulLifeYears     string
	UsefulLifeMonths    string
}

func (v VehicleCatalogue) GetType() string {
	return "Building"
}

type EquipmentCatalogue struct {
	ID                    primitive.ObjectID `json:"_d,omitempty" bson:"_id, omitempty"`
	CatalogueName         string             `json:"catalogue_name"`
	EquipmentType         string
	EquipmentTypeCategory string
	AssetType             string
	Make                  string
	Model                 string
	UsefulLifeYears       string
	UsefulLifeMonths      string
	ManufactureCountry    string
}

func (b EquipmentCatalogue) GetType() string {
	return "Equipment"
}

type CatalogueSchema struct {
	AssetType             string
	AssetName             string
	Location              string
	BuildingType          string
	UsefulLifeYears       string
	UsefulLifeMonths      string
	ManufactureCountry    string
	Make                  string
	Model                 string
	EquipmentType         string
	EquipmentTypeCategory string
	Vehicle_type          string
	FuelType              string
	RoomNumber            string
	BuildingSize          string
}
