package context

import (
	"fmt"
	"code.google.com/p/go-sqlite/go1/sqlite3"
)

/*
   #########################################################################
   # Event
   #########################################################################
*/
type Event struct {
	id             int
	ThreadID      int
	start_location string
	end_location   string
	StartTime     float64
	EndTime       float64
	StartType     string
	EndType       string
	start_text     string
	end_text       string
	count          int
	Depth          int
}

func (self *Event) NewEvent(query *sqlite3.Stmt) {
	query.Scan(
		&self.id,
		&self.ThreadID,
		&self.start_location, &self.end_location,
		&self.StartTime, &self.EndTime,
		&self.StartType, &self.EndType,
		&self.start_text, &self.end_text,
	)

	self.count = 1
	self.Depth = 0
}

func (self *Event) CanMerge(other Event, threshold float64) bool {
	return (other.Depth == self.Depth &&
		other.ThreadID == self.ThreadID &&
		other.StartTime-self.EndTime < 0.001 &&
		other.Length() < threshold &&
		other.start_text == self.start_text)
}

func (self *Event) Merge(other Event) {
	self.EndTime = other.EndTime
	self.count += 1
}

func (self *Event) Text() string {
	var text string

	if self.start_text == self.end_text || self.end_text == "" {
		text = self.start_text
	} else {
		text = self.start_text + "\n" + self.end_text
	}

	if self.count > 1 {
		text = fmt.Sprintf("%d x %s", self.count, text)
	}

	return text
}

func (self *Event) Length() float64 {
	return self.EndTime - self.StartTime
}
