package i18n

// Message is a map of MessageKey and string, used to store the messages for a locale
type Message = map[MessageKey]string

// Dictionary is a map of Locale and Message, used to store the messages each locale
type Dictionary = map[Locale]Message

var MessageDictionary = Dictionary{
	ES: {
		REQUIRED:             "%s es requerido",
		NOT_FOUND:            "%s no encontrado",
		INVALID:              "%s es inválido",
		ALREADY_EXISTS:       "%s ya existe",
		NOT_FOUND_OR_INVALID: "%s no encontrado o inválido",
		FOUND:                "%s encontrado",
		CREATED:              "%s creado",
		UPDATED:              "%s actualizado",
		NOT_DELETED:          "%s no eliminado",
		DELETED:              "%s eliminado",
		GET_ALL:              "obtener todos los %s",
		GET:                  "obtener %s",
		CREATE:               "crear %s",
		UPDATE:               "actualizar %s",
		DELETE:               "eliminar %s",
		SERVER_RUNNING:       "El servidor se está ejecutando",
		HEALTH_CHECK:         "El servidor se está ejecutando",
	},
	EN: {
		REQUIRED:             "%s is required",
		NOT_FOUND:            "%s not found",
		INVALID:              "%s is invalid",
		ALREADY_EXISTS:       "%s already exists",
		NOT_FOUND_OR_INVALID: "%s not found or invalid",
		FOUND:                "%s found",
		CREATED:              "%s created",
		UPDATED:              "%s updated",
		NOT_DELETED:          "%s not deleted",
		DELETED:              "%s deleted",
		GET_ALL:              "get all %s",
		GET:                  "get %s",
		CREATE:               "create %s",
		UPDATE:               "update %s",
		DELETE:               "delete %s",
		SERVER_RUNNING:       "Server is running",
		HEALTH_CHECK:         "Server is running",
	},
}
