FROM postgres:17.1

ENV POSTGRES_USER="paul"
ENV POSTGRES_DB="gofinance"

RUN echo $POSTGRES_PASSWORD

COPY schema.sql  /docker-entrypoint-initdb.d/
