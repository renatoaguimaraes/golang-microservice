# golang-microservice

### Project setup
```console
$ mkdir -p $GOTPATH/src/github.com/renatoaguimaraes/
$ cd $GOTPATH/src/github.com/renatoaguimaraes/
$ git clone https://github.com/renatoaguimaraes/golang-microservice.git
$ cd golang-microservice
$ go get github.com/gocql/gocql
$ go get github.com/gorilla/mux
```

### Cassandra setup
```console
$ docker run --name some-cassandra -p 9042:9042 -d cassandra
$ docker exec -it some-cassandra bash
root@4d3c1fdd8f4c:/# cqlsh
Connected to Test Cluster at 127.0.0.1:9042.
[cqlsh 5.0.1 | Cassandra 3.11.4 | CQL spec 3.4.4 | Native protocol v4]
Use HELP for help.
cqlsh> CREATE KEYSPACE users WITH replication = {'class':'SimpleStrategy', 'replication_factor':1};
cqlsh> CREATE TABLE users.users (id text, firstname text, lastname text, PRIMARY KEY(id));
```
### Project build and run
```console
$ go build main.go
$ ./main
```
### Resources
```
List users  - GET    - http://localhost:8000/users
Get user    - GET    - http://localhost:8000/users/{id}
Create user - POST   - http://localhost:8000/users/{id}
Remove user - DELETE - http://localhost:8000/users/{id}
```

### Cassandra driver

[4 simple rules when using the DataStax drivers for Cassandra](https://www.datastax.com/dev/blog/4-simple-rules-when-using-the-datastax-drivers-for-cassandra)
