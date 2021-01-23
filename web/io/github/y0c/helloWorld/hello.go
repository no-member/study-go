package main

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	return englishHelloPrefix + name
}
