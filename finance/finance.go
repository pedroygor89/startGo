package finance

import "time"

type Record struct {
	Date   time.Time
	Artist string
	Client string
	Value  float64
	Method string
}

type Ledger struct {
	records []Record
}

func (l *Ledger) Add(r Record) {
	l.records = append(l.records, r)
}

func (l *Ledger) All() []Record {
	return append([]Record(nil), l.records...)
}
