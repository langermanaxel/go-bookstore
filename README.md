# Go Bookstore

Este es un proyecto de API RESTful para una aplicación de gestión de librerías construida en Go (Golang) con conexión a una base de datos MongoDB. El objetivo del proyecto es ofrecer operaciones CRUD (Crear, Leer, Actualizar, Eliminar) para gestionar libros en una librería.

## 📋 Índice

1. [Descripción](#-descripción)
2. [Características](#-características)
3. [Requisitos Previos](#-requisitos-previos)
4. [Instalación](#-instalación)
5. [Uso](#-uso)
6. [Endpoints](#-endpoints)
7. [Contribuir](#-contribuir)
8. [Licencia](#-licencia)
9. [Contacto](#-contacto)

## 📖 Descripción

El proyecto `go-bookstore` es una API REST que permite gestionar un inventario de libros. Cada libro tiene información como título, autor, año de publicación y número de páginas. La API permite realizar operaciones CRUD sobre los libros almacenados en una base de datos MongoDB.

## ✨ Características

- CRUD completo para la gestión de libros.
- Conexión con base de datos MongoDB.
- Estructura modular y limpia usando las mejores prácticas en Go.
- Gestión de errores adecuada.
- Configuración fácil y rápida.

## 📦 Requisitos Previos

Asegúrate de tener instalados los siguientes elementos en tu entorno:

- [Go](https://golang.org/doc/install) (versión 1.18 o superior)
- [MongoDB](https://www.mongodb.com/try/download/community)
- Git

## ⚙️ Instalación

Sigue estos pasos para instalar y ejecutar el proyecto localmente:

1. **Clonar el repositorio:**
   
   ```bash
   git clone https://github.com/langermanaxel/go-bookstore.git

2. **Ir al directorio del proyecto:**

   ```bash
   cd go-bookstore

3. **Instalar las dependencias:**

   ```bash
   go mod tidy

4. **Configurar la conexión a la base de datos:**
   - Crea un archivo .env en la raíz del proyecto y añade la URL de conexión a MongoDB:

   MONGODB_URI=mongodb://localhost:27017

5. **Ejecutar el servidor:**

   go run main.go

   - El servidor debería estar corriendo en http://localhost:8000.

## 🚀 Uso

Para interactuar con la API, puedes usar herramientas como Postman o curl. A continuación, se presentan algunos ejemplos de cómo utilizar la API.

### Crear un libro

curl -X POST http://localhost:8000/books -H "Content-Type: application/json" -d '{
    "title": "El Principito",
    "author": "Antoine de Saint-Exupéry",
    "year": 1943,
    "pages": 96
}'

### Obtener todos los libros

curl http://localhost:8000/books

### Actualizar un libro

curl -X PUT http://localhost:8000/books/{id} -H "Content-Type: application/json" -d '{
    "title": "El Principito - Edición Especial",
    "author": "Antoine de Saint-Exupéry",
    "year": 1943,
    "pages": 100
}'

### Eliminar un libro

curl -X DELETE http://localhost:8000/books/{id}

## 📚 Endpoints

Aquí están los principales endpoints de la API:

| Método | Endpoint           | Descripción                     |
|--------|--------------------|---------------------------------|
| GET    | `/books`            | Obtiene todos los libros        |
| GET    | `/books/{id}`       | Obtiene un libro por su ID      |
| POST   | `/books`            | Crea un nuevo libro             |
| PUT    | `/books/{id}`       | Actualiza un libro por su ID    |
| DELETE | `/books/{id}`       | Elimina un libro por su ID      |

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Si deseas contribuir al proyecto, sigue estos pasos:

1. Realiza un **fork** del repositorio.
2. Crea una nueva **rama** para tus cambios:
   
   ```bash
   git checkout -b nombre-de-la-rama

   
