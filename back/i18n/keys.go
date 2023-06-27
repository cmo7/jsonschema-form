package i18n

type MessageKey string

const (
	REQUIRED             MessageKey = "required"
	NOT_FOUND            MessageKey = "notFound"
	INVALID              MessageKey = "invalid"
	ALREADY_EXISTS       MessageKey = "alreadyExists"
	NOT_FOUND_OR_INVALID MessageKey = "notFoundOrInvalid"
	FOUND                MessageKey = "found"
	CREATED              MessageKey = "created"
	UPDATED              MessageKey = "updated"
	NOT_DELETED          MessageKey = "notDeleted"
	DELETED              MessageKey = "deleted"
	GET_ALL              MessageKey = "getAll"
	GET                  MessageKey = "get"
	CREATE               MessageKey = "create"
	UPDATE               MessageKey = "update"
	DELETE               MessageKey = "delete"
)
