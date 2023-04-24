# Vault-service
Сервис представляет собой хранилище JSON-объектов с HTTP-интерфейсом. Сохраненные объекты размещаются в оперативной памяти, имеется возможность задать время жизни объекта.

команды
make build - сборка сервиса

make build-img - сборка докер образа

make run-img - запуск на портах 8040 и 3110 (дебаг)

curl --location --request GET 'http://localhost:3110/live' 

curl --location --request GET 'http://localhost:3110/ready'

curl --location --request GET 'http://localhost:3110/metrics'


Добавление ключа в хранилище:
curl --location --request PUT 'http://localhost:8040/objects/firstkey' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "object1": "object1",
    "object2": "object2"
}'

Получение ключа из хранилища

curl --location --request GET 'http://localhost:8040/objects/firstkey' \
--header 'Content-Type: text/plain'
