package main

import "testing"

func TestCoucou(t *testing.T) {
	templateString := "COUCOU"
	templateMorse := "-.-. --- ..- -.-. --- ..-"

	morse := morse(templateString)
	if morse != templateMorse {
		t.Error("Fail to encode in Morse: " + morse)
	}

	demorse := demorse(templateMorse)
	if demorse != templateString {
		t.Error("Fail to decode from Morse: ")
	}
}
func TestComplex(t *testing.T) {
	templateString := "THIS IS LIKE A CHATEVER COMPLEX STRING, IS IT? IT IS!"
	templateMorse := "- .... .. ...  .. ...  .-.. .. -.- .  .-  -.-. .... .- - . ...- . .-.  -.-. --- -- .--. .-.. . -..-  ... - .-. .. -. --. .--.-. .. ...  .. - ..--..  .. -  .. ... -.-.--"

	morse := morse(templateString)
	if morse != templateMorse {
		t.Error("Fail to encode in Morse: " + morse)
	}

	demorse := demorse(templateMorse)
	if demorse != templateString {
		t.Error("Fail to decode from Morse: " + demorse)
	}
}
