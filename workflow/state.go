package workflow

type State string

const (
	StateNew       State = "new"
	StateValidated State = "validated"
	StateCancelled State = "cancelled"
	StateCommitted State = "committed"
)

func NextState(s State, event string) State {
	switch s {
	case StateNew:
		if event == "validate" {
			return StateValidated
		}
	case StateValidated:
		if event == "cancel" {
			return StateCancelled
		}
		if event == "commit" {
			return StateCommitted
		}
	}
	return s
}
