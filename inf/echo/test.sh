#!/usr/bin/zsh

curl -H "Content-Type: application/json" -X POST \
-d '{"login":"new","password":"new4","rule":1,"name":"first name new","last_name":"last name new","dob":200000}' \
-u admin:admin http://localhost:8000/api/users/v1

curl -H "Content-Type: application/json" -X PUT \
-d '{"login":"new1","password":"new","rule":1,"name":"first name new","last_name":"last name new","dob":1}' \
-u admin:admin http://localhost:8000/api/users/v1/new

curl -u admin:admin http://localhost:8000/api/users/v1/new1

curl -H "Content-Type: application/json" -X DELETE \
-u admin:admin http://localhost:8000/api/users/v1/new1

curl -u admin:admin http://localhost:8000/api/users/v1/