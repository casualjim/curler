# Curler

A golang HTTP middleware that dumps request as curl commands.
When the request has a json body then it will also dump the json.

```go
curler.New(handler)
```