# goquery-raknet
Support for UT3 query in go-raknet and gophertunnel

# Usage
Set UpstreamPacketListener to QueryUpstreamPacketListener to use in go-raknet. In Gophertunnel, you can use CreateGophertunnelNetwork to create a network with query support. Gophertunnel example:  
```go
q := goquery.New(map[string]string{}, []string{})
goqueryraknet.CreateGophertunnelNetwork("raknetquery", q)
list, err := minecraft.ListenConfig{}.Listen("raknetquery", "0.0.0.0:19132")
if err != nil {
  panic(err)
}
```
