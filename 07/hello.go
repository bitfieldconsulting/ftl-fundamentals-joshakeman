package hello

func ReturnGreeting(s string) string {
	if s == "Hi there" {
		return "Hi there yourself!"
	} else {
		return "I don't recognize what you're saying"
	}
}
