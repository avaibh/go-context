# Learn how context work in Go

This demo show what happens during context timeout. This demo is only scoped to TCP/IP model.

This demo is aimed at testing the following hypothesis:
> A context timeout only kills an HTTP request and not the TCP connection.

In this demo we spin up three web servers:
1. Client (`localhost:9000`)
2. Middleware (`localhost:9001`)
3. Server (`localhost:9002`)

`Client` calls the `Server` through a `Middleware`. 

Preface: In the client I have setup a context timeout of 80 msec. And the middleware sleeps for 2 secs before calling the server. 

I have created a custom http client to connect with the middleware. This gives a reusable TCP connection (with lifetime of 10 secs), i.e. all http requests to the middleware from the client within 10 secs will go through the same TCP connection.
```go
client := &http.Client{
    Transport: &http.Transport{
        MaxIdleConnsPerHost: 20,
    },
    Timeout: 10 * time.Second,
}
```

We will prove that a context timeout does not kill/end a TCP connection through following two observations:

1. If context timeout kills a TCP connection, we would expect to see a failure for the 2nd http request from the client to the middleware. But from this demo it was evident that we succesfully sent an HTTP request to the middleware.
2. If context timeout kills a TCP connection, we would expect that once the 80 msec context times out, the middleware process ends and it drops the request in the middle of its sleep. In reality though, the connection between client and middleware is not affected by the request's context timeout. And that is testified when middleware succesfully wakes up from its short 2 sec nap.

So what good is the context timeout then. `If context is passed correctly, the context timeout make sure all blocking calls are cancelled`. And so in the demo you will find that, by the time middleware wakes up from the sleep the context of http request sent from client has already timed out. So the http blocking call in client is cancelled and the client process ends before the middleware process.


Key learnings:
1. Context timeout does not end a downstream process. It only frees the client from the blocking call so that the client is not kept waiting for the it to end. Note the downstream process will still run after the context timeout. And so we need to make sure context flows naturally through our program.
2. Context timeout does not end a TCP connection. It's only the HTTP blocking call that gets cancelled.
3. Using this we can say context in go operates at application layer in OSI model. [disclaimer: this is my observation]

Above learnings can be used to understand how context timeout will work in any TCP network communication. For example a communication between Client, PGBouncer and Postgres happens over TCP. 
