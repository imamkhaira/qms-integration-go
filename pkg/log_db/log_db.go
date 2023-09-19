package logdb

import (
	"time"
)

type LogDbOptions struct {
	Host   string
	User   string
	Pass   string
	Dbname string
}

type LogDb struct {
	LogDbOptions
}

// create a new connection to the QMS internal database and abstraction
func New(options LogDbOptions) *LogDb {
	return &LogDb{}
}

// kontol.
func Kon() []RawLog {
	logs := []RawLog{}
	return logs
}

// get QMS log by time range (start to end)
func (l *LogDb) GetByTimeRange(begin time.Time, end time.Time) []RawLog {
	records := Kon()
	return records
}

// get full QMS log on that specific date from 00:00:00 to 23:59:59
func (l *LogDb) GetByDate(date time.Time) []RawLog {
	return l.GetByTimeRange(time.Now(), time.Now())
}

func (l *LogDb) GetYesterday() []RawLog {
	return l.GetByDate(time.Now())
}
