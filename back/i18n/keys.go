package i18n

// MessageKey is a type for all message keys
type MessageKey int64

// Numeric enum, easy to extend.
const (
	REQUIRED MessageKey = iota
	NOT_FOUND
	INVALID
	ALREADY_EXISTS
	NOT_FOUND_OR_INVALID
	FOUND
	CREATED
	UPDATED
	NOT_DELETED
	DELETED
	GET_ALL
	GET
	CREATE
	UPDATE
	DELETE
	SERVER_RUNNING
	HEALTH_CHECK
)
