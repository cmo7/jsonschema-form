# jsonschema-form

## Desarrollo

### Sincronización de tipos

```bash
go install github.com/gzuidhof/tygo@latest
```

### Backend

Instalar Air para tener recargas automácias.

```bash
go install github.com/cosmtrek/air@latest
```

Ejecutar Air

```bash
cd back && air
```

### Frontend

```bash
cd front && pnpm dev
```

### TLS

Comando para certificado autofirmado

```bash
openssl.exe req -x509 -newkey rsa:4096 -nodes -out ./cert.pem -keyout ./key.pem -days 365
```

## Paquetes Backend

### App

Principal paquete de la aplicación. Todo lo referente al servidor y la aplicación se encuentra aquí.

#### Controllers

Funciones manejadoras de las rutas.

#### Middleware

Funciones middleware para las rutas.

#### Routes

Rutas de la aplicación. Están organizadas por módulos. Cada endpoint tiene su propio archivo en el que se genera un router con las rutas del endpoint.
por ejemplo: `routes/auth.go` contiene las rutas del endpoint de autenticación.

Las funciones de cada archivo de rutas devuelven un router que se monta en el router principal de la aplicación.

#### Types

### Certs

### Codegen

### Config

### Database

### Lib

### Models

### Public

### Services

### Tests
