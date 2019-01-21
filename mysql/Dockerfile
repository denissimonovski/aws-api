FROM mysql
ENV MYSQL_ROOT_PASSWORD supersecret
ENV MYSQL_DATABASE banka
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/
