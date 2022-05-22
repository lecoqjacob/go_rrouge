package main

type logStyle int

const (
	logNormal logStyle = iota
	logCritic
	logNotable
	logDamage
	logSpecial
	logStatusEnd
	logError
	logConfirm
)

type logEntry struct {
	Text  string
	MText string
	Index int
	Tick  bool
	Style logStyle
	Dups  int
}
