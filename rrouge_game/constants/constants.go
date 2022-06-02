package constants

const GameName = "rrouge"

const (
	UI_WIDTH  = 80
	UI_HEIGHT = 24
	UI_COUNT  = UI_WIDTH * UI_HEIGHT
)

var (
	Xterm256Color = false
	Terminal      = false
	LogGame       = false
)

// type LogType string

// const (
// 	Info   LogType = "i"
// 	Warn   LogType = "w"
// 	Bad    LogType = "b"
// 	Danger LogType = "d"
// 	Good   LogType = "g"
// )

// var LogColors = map[LogType]color.Color{
// 	Info:   color.White,
// 	Warn:   palette.PColor(palette.Orange, 0.6),
// 	Bad:    palette.PColor(palette.Red, 0.6),
// 	Danger: palette.PColor(palette.Red, 0.3),
// 	Good:   palette.PColor(palette.Green, 0.6),
// }

type EquipSlot string

const (
	EquipHead   EquipSlot = "head"
	EquipWeapon EquipSlot = "weapon"
	EquipBoots  EquipSlot = "boot"
	EquipArmor  EquipSlot = "armor"
)

var EquipmentSlots = []EquipSlot{
	EquipHead,
	EquipWeapon,
	EquipBoots,
	EquipArmor,
}
