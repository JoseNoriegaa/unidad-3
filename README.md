# unidad-3
> Simple API Rest en Gin-Gonic y GORM.

## API
### Autor
| Método | Código | Ruta | Datos | Description |
|--|--|--|--|--|
| GET | 200 | /api/v1/authors/ |  | Retorna todos los autores registrados |
| GET | 200, 404 | /api/v1/authors/:fullname | **param**: fullname | Retorna la primer coincidencia de autor con el nombre proporcionado. |
| POST | 201, 400 | /api/v1/authors/ | **body**: fullname | Registra un nuevo autor. |
| PUT | 200, 404, 400 | /api/v1/authors/:fullname | **body**: fullname | Actualiza la primer coincidencia de autor con el nombre proporcionado. |
| DELETE | 200, 404 | /api/v1/authors/:fullname | **param**: fullname | Borra la primer coincidencia de autor con el nombre proporcionado. |

### Libro
| Método | Código | Ruta | Datos | Description |
|--|--|--|--|--|
| GET | 200 | /api/v1/books/ |  | Retorna todos los libros registrados |
| GET | 200, 404 | /api/v1/books/:id | **param**: id | Retorna un libro por el id proporcionado. |
| POST | 201, 400 | /api/v1/books/ | **body**: title, description, author, editorial, publicationDate | Registra un nuevo libro. |
| PUT | 200, 404, 400 | /api/v1/books/:id | **body**: title, description, author, editorial, publicationDate | Actualiza un libro por el id proporcionado. |
| DELETE | 200, 404 | /api/v1/books/:id | **param**: id | Borra un libro por el id proporcionado. |

## Instalación
### Docker
> Si tienes docker instalado en tu estación de trabajo sigue estas instrucciones, de lo contrario, ve a la sección de instalación manual. 
1. Clona el repositorio.
2. Abre una terminal en el directorio raíz del proyecto.
3. Ejecuta el siguiente comando para construir las imagenes locales de docker:
    ```bash
    $ docker-compose build
    ```
4. Ejecuta el siguiente comando para ejecutar la aplicación.
    > Nota: Es importante que el puerto 3306 y 3000 de tu estación de tu estación de trabajo estén libres.
5. Abre algún cliente de MySQL en el servidor local, y ejecuta el script `db.sql`.

### Manual
1. Obten el proyecto localmente con el siguiente comando:
    ```bash
    $ go get github.com/josenoriegaa/unidad-3
    ```
2. Abre una terminal en el directorio raíz del proyecto.
3. Instala las siguientes dependencias:
    ```bash
    $ go get -u github.com/gin-gonic/gin
    $ go get -u github.com/jinzhu/gorm
    $ go get -u github.com/jinzhu/gorm/dialects/mysql
    ```
4. Ejecuta el script sql `db.sql` en tu administrador de bases de datos de MySQL.
5. Abre el archivo `main.go` en tu editor de texto y reemplaza variables de entorno (línea 21) por las credenciales de tu base de datos.
    > Ejemplo:
    ```go
    DB_USER := "root"
    DB_PWD := ""
    DB_NAME := "db_test"
    DB_HOST := "127.0.0.1"
    ```
6. Corre el proyecto localmente con el siguiente comando:
    ```bash
    $ go run main.go
    ```
