#!/usr/bin/bash

curl -H 'Content-Type: application/json' \
     -d '{ "hostname": "localhost:8090", "port": 7000, "type": 0 }' \
     -X POST localhost:8080/services/register
