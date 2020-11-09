FROM postgres:alpine

# could use _FILE
ENV POSTGRES_USER=
ENV POSTGRES_PASSWORD=

WORKDIR /docker-entrypoint-initdb.d
COPY initdb-scripts/* .
