##### Swagger :
#####
https://github.com/vv198x/userSL/blob/main/docs/swagger.yaml

##### Docker :
#####
https://github.com/vv198x/userSL/blob/main/Docker/docker-compose.yml

##### Endpoints :
#####
| Метод      | Параметры               | Описание                    |
|------------|-------------------------|-----------------------------| 
| POST (C)   | /api/users/v1/          | Создать пользователя        | 
| GET (R)    | /api/users/v1/          | Показать всех пользователей | 
| GET (R)    | /api/users/v1/user_name | Показать одного             | 
| PUT (U)    | /api/users/v1/user_name | Обновить                    | 
| DELETE (D) | /api/users/v1/user_name | Удалить                     | 



| Пакеты использовал                  |
|-------------------------------------|
| github.com/go-pg/pg                 |
| github.com/go-playground/validator  |
| github.com/labstack/echo  |
| github.com/swaggo/  |

