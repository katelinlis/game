#!/bin/sh

# Запуск бэкенда
/main &

# Запуск nginx
nginx -g 'daemon off;'