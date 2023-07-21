# NGR Stack Seed

Sencillo repositorio para iniciar un proyecto de aplicación web FullStack con Go y React.

Información específica del backend en [backend/README.md](./back/README.md)
Información específica del frontend en [frontend/README.md](./front/README.md)

Se intentará mantener documentación actualizada y exahustiva del proyecto para facilitar el uso del stack a desarrolladores que no estén familiarizados con Go o React.

## Información general y decisiones de diseño

1. El backend está desarrollado en Go y el frontend en React.
2. El backend está desarrollado con el framework [Fiber](https://gofiber.io/) y el ORM [GORM](https://gorm.io/)
3. El frontend está desarrollado con [Vite](https://vitejs.dev/).
4. La estructura del proyecto está fuertemente influenciada por [Laravel](https://laravel.com/).

### Los objetivos que se persiguen son:

1. **Rapidez de desarrollo**: Se busca que el desarrollador pueda enfocarse en la lógica de negocio y no en la configuración de la aplicación.
2. **Escalabilidad**: Se busca que la aplicación pueda escalar de forma sencilla y sin problemas.
3. **Mantenibilidad**: Se busca que el código sea fácil de mantener y de entender.
4. **Testeabilidad**: Se busca que el código sea fácil de testear.
5. **Facilidad de despliegue**: Se busca que el despliegue de la aplicación sea lo más sencillo posible.

### Decisiones de diseño:

- **Uso extensivo de linting y formateadores**: Se busca que el código sea consistente y fácil de leer. Se usan formateadores y linters para asegurar que el código cumple con los estándares de la comunidad. Todos los archivos de código deben ser formateados y pasar los linters antes de ser subidos al repositorio.
- **Uso de Docker** (TODO): Se busca que el despliegue de la aplicación sea lo más sencillo posible. Se usa Docker para crear imágenes de la aplicación y desplegarlas en cualquier máquina que tenga Docker instalado.
- **Entorno de desarrollo con [Visual Studio Code](https://code.visualstudio.com/)**: Se busca que el desarrollador pueda enfocarse en la lógica de negocio y no en la configuración de la aplicación. Se usa VSCode por su gran cantidad de herramientas para Go y React. El repositorio además incluye sugerencias de extensiones para VSCode, configuración inicial y tareas de automatización para ejecutar el backend y el frontend en modo desarrollo.

### Principales diferencias con un proyecto en Java y Spring

- **Composición en lugar de herencia**: Evitamos crear jerarquías de tipos. La composición es más flexible y permite una mejor escalabilidad.
- **Convención sobre configuración**: Minimizamos la configuración. Siempre que sea posible, se usan los valores por defecto.
- **Poca Magia**: Intentamos que el código sea lo más sencillo y directo posible. No hay magia. Si la hay es poca y está bien documentada.
- **Herramientas**: Podemos usar cualquier editor de texto, pero VSCode cuenta con una gran cantidad de herramientas para Go y React.
- **Builds reproduciibles**: El proyecto cuenta con un Dockerfile para el backend y otro para el frontend. Esto permite que el proyecto pueda ser construido y ejecutado en cualquier máquina que tenga Docker instalado. (TODO)

## Hoja de Ruta

Actualmente hay una versión inicial del proyecto que permite crear usuarios y autenticarse. Otra funcionalidad de negocio dependerá del proyecto concreto que se esté desarrollando.
Aun así este proyecto semilla tiene pendientes algunas funcionalidades interesantes:

- [ ] **Tests**: Se debe añadir tests unitarios y de integración.
- [ ] **Documentación**: Se debe añadir documentación de la API y de la base de datos.
- [ ] **Docker**: Se debe añadir un Dockerfile para el backend y otro para el frontend.
- [ ] **CI/CD**: Se debe añadir un pipeline de CI/CD para el proyecto.
- [ ] **Ejemplo de uso**: Se deben crear varios proyectos de aplicación que usen este stack y documentarlos.

## Contribuciones

Si quieres contribuir al proyecto, puedes hacerlo de varias formas:

- **Reportando errores**: Si encuentras algún error en el código o en la documentación, puedes abrir un issue en el repositorio.
- **Añadiendo funcionalidad**: Si quieres añadir alguna funcionalidad al proyecto, puedes abrir un pull request en el repositorio.
- **Añadiendo documentación**: Si quieres añadir documentación al proyecto, puedes abrir un pull request en el repositorio.
