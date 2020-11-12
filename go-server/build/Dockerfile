FROM postgres:alpine

WORKDIR /docker-entrypoint-initdb.d
COPY initdb-scripts/* .

# could use _FILE
ENV POSTGRES_USER=username
ENV POSTGRES_PASSWORD=password