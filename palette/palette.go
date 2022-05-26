package palette

import "github.com/anaseto/gruid"

// Color definitions. For now, we use a special color for FOV. We start from 1,
// because 0 is gruid.ColorDefault, which we use for default foreground and
// background.
const (
	ColorPlayer gruid.Color = 1 + iota // skip special zero value gruid.ColorDefault
	ColorMonster
	ColorFOV
	ColorDark
	ColorLogPlayerAttack
	ColorLogItemUse
	ColorLogMonsterAttack
	ColorLogSpecial
	ColorStatusHealthy
	ColorStatusWounded
	ColorConsumable
)

var Only8Colors bool

const (
	AttrReverse = 1 << iota
)
