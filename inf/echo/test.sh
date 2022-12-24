#!/usr/bin/zsh
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFkbWluIiwicnVsZSI6MH0.lhnrxwZZtcg1ZH47P_v4hEZEuUGLX8uB44OlYwvZ8ms"
HOST="localhost:8000"

curl -H "Content-Type: application/json" -X POST \
-d '{"login":"admin","password":"admin"}' \
http://$HOST/api/users/v1/auth/

curl -H "Content-Type: application/json" -X POST \
-d '{"login":"new","password":"new1","rule":2,"name":"first name new","last_name":"last name new","dob":"21-11-2000"}' \
"http://$HOST/api/users/v1/?token=$TOKEN"

curl -H "Content-Type: application/json" -X PUT \
-d '{"login":"new1","password":"new","rule":1,"name":"first name new","last_name":"last name new","dob":"21-11-2001"}' \
"http://$HOST/api/users/v1/new?token=$TOKEN"


curl "http://$HOST/api/users/v1/new1?token=$TOKEN"

curl -H "Content-Type: application/json" -X DELETE \
"http://$HOST/api/users/v1/new1?token=$TOKEN"


curl "http://$HOST/api/users/v1/?token=$TOKEN"
