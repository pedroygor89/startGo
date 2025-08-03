package scheduler

import "time"

type Appointment struct {
	ID          string
	Client      string
	Artist      string
	Time        time.Time
	Description string
	Completed   bool
}

type Repository struct {
	appointments []Appointment
}

func (r *Repository) Add(a Appointment) {
	r.appointments = append(r.appointments, a)
}

func (r *Repository) ListByDate(date time.Time) []Appointment {
	var list []Appointment
	for _, a := range r.appointments {
		if sameDay(a.Time, date) {
			list = append(list, a)
		}
	}
	return list
}

func (r *Repository) MarkCompleted(id string) {
	for i, a := range r.appointments {
		if a.ID == id {
			r.appointments[i].Completed = true
			return
		}
	}
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
