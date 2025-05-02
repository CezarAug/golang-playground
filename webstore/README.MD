[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=CezarAug_golang-playground&metric=alert_status&token=d7bebff9b85a057a22169ec63457ef63b29d7555)](https://sonarcloud.io/summary/new_code?id=CezarAug_golang-playground)

### How to run

#### Preparing the DB

Suggestion, create a docker image with postgres:

-- DB
docker run -it --name postgres-db -e POSTGRES_PASSWORD=qwerty -d -p 5051:5432 postgres
-- Admin
docker run -it --name pgadmin-web -e PGADMIN_DEFAULT_EMAIL=admin@admin.com -e PGADMIN_DEFAULT_PASSWORD=qwerty -d -p 5050:80  dpage/pgadmin4


-- Check the IPAddress for the DB image
docker inspect postgres-db

"IPAddress": "172.17.0.2",

-- Admin URL:
localhost:5050/admin

User: admin@admin.com
Password: querty

Once accessing, the connection can be created with the used credentials, IPAddress and database name postgres

<!-- TODO: Prepare a docker compose that would run the whole environment -->


### Used Libraries

pq: For connecting with postgresql. https://pkg.go.dev/github.com/lib/pq