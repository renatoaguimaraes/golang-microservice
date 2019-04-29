# golang-microservice

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
