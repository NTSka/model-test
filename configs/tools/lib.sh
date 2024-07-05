#!/bin/bash

error()
{
    local MSG="$1"
    if [ -z "$MSG" ]; then
        echo "ОШИБКА" >&2
    else
        echo "ОШИБКА: $MSG" >&2
    fi
    exit 1
}

get_env_os() {
  uname="$(uname -s)"
  case "${uname}" in
    Darwin*)    os=darwin;;
    *)          os=linux
  esac
  echo "${os}"
}

migrate_help() {
  printf "Usage:\n"
  printf "./migrate.sh [/path/to/migrations] [db_driver://db_dsn] [command] [param]\n"
  printf "CLICKHOUSE\t \"clickhouse://<url>:<port>?password=<password>&database=<db_name>\"\n"
  printf "MYSQL\t\t \"mysql://<user>:<password>@tcp(<url>:<port>)/<db_name>\"\n"
}

migrate() {
  MIGRATION_PATH=$1
  DATABASE=$2
  COMMAND=$3
  PARAM=$4

  CMD="$PATH_ROOT/bin/migrate_$(get_env_os) -path=$MIGRATION_PATH -database \"$DATABASE\" $COMMAND $PARAM"
  echo "$CMD"
  eval "$CMD"
}