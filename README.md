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
