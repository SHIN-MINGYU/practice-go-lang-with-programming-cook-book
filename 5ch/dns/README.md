# check domain name

- net package provide various useful feature for DNS lookup

- this information is simliar with result by unix's 'dig' command

- this example show how to collect this data

- Search how to mapping URL at all of IPv4 address and IPv6 address

- when GODEBUG=netdns= to 'go' => this is use pure go dns resovler,
- when GODEBUG=netdns= to 'cgo' => this is use cgo resolver.
- default is first one
