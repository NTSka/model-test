#!/bin/bash

PATH_ROOT=$(dirname "$0")
source "$PATH_ROOT/tools/lib.sh" || exit 1

MIGRATION_PATH=$1
DATABASE=$2
COMMAND=$3
PARAM=$4

if [ "$3" = "" ]; then
  migrate_help
  error "Ошибка параметров"
fi

if [[ "$COMMAND" != "down" && "$COMMAND" != "up" && "$COMMAND" != "force" ]]; then
  error "Варианты команд down|up|force: ${COMMAND}"
fi

if ! [ -d "$MIGRATION_PATH" ]; then
  error "Путь к миграции не существует: ${MIGRATION_PATH}"
fi

migrate "$MIGRATION_PATH" "$DATABASE" "$COMMAND" "$PARAM"