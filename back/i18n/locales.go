package i18n

// Locale is a type for the supported locales
type Locale string

// String enum, so it can be loaded easily from config file.
const (
	ES Locale = "es"
	EN Locale = "en"
)
