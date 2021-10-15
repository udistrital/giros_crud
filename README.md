# giros_crud

El API giros_crud, proporciona interfaces para la manipulación(CRUD) de los datos almacenados en una base de datos relacional PostgreSQl (registra giros, contabilización giros, ordenes de pago relaciona al giro, etc). Esta API representa la capa de datos del subsistema de giros, el cual, a su vez, hace parte de el sistema de gestión financiero KRONOS.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Postgres](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

```shell
# Ejemplo que se debe actualizar acorde al proyecto
      - GIROS_CRUD_HTTP_PORT=[Puerto asignado para la ejecución del API]
      - GIROS_CRUD_PGUSER=[Usuario de la base de datos]
      - GIROS_CRUD_PGPASS=[Clave del usuario para la conexión a la base de datos]
      - GIROS_CRUD_PGHOST=[Host de conexión a la base de datos]
      - GIROS_CRUD_PGDB=[Nombre de la base de datos]
      - GIROS_CRUD_PGSCHEMA=[Esquema de la base de datos]
      - GIROS_CRUD_RUN_MODE=dev

      # Ejemplo
      - GIROS_CRUD_HTTP_PORT=8080
      - GIROS_CRUD_PGUSER=postgres
      - GIROS_CRUD_PGPASS=***
      - GIROS_CRUD_PGHOST=127.0.0.1
      - GIROS_CRUD_PGPORT=5432
      - GIROS_CRUD_PGDB=giros_db
      - GIROS_CRUD_PGSCHEMA=giros
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con GIROS_CRUD_...

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/giros_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/giros_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
GIROS_CRUD_HTTP_PORT=8080 GIROS_CRUD_PGURL=127.0.0.1 GIROS_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile

```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/giros_crud

#2. Moverse a la carpeta del repositorio
cd giros_crud

#3. Crear un fichero con el nombre **custom.env**
touch .env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias

```shell
# Not Data
```

## Modelo de datos

[Mode Relacional giros_crud](database/Modelo_giros.svg)

## Estado CI

| Develop | Release | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/giros_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/giros_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/giros_crud/status.svg?ref=refs/heads/release)](https://hubci.portaloas.udistrital.edu.co/udistrital/giros_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/giros_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/giros_crud) |

## Licencia

This file is part of giros_crud

giros_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

giros_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with giros_crud. If not, see https://www.gnu.org/licenses/.
