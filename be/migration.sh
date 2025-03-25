#!/bin/bash

if [ $# -ne 2 ]; then
  echo "引数がおかしい"
  exit 9
fi

case $1 in
  1 ) echo "migrationファイルを作成"
      goose -dir ./infra/database/migrations create $2 sql
      ;;
  2 ) echo "migrationファイルを適応"
      goose -dir ./infra/database/migrations up
      ;;
  3 ) echo "migrationの最新をロールバック"
      goose -dir ./infra/database/migrations down
      ;;
  * ) echo "実行引数がおかしい"
      exit 9
      ;;
esac