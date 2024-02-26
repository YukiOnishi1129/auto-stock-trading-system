#!/usr/bin/env bash

# careful
# don't act if you don't install sqlboiler and sqlboiler-mysql in mac
#  go install github.com/volatiletech/sqlboiler/v4@latest
#  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

# create entity at batch-service
sqlboiler mysql \
    -c db/sqlboiler.toml \
    -o micro-service/batch-service/database/entity \
    -p entity \
    --no-tests --wipe

# create entity at user-service
sqlboiler mysql \
    -c db/sqlboiler.toml \
    -o micro-service/user-service/database/entity \
    -p entity \
    --no-tests --wipe