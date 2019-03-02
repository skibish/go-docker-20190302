package commands

import "fmt"

const (
	// Name is const for NAME cmd
	Name = "NAME"
	// Message is const for MSG cmd
	Message = "MSG"
	// Exit is const for EXIT cmd
	Exit = "EXIT"
)

// CreateNameCmd returns formatted string for NAME command
func CreateNameCmd(name string) string {
	return fmt.Sprintf("%s %s", Name, name)
}

// CreateMessageCmd returns formatted string for MSG command
func CreateMessageCmd(message string) string {
	return fmt.Sprintf("%s %s", Message, message)
}

// CreateExitCmd returns formatted string for EXIT command
func CreateExitCmd() string {
	return fmt.Sprintf("%s ", Exit)
}
