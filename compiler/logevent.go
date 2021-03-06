package compiler

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

type LogEvent struct {
	Timestamp float64
	Node      string
	Process   int64
	Thread    string
	Type      string
	Location  string
	Text      string
}

func (self *LogEvent) FromLine(line string) error {
	var err error
	// regex?
	/*
		n, _ := fmt.Sscanf(line, "%f %s %d %s %s %s %s\n",
			&self.Timestamp, &self.Node, &self.Process, &self.Thread,
			&self.Type, &self.Location, &self.Text)
		if n < 6 {
			fmt.Printf("Error parsing %s\n", line)
		}
	*/
	trimmed := strings.Trim(line, "\n")
	parts := strings.SplitN(trimmed, " ", 7)
	if len(parts) != 7 {return errors.New("Not enough fields")}

	//fmt.Printf("parts: %d %s\n", len(parts), parts)
	self.Timestamp, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {return err}
	self.Node = parts[1]
	self.Process, err = strconv.ParseInt(parts[2], 10, 32)
	if err != nil {return err}
	self.Thread = parts[3]
	self.Type = parts[4]
	self.Location = parts[5]
	self.Text = parts[6]
	return nil
}

func (self *LogEvent) ThreadID() string {
	return fmt.Sprintf("%s %d %s", self.Node, self.Process, self.Thread)
}

func (self *LogEvent) EventStr() string {
	return fmt.Sprintf("%s %s:%s", self.Location, self.Type, self.Text)
}

func (self *LogEvent) ToString() string {
	return self.ThreadID() + " " + self.EventStr()
}
