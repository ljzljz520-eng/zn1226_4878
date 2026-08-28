package domain

type Event struct {
	Name, Subject string
	At            int64
}
type EventLog struct{ Events []Event }

func (l *EventLog) Add(e Event) { l.Events = append(l.Events, e) }
func (l EventLog) Names() []string {
	out := make([]string, 0, len(l.Events))
	for _, e := range l.Events {
		out = append(out, e.Name)
	}
	return out
}
