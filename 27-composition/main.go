package main

import "github.com.roy6lty/compostion/pillar"

type StartAble interface {
	Start()
}

func startEngines(cars ...StartAble) { // a variadic function which accepts multiple values of type startAble
	//the values are passed as an interface and can be looped through
	for _, car := range cars {
		car.Start()
	}
}

func main() {
	myConvertible := Convertable{Engine{}, EnhancedTransmission{}, SteeringWheel{}}
	myTruck := Truck{Engine{}, Transmission{}, SteeringWheel{}}
	startEngines(myConvertible, myTruck)

	myConvertible.Start()
	myConvertible.CoverTop()
	myTruck.Start()
	myTruck.FourWheelDrive()

	user1 := pillar.ServiceUser{First_name: "Ayobami", Last_Name: "Olowoleru", Age: 27, Weight: 71, Height: 172}
	//user2 := pillar.ServiceUser{First_name: "Fola", Last_Name: "Olowoleru", Age: 24, Weight: 65, Height: 176}
	//base_logger := pillar.NewLogger()

	su_logger := pillar.ServiceUserHistoryLogger{Owner: user1, OperationType: "Creation", LoggingMessage: "Creating Record", Logger: pillar.NewLogger()}

	//create a new record

	medication := pillar.MedicationRecord{RecordID: "123", MedicationName: "paracetamol", MedicationType: "NSAIDS", Dosage: "20mg", Owner: user1, Logger: su_logger}

	medication.Create(&su_logger)

}
