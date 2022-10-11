package main

import "testing"

//test printer function
func TestPrinter(t *testing.T) {
	cases := []struct {
		name string
		lang string
		want string
	}{
		{"Lucien", "english", "Hello Lucien"},
		{"Algosup", "english", "Hello Algosup"},
		{"Lucien", "french", "Bonjour Lucien"},
		{"Lucien", "spanish", "Hola Lucien"},
	}
	for _, item := range cases {
		t.Run("greet "+item.name+" in "+item.lang, func(t *testing.T) {
			var a = greeter(item.name, item.lang)
			if a != item.want {
				t.Errorf("Expected %q, Got %q", item.want, a)
			}
		})
	}
}
