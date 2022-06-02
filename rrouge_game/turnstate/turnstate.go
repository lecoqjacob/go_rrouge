package turnstate

// This pkg is only so that I can
type TurnState int

const (
	PreRun TurnState = iota
	AwaitingInput
	PlayerTurn
	MonsterTurn
)

func (ts TurnState) String() string {
	return [...]string{"PreRun", "AwaitingInput", "PlayerTurn", "MonsterTurn"}[ts]
}

func (ts TurnState) NextState() TurnState {
	var nextState TurnState
	switch ts {
	case PreRun:
		nextState = AwaitingInput
	case AwaitingInput:
		nextState = PlayerTurn
	case PlayerTurn:
		nextState = MonsterTurn
	case MonsterTurn:
		nextState = AwaitingInput
	}

	return nextState
}
