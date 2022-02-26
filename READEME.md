# Learn how context work in Golang

This demo show what happens when context gets timed out. This demo is only constrained to TCP/IP model.

Key learnings:
1. Context timeout do not kill the TCP connection. As you will see it's only the HTTP request that get's cancelled.
2. Using this we can say context in go operates at application layer in OSI model. 
3. Above learning can be used to understand how context timeout will work in any TCP network communication.