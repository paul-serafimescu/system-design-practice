#!/usr/bin/bash

curl -H 'Content-Type: application/json' \
     -d '{ "email":"noreply@gmail.com","password":"8dccd2fa9e2340d455740d1b4ae2dca6774c3b7d7dbdef9f362c1c59ecec4016"}' \
     -X POST localhost:8080/auth/login
