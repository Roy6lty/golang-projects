package pillar

import (
	"fmt"
	"time"
)

type Record interface {
	Create()
	Update()
	Delete()
	Read()
}

type MedicationRecord struct {
	RecordID        string
	MedicationName  string
	MedicationType  string
	MedicationRoute string
	Dosage          string
	Owner           ServiceUser
	Logger          ServiceUserHistoryLogger
}

func (r *MedicationRecord) Create(l *ServiceUserHistoryLogger) {
	// fmt.Printf("Creating new Medication Record for  %v s %v.", r.owner.First_name, r.owner.Last_Name)
	// fmt.Println("Record with id Created")
	fmt.Printf("Record with %v %v %v %v", r.MedicationName, r.MedicationRoute, r.MedicationType, r.Dosage)
	//service user Logs
	fmt.Printf("Logging record %v message %v", l.Owner.First_name, l.LoggingMessage)
}

func (r *MedicationRecord) Update() {
	fmt.Println("Updating  Medication Record.....")
}

func (r *MedicationRecord) Delete() {
	fmt.Println("Deleting  Medication Record.....")
}

func (r *MedicationRecord) Read() {
	fmt.Println("Reading  Medication Record.....")
}

type AllergyRecord struct {
	substance    string
	reactionType string
	severity     string
}

func (r *AllergyRecord) Create() {
	fmt.Println("Creating new Allergy Record.....")
}

func (r *AllergyRecord) Update() {
	fmt.Println("Updating  Allergy Record.....")
}

func (r *AllergyRecord) Delete() {
	fmt.Println("Deleting  Allergy Record.....")
}

func (r *AllergyRecord) Read() {
	fmt.Println("Reading  Allergy Record.....")
}

type ReferralRecord struct {
	substance    string
	reactionType string
	severity     string
}

func (r *ReferralRecord) Create() {
	fmt.Println("Creating new Referral Record.....")
}

func (r *ReferralRecord) Update() {
	fmt.Println("Updating  Referral Record.....")
}

func (r *ReferralRecord) Delete() {
	fmt.Println("Deleting  Referral Record.....")
}

func (r *ReferralRecord) Read() {
	fmt.Println("Reading  Referral Record.....")
}

type MessagesRecord struct {
	Sender   string
	receiver string
	timeSent time.Time
}

func (r *MessagesRecord) Create() {
	fmt.Println("Creating new message Record.....")
}

func (r *MessagesRecord) Update() {
	fmt.Println("Updating  message Record.....")
}

func (r *MessagesRecord) Delete() {
	fmt.Println("Deleting  message Record.....")
}

func (r *MessagesRecord) Read() {
	fmt.Println("Reading  message Record.....")
}
