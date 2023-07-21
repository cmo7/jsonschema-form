# NGR-Stack (Backend)

Este proyecto es una API REST desarrollada con Go, Fiber, GORM y PostgreSQL.

Algunas de las características de este proyecto son:

-   Patrón de diseño MVC
-   Permite generar tipos de datos para el frontend
-   Permite generar JSON Schemas para los formularios
-   Generador de controladores CRUD
-   Generador de rutas CRUD
-   Autenticación con JWT
-   API REST
-   Internacionalización
-   Factories y Seeders
-   Migraciones automáticas
-   Agnosticismo de base de datos
-   Soporte para Docker (TODO)
-   Soporte para Docker Compose (TODO)
-   Soporte para TLS
-   Soporte para CORS
-   Soporte para Helmet
-   Soporte para Compresión
-   Soporte para Cache
-   Soporte para Logger
-   Soporte para Paginación
-   Soporte para Herramientas de Desarrollo
-   Soporte para Validación de Formularios

## Tecnologías y Dependencias

-   [Go](https://golang.org/)
-   [Fiber](https://gofiber.io/)
-   [GORM](https://gorm.io/)
-   [Docker](https://www.docker.com/)
-   [Docker Compose](https://docs.docker.com/compose/)

## Instalación

1. Clonar el repositorio

```bash
git clone XXX
```

2. Crear el archivo `.env` a partir del archivo `.env.example`

```bash
cp .env.example .env
```

3. Instalar las dependencias

```bash
go mod download
```

4. Iniciar el servidor

```bash
go run .
```

## Docker

TODO:

## Configuración

El archivo `.env` contiene las siguientes variables de entorno:

| Variable                       | Descripción                                 | Valor por defecto    |
| ------------------------------ | ------------------------------------------- | -------------------- |
| `ENVIROMENT`                   | Entorno de ejecución                        | `development`        |
| `APP_NAME`                     | Nombre de la aplicación                     | `Nartex Go App`      |
| `LOCALE`                       | Localización de la aplicación               | `ES`                 |
| `DB_ENGINE`                    | Motor de base de datos                      | `postgres`           |
| `DB_HOST`                      | Host de la base de datos                    | `localhost`          |
| `DB_PORT`                      | Puerto de la base de datos                  | `5432`               |
| `DB_USER`                      | Usuario de la base de datos                 | `postgres`           |
| `DB_PASSWORD`                  | Contraseña de la base de datos              | `postgres`           |
| `DB_DATABASE`                  | Nombre de la base de datos                  | `postgres`           |
| `MID_HELMET`                   | Helmet Middleware                           | `true`               |
| `MID_COMPRESS`                 | Compress Middleware                         | `true`               |
| `MID_CORS`                     | CORS Middleware                             | `true`               |
| `MID_LOGGER`                   | Logger Middleware                           | `true`               |
| `MID_CACHE`                    | Cache Middleware                            | `true`               |
| `GEN_FRONT_TYPES`              | Generar tipos de datos para el frontend     | `true`               |
| `GEN_FRONT_TYPES_PATH`         | Ruta de los tipos de datos para el frontend | `../front/src/types` |
| `GEN_AUTO_MIGRATE`             | Migrar automáticamente la base de datos     | `true`               |
| `CLIENT_MODE`                  | Modo cliente                                | `Internal`           |
| `CLIENT_URL`                   | URL del cliente                             | `http://localhost`   |
| `TLS`                          | Habilitar TLS                               | `true`               |
| `TLS_CERT`                     | Certificado TLS                             | `./certs/cert.pem`   |
| `TLS_KEY`                      | Llave TLS                                   | `./certs/key.pem`    |
| `TLS_PORT`                     | Puerto TLS                                  | `443`                |
| `JWT_ISSUER`                   | Emisor del JWT                              | `Nartex`             |
| `JWT_SECRET`                   | Secreto del JWT                             | `secret`             |
| `JWT_EXPIRATION`               | Expiración del JWT                          | `24h`                |
| `JWT_MAX_AGES`                 | Máxima edad del JWT                         | `24h`                |
| `LOG_ENABLED`                  | Habilitar logs                              | `true`               |
| `DB_QUERY_LOGGING`             | Habilitar logs de consultas                 | `true`               |
| `DEV_TOOLS`                    | Habilitar herramientas de desarrollo        | `false`              |
| `DATABASE_SEED``               | Habilitar seed de la base de datos          | `false`              |
| `PAGINATION_MAX_PAGE_SIZE`     | Tamaño máximo de página de paginación       | `100`                |
| `PAGINATION_DEFAULT_PAGE_SIZE` | Tamaño por defecto de página de paginación  | `10`                 |

## Paquete `app`

### Modelos

Los modelos se encuentran en `app/models` y se componen de dos structs:

-   Entidad: Representa una tabla de la base de datos.
-   DTO: Representa un objeto de transferencia de datos.
-   Formularios: Representa un formulario de entrada de datos.

#### Entidad

Las entidades son structs anotados con `gorm` de forma que el ORM pueda mapearlos a la base de datos. Estos structs deben implementar la interfaz `Entity` que contiene los métodos `GetID` y `ToDto`.

```go
type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	baseEntity
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Avatar   string `gorm:"type:bytes;not null"`
	Provider string `gorm:"type:varchar(255);not null;default:'local'"`
	// Relationships

	Roles []*Role `gorm:"many2many:user_roles;"`
	Posts []*Post
}

func (user *User) GetId() uuid.UUID {
	return user.ID
}

func (user *User) ToDto() UserDTO {
	filteredRoles := make([]RoleDto, len(user.Roles))
	for i, role := range user.Roles {
		filteredRoles[i] = role.ToDto().(RoleDto)
	}

	u := UserDTO{}
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.Avatar = user.Avatar
	u.Roles = filteredRoles
	u.Provider = user.Provider
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	return u
}
```

#### DTO

Los DTOs son structs que representan un objeto de transferencia de datos
Están anotados con `json` para poder ser serializados y deserializados.
Están anotados con `ts_type` y `ts_transform` para poder generar los tipos de datos para el frontend.
Estos structs deben implementar la interfaz `Dto` que contiene el método `GetId` y `Validate`.

```go
type UserDTO struct {
	ID uuid.UUID `json:"id,omitempty" ts_type:"string" ts_transform:"__VALUE__.toString()"`
	baseDTO
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	Roles    []RoleDto `json:"roles,omitempty" ts_type:"Role[]" ts_transform:"__VALUE__.map((role: Role) => role.Name.toString())"`
	Provider string    `json:"provider,omitempty"`
}

func (userDTO UserDTO) GetId() uuid.UUID {
	return userDTO.ID
}

func (userDTO UserDTO) Validate() []*ErrorResponse {
	return ValidateStruct(userDTO)
}

```

#### Formularios

Estos structs están anotados con json para poder serializarlos como JSON Schemas.
Están anotados con validate para poder ser validados.
Están anotados con title para tener label en el JSON Schema
Están anotados con widget para especificar que widget usar en el frontend
Están anotados con option para pasar dichas options al JSON UI Schema

```go
type SignUpInput struct {
	Name            string `json:"name" validate:"required" title:"Name"`
	Email           string `json:"email" validate:"required,email" title:"Email" widget:"email"`
	Password        string `json:"password" validate:"required,min=8" title:"Password" widget:"password"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8,eqfield=Password" title:"Password Confirm" widget:"password"`
	Avatar          string `json:"avatar" title:"Avatar" widget:"file" options:"accept:image/*,filePreview:false"`
}
```

### Controladores

Los controladores se encuentran en `app/controllers`

Distintos controladores pueden implementar distintos manejadores, por ejemplo `healthcheck.controller.go` contiene

```go
func HealthCheck(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Server is running",
	})
}
```

Un caso habitual de controlador es el controlador CRUD para un recurso dado. Para facilitar la creación de estos controladores tenemos un generador de controladores genéricos que se encuentra en `app/controllers/generic.controller.go` y se puede usar así:

```go
var RoleController = NewStandardCrudController[models.Role](
	CrudControllerOptions{
		ResourceName:       "Role",
		ResourceSlug:       "role",
		ResourcePluralName: "Roles",
		ResourcePluralSlug: "roles",
	},
)
```

El struct generado contiene los siguientes métodos:

```go
func (DefaultCrudController[Res]).Create() func(*fiber.Ctx) error
func (DefaultCrudController[Res]).Delete() func(*fiber.Ctx) error
func (DefaultCrudController[Res]).Get() func(*fiber.Ctx) error
func (DefaultCrudController[Res]).GetAll() func(*fiber.Ctx) error
func (DefaultCrudController[Res]).Update() func(*fiber.Ctx) error
func (DefaultCrudController[Res]).GetOptions() CrudControllerOptions
```

-   `Create()` retorna un manejador para crear un recurso
-   `Delete()` retorna un manejador para eliminar un recurso
-   `Get()` retorna un manejador para obtener un recurso
-   `GetAll()` retorna un manejador para obtener todos los recursos
-   `Update()` retorna un manejador para actualizar un recurso
-   `GetOptions()` retorna las opciones del controlador

### Rutas

Las rutas se encuentran en `app/routes`

Fiber sigue una sintaxis similar a Express para definir rutas:

```go
func ApiRoutes() *fiber.App {

	api := fiber.New()

	// Health Check
	api.Get("/healthcheck", controllers.HealthCheck).Name("Health Check")
	// Json Schemas for Forms Routes

	api.Mount("/auth", authRoutes())
	api.Mount("/administration", adminRoutes())
	api.Mount("/schema", schemaRoutes())
	api.Mount("/api-info", apiInfoRoutes())
	api.Mount("/user", userRoutes())
	api.Mount("/role", genericRoutes(controllers.RoleController))
	api.Mount("/analytic", genericRoutes(controllers.AnalyticsController))

	return api
}
```

Las funciones como authRoutes() retornan un router de Fiber que puede ser montado en otro router.

Para facilitar la creación de rutas tenemos un generador de rutas genéricas que se encuentra en `app/routes/generic.routes.go` y nos permite generar un router de Fiber a partir de un controlador, con las rutas CRUD para el recurso dado, siguiendo la convención de rutas públicas y protegidas standard.

Un ejemplo de una de las funcioens generadoras de router Fiber:

```go
func userRoutes() *fiber.App {
	router := fiber.New()
	controller := controllers.UserController

	// Public routes, no token required
	public := router.Group("/")
	public.Get("/", controller.GetAll).
		Name(i18n.S(i18n.GET_ALL, controller.ResourcePluralName))

	public.Get("/:id", controller.Get).
		Name(i18n.S(i18n.GET, controller.ResourceName))

	// Protected routes, token required, only admin
	protected := router.Group("/").
		Use(middleware.ValidateToken).
		Use(middleware.OnlyAdmin)

	protected.Post("/", controller.Create).
		Name(i18n.S(i18n.CREATE, controller.ResourceName))

	protected.Put("/:id", controller.Update).
		Name(i18n.S(i18n.UPDATE, controller.ResourceName))

	protected.Delete("/:id", controller.Delete).
		Name(i18n.S(i18n.DELETE, controller.ResourceName))

	return router
}
```

### Middlewares

Los middlewares se encuentran en `app/middlewares`

Los middlewares son funciones que se ejecutan antes de que se ejecute el manejador de una ruta. Estas funciones pueden modificar el contexto de la petición y retornar un error.

La función `Next()` del contexto de Fiber permite continuar con la ejecución de la ruta.

Esta forma de definir middlewares es similar a la de Express.

Un ejemplo de middleware:

```go
func DeserializeUser(c *fiber.Ctx) error {

	// Get claims from context
	claims := c.Locals("claims").(jwt.MapClaims)
	// Find user
	var user models.User
	database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	// If user is not present
	if user.ID.String() == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Set user in locals
	c.Locals("user", user.ToDto())

	// Continue
	return c.Next()
}
```

### Paquete `database`

El paquete `database` se encuentra en `app/database` y contiene la configuración de la base de datos.

Expone un objeto `DB` que contiene la instancia de la base de datos. Utiliza el paquete `gorm` para la conexión con la base de datos y como ORM.

El paquete también se encarga de migrar la base de datos y de cargar los seeders.

#### Factories

Las factories se encuentran en `app/database/factories` y se utilizan para generar datos de prueba. Utilizan el paquete `github.com/jaswdr/faker` para generar datos aleatorios.

Todas las factories deben implementar la interfaz `Factory` que contien los métodos `Create`, `CreateMany`, `CreateOneWithData` y `CreateManyWithData`.

```go
type Factory interface {
	createOne() interface{}
	createMany(int) []interface{}
	createOneWithData(interface{}) interface{}
	createManyWithData(int, interface{}) []interface{}
}
```

Un ejemplo las funciones createOne() y createMany() de factory:

```go
func (UserFactory) CreateOne() *models.User {
	user := models.User{
		Name:     Faker.Person().Name(),
		Email:    Faker.Internet().Email(),
		Password: Faker.Internet().Password(),
		Avatar:   "https://i.pravatar.cc/300?u=" + Faker.UUID().V4(),
		Provider: "local",
	}
	return &user
}

func (UserFactory) CreateMany(count int) []*models.User {
	var users []*models.User
	var user *models.User
	for i := 0; i < count; i++ {
		user = UserFactory{}.CreateOne()
		users = append(users, user)
	}
	return users
}
```

#### Seeders

Los seeders se encuentran en `app/database/seeders` y se utilizan para poblar la base de datos con datos de prueba.

Son scripts que usan las factories y son llamados al iniciar la aplicación si la variable de entorno `DATABASE_SEED` es `true`.

### Paquete `i18n`

El paquete `i18n` se encuentra en `app/i18n` y contiene las traducciones de la aplicación.

Expone la función `M` que retorna una traducción en el idioma actual.

```go
i18n.M(i18n.HEALTH_CHECK)
```

También expone la función `S` que retorna una traducción en el idioma actual, con los parámetros dados. Los parámetros se reemplazan en el orden en el que se pasan, con una sintaxis similar a la de `fmt.Sprintf`.

```go
i18n.S(i18n.FOUND, "User")
```
