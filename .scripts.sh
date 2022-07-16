#!/usr/bin/env bash
# Nota 1: Para más información:
# https://stackoverflow.com/questions/10376206/what-is-the-preferred-bash-shebang
# Nota 2: El anterior shebang no es suficiente, de momento se debe usar con source

# La idea de este script es dejar funciones que se puedan usar principalmente en
# CI, pero también podrían usarse localmente. Modificar según sea necesario

# Para chequear que una variable de entorno exista
function checkEnvVar() {
  if [ -z "\$${1}" ]; then
    echo "${1} unset"
  else
    echo "${1} set"
  fi
}

function checkEnvMode() {
  checkEnvVar GIROS_CRUD_RUN_MODE
}

# checkEnvMode
