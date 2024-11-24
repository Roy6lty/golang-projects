package pillar

import (
	"time"
)

type Logger struct {
	logTime time.Time
	logDate string
}

type ServiceUserHistoryLogger struct {
	Owner          ServiceUser
	OperationType  string
	LoggingMessage string
	Logger         Logger
}

func NewLogger() Logger {
	return Logger{
		logTime: time.Now(),
		logDate: time.Now().Format("2006-01-02"),
	}
}

// func (s *ServiceUserHistoryLogger) Logging() {
// 	//Logs All activities on service  user profile
// 	fmt.Printf("service user %v and  %v at %v and %v record created \n", s.Owner.First_name, s.Owner.Last_Name, s.logger.logDate, s.logger.logTime)

// }

type RecordUserHistoryLogger struct {
	RecordID string
	logger   Logger
}

func (s *RecordUserHistoryLogger) Logging() {
	//Logs All activities on service  user profile
	//fmt.Printf("service user %v and  %v at %v and %v record created \n", s.Owner.First_name, s.Owner.Last_Name, s.logger.logDate, s.logger.logTime)

}

type ServiceUserLogs struct {
	serviceUser_ID string
}

func (r *ServiceUserLogs) Logging(owner *ServiceUserLogs, l *Logger) {
	//Logs All activities on service  user profile
	//fmt.Printf("service user %v record %v at %v and %v \n", owner.serviceUser_ID, l.operationType, l.logDate, l.logTime)

}

type HCPLogs struct {
}

type RecordLogs struct {
	//logs changes in the record
	RecordID       string
	LoggingMessage string
}

func (r *RecordLogs) Logging(owner *RecordLogs, l *Logger) {
	//Logs All activities on service  user profile
	//fmt.Printf("%v record  %v at %v and %v \n", owner.RecordID, l.operationType, l.logDate, l.logTime)

}

type Logs struct {
}
