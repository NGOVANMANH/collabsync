package service

func GetHelloMessage(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
