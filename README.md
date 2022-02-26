# Learn how context work in Golang

This demo show what happens when context gets timed out during http request. The learnings from this demo is only constrained to TCP/IP model.

This demo aimed at testing the following hypothesis:
> A context timeout only kills an HTTP request and not the TCP connection.

In this demo we spin up three web servers:
1. Client
2. Middleware
3. Server

We (the client `localhost:9000`) calls the server through a middleware. In our client we setup a context timeout of 80 msec. 

Our middleware sleeps for 2 secs before calling the server.

Now if context timeouts kills a TCP connection, we would expect that once the 80 msec context times out, the middleware drops the request in middle of its sleep.

In reality though, the connection between client and middleware is not affected by the context timeout (we configured in client) and middleware succesfully wakes up. 

So what good is the context timeout. Context timeout will make sure the middleware does not send the timedout request to server. And so in the demo you will find that, by the time middleware wakes up from the sleep the context of http request sent from client has already timed out. And so middleware does nothing and returns.


Key learnings:
1. Context timeout do not kill the TCP connection. It's only the HTTP request that get's cancelled.
2. Using this we can say context in go operates at application layer in OSI model. 
3. Above learning can be used to understand how context timeout will work in any TCP network communication. For example a communication between client, PGBouncer and Postgres happens over TCP. 