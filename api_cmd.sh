Api platform
  Tests
  auth
	


curl localhost:8000/api/users \
--request GET

curl localhost:8000/api/users \
--request POST \
--data-raw '{
	"Username": "Vishal21",
	"Password": "Vishal21",
	"Email": "Vishal2xx"
}'

curl -vv localhost:8000/api/users/2 \
--request DELETE 

