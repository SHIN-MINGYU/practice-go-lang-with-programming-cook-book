package dns

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// Lookup struct save DNS information what we have interested in
type Lookup struct {
	cname string
	hosts []string
}

// this function can use to Lookup's object
func (l *Lookup)String() string{
	result := ""
	for _,host := range l.hosts{
		result += fmt.Sprintf("%s IN A %s\n", l.cname, host)
	}
	return result
}

// LookupAddress return DNSLookup structure about input address
func LookupAddress(address string)(*Lookup, error){
	cname, err := net.LookupCNAME(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up CNAME")
	}
	hosts, err := net.LookupHost(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up HOST")
	}

	return &Lookup{cname:cname, hosts:hosts},nil
}