
sudo docker run -it -e POSTGRES_PASSWORD=password -p 5432:5432 postgres:14-alpine

sudo docker run -d -p 5432:5432 postgres:14-alpine

sudo docker volume create pg_vol

sudo docker run --name pg_kafka -d -e POSTGRES_PASSWORD=password -v pg_vol:/var/lib/postgresql/data -p 5432:5432 postgres:14-alpine

sudo docker run -d --name broker apache/kafka:latest

sudo docker run -d -p 9092:9092 --name broker apache/kafka:latest

