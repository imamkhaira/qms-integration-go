package logdb

type RawLog struct {
	id          int
	transcodeId string
	ticketNo    string
	keypad      int
	printTime   string
	callTime    string
	endTime     string
}
