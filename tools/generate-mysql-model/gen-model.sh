#!/usr/bin/env bash

# import related config
source .config
convert_tool=db2struct


# check if db2struct exist
if !(hash ${convert_tool} 2>/dev/null); then
  echo "convert tool ${convert_tool} is absent, install..."
  go get github.com/Shelnutt2/db2struct/cmd/db2struct
fi


# generate golang model
for table in ${TABLES[@]};
do
  echo "del old model/${table}.go" && rm ../../model/${table}.go
  struct_name="$(echo ${table:0:1} | tr '[a-z]' '[A-Z]')${table:1}"
  db2struct -H ${HOST} -u ${USER} -p ${PASSWORD}  \
    --package ${GO_PACKAGE_NAME} -d ${DB}  \
    --table ${table} --struct ${struct_name}  \
    --target=../../model/${table}.go \
    --json -v
  echo "finish model/${table}.go\n"
done
