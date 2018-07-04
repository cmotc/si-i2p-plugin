package jumpresolver

import (
	"golang.org/x/net/context"
	"net"
	"net/http"
)

import (
	"github.com/eyedeekay/si-i2p-plugin/src/addresshelper"
)

type JumpResolver struct {
	addressbook    *dii2pah.AddressHelper
	jumpHostString string
	jumpPortString string
}

func (j *JumpResolver) CheckAddressHelper(url *http.Request) (*http.Request, bool) {
	return j.CheckAddressHelper(url)
}

func (j JumpResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	addr, err := net.ResolveIPAddr("ip", name)
	//addr :=
	if err != nil {
		return ctx, nil, err
	}
	return ctx, addr.IP, err
}

func NewJumpResolver(host, port string) (*JumpResolver, error) {
	return NewJumpResolverFromOptions(
		SetJumpResolverHost(host),
		SetJumpResolverPort(port),
	)
}

func NewJumpResolverFromOptions(opts ...func(*JumpResolver) error) (*JumpResolver, error) {
	var j JumpResolver
	//j.addressHelperURL = "inr.i2p"
	j.jumpHostString = "127.0.0.1"
	j.jumpPortString = "7054"
	//j.bookPath = "addressbook.txt"
	for _, o := range opts {
		if err := o(&j); err != nil {
			return nil, err
		}
	}
	return &j, nil
}