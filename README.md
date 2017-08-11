# go-manage
Go helper library for fident management API

Import this library into your project
```go
import client "github.com/fident/go-manage"
```

Init a new client using your fident instance address and your service key file
```go
testClient, err := client.New("./testkey.json", client.FidentInstanceAddressLocal)
if err != nil {
	panic(err)
}
```

Start querying your fident project
```go
details, err := testClient.GetAccountDetailsForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
if err != nil {
	panic(err)
}
```
