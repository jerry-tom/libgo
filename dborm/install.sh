#!/bin/bash

set -euo pipefail

# install gorm
echo "install gorm"
go get -u gorm.io/gorm

# install mysql
echo "install mysql pkg"
go get -u gorm.io/driver/mysql

# install sqlite
echo "install sqlite pkg"
go get -u gorm.io/driver/sqlite