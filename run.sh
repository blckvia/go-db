#!/bin/sh -e

echo "Запущен скрипт с параметрами: $@"

REPO_NAME="go-queue"

CGO_ENABLED=0

PPROF_UI_PORT=77345
PPROF_DEFAULT_PORT=8080
PPROF_DEFAULT_CPU_DURATION=5
UNIT_COVERAGE_MIN=80
# Подтянуть зависимости
deps(){
  go get ./...
}

# Собрать исполняемый файл
build(){
  deps
  go build ./cmd/order
}

# Собрать docker образ
build_docker() {
  build
  docker build -t "$REPO_NAME:local" .
  rm ./"$REPO_NAME"
}

# Cписок команд
using(){
  echo "  Укажите команду при запуске: ./run.sh [command]"
  echo "  Список команд:"
  echo "  deps - подтянуть зависимости"
  echo "  unit_coverage - запуск unit тестов и проверка покрытия кода тестами"
  echo "  build - собрать приложение"
  echo "  build_docker - собрать локальный docker образ"
  echo "  fmt - форматирование кода при помощи 'go fmt'"
  echo "  vet - проверка правильности форматирования кода"
  echo "  run - запуск приложения локально через go run"
  echo "  pprof_cpu HOST [SECONDS] - сбор метрик нагрузки на cpu из pprof (НЕ НАСТРОЕНО)"
  echo "  pprof_heap HOST - запустить сбор метрик памяти из pprof (НЕ НАСТРОЕНО)"
}

# Запустить контейнер с приложением
compose(){
  docker build -t go-queue:develop .
}

fmt() {
  echo "run go fmt"
  go fmt ./...
}

vet() {
  echo "run go vet"
  go vet ./...
}

# Запустить сбор метрик нагрузки на cpu из pprof
pprof_cpu(){
  local SECS=${3:-$PPROF_DEFAULT_CPU_DURATION}
  local HOST=$2

  go tool pprof -http :$PPROF_UI_PORT $HOST/debug/pprof/profile?seconds=$SECS
}

# Запустить сбор метрик памяти из pprof
pprof_heap(){
  local HOST=$2

  go tool pprof -http :$PPROF_UI_PORT $HOST/debug/pprof/heap
}

run(){
  go run ./cmd/app/main.go
}

unit_coverage() {
  echo "run test coverage"
  go test -coverpkg=./... -coverprofile=cover_profile.out.tmp $(go list ./...)
  # remove generated code and mocks from coverage
  < cover_profile.out.tmp grep -v -e "mock"  > cover_profile.out
  rm cover_profile.out.tmp
  CUR_COVERAGE=$( go tool cover -func=cover_profile.out | tail -n 1 | awk '{ print $3 }' | sed -e 's/^\([0-9]*\).*$/\1/g' )
  rm cover_profile.out
  if [ "$CUR_COVERAGE" -lt $UNIT_COVERAGE_MIN ]
  then
    echo "coverage is not enough $CUR_COVERAGE < $UNIT_COVERAGE_MIN"
    return 1
  else
    echo "coverage is enough $CUR_COVERAGE >= $UNIT_COVERAGE_MIN"
  fi
}

############### НЕ МЕНЯЙТЕ КОД НИЖЕ ЭТОЙ СТРОКИ #################
command="$1"
if [ -z "$command" ]
then
 using
 exit 0;
else
 $command $@
fi

