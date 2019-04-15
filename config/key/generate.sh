#!/bin/bash

KEYNAME=${1:-jwtrsa}
openssl genrsa -out private.key 1024 && openssl rsa -in private.key -outform PEM -pubout -out public.pem