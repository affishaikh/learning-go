package greetings

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := "Hi, " + name + ". Welcome!"
	return message
}
