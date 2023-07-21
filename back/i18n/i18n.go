package i18n

import "fmt"

// Establishes the default active locale for the API
var active = EN

// Get message in active locale
func M(key MessageKey) string {
	return MessageDictionary[active][key]
}

// Get interpolated string in message in active locale
func S(key MessageKey, value ...string) string {
	return fmt.Sprintf(MessageDictionary[active][key], value)
}

// SetActive sets the active locale for the API
func SetActive(l Locale) {
	active = l
}
