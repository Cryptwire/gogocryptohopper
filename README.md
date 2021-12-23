# GoGoCryptohopper

A [Go](https://golang.org/) client for working with the [Cryptohopper API](https://www.cryptohopper.com/api-documentation)

## Structure of the client

### Request/Response

The definitions for request and response formats are located in `/pkg` and
can be used independently of the client itself.

## Examples

See the examples folder for code samples.

Example: if you'd like to run the fetch Hoppers sample. You will need to define
the access token variable which needs to contain a valid Cryptohopper access
token

```go
go run examples/hopper.go
```

## Licence

MIT

## Contact

Twitter: [@cryptwireapp](https://twitter.com/cryptwireapp)
