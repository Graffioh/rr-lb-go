# toy load balancer in go
simple round robin load balancer implementation from scratch

## how it works?

simply run
- the backend servers (inside be folder):

~~~
go run be.go
~~~

- the load balancer (inside lb folder):

~~~
go run lb.go
~~~

then to test it just curl localhost:6969 and see the load balancer redirect your request to different backend servers using round robin algorithm
