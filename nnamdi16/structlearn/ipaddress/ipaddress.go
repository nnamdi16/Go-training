package ipaddress

import "fmt"

//IPAddr with 4 byte array
type IPAddr [4]byte

func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",a[0],a[1],a[2],a[3])
}

