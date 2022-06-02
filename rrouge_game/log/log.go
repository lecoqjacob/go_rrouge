package log

import (
	"fmt"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/mitchellh/colorstring"
)

func PrintTerminalError(msg string, a ...interface{}) string {
	full_msg := fmt.Sprintf("[red][ERROR]: %s", fmt.Sprintf(msg, a...))
	colorstring.Printf(full_msg)
	return full_msg
}

func PrintNotFoundComponentError(entityId uint64, componentType ecs.ComponentType) string {
	return PrintTerminalError(fmt.Sprintf("Entity %d has no [%s] component", entityId, componentType))
}

func Debug(msg string, a ...interface{}) string {
	full_msg := fmt.Sprintf("[yellow][DEBUG]: %s\n", fmt.Sprintf(msg, a...))
	colorstring.Printf(full_msg)
	return full_msg
}
