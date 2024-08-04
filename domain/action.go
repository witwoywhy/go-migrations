package domain

type Action string

const (
	Up   Action = "up"
	Down Action = "down"
)

func IsNotAction(action Action) bool {
	switch action {
	case Up, Down:
		return false
	default:
		return true
	}
}
