#!/usr/bin/zsh

curl -H "Content-Type: application/json" -X POST \
-d '{"login":"user","password":"user"}' \
http://localhost:8000/api/users/v1/auth/

curl -H "Content-Type: application/json" -X POST \
-d '{"login":"new","password":"new1","rule":2,"name":"first name new","last_name":"last name new","dob":200000}' \
"http://localhost:8000/api/users/v1/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFkbWluIiwicnVsZSI6MH0.lhnrxwZZtcg1ZH47P_v4hEZEuUGLX8uB44OlYwvZ8ms"

curl -H "Content-Type: application/json" -X PUT \
-d '{"login":"new1","password":"new","rule":1,"name":"first name new","last_name":"last name new","dob":1}' \
"http://localhost:8000/api/users/v1/new?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFkbWluIiwicnVsZSI6MH0.lhnrxwZZtcg1ZH47P_v4hEZEuUGLX8uB44OlYwvZ8ms"


curl "http://localhost:8000/api/users/v1/new1?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFkbWluIiwicnVsZSI6MH0.lhnrxwZZtcg1ZH47P_v4hEZEuUGLX8uB44OlYwvZ8ms"

curl -H "Content-Type: application/json" -X DELETE \
"http://localhost:8000/api/users/v1/new1?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFkbWluIiwicnVsZSI6MH0.lhnrxwZZtcg1ZH47P_v4hEZEuUGLX8uB44OlYwvZ8ms"


curl "http://localhost:8000/api/users/v1/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFkbWluIiwicnVsZSI6MH0.lhnrxwZZtcg1ZH47P_v4hEZEuUGLX8uB44OlYwvZ8ms"
