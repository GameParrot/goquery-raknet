package goqueryraknet

import (
	"context"
	"log/slog"
	"net"

	"github.com/gameparrot/goquery"
	"github.com/sandertv/go-raknet"
	"github.com/sandertv/gophertunnel/minecraft"
)

func CreateGophertunnelNetwork(name string, queryServer *goquery.QueryServer) {
	minecraft.RegisterNetwork(name, func(l *slog.Logger) minecraft.Network { return RakNetQuery{l: l, q: queryServer} })
}

// RakNet is an implementation of a RakNet v10 Network with query.
type RakNetQuery struct {
	l *slog.Logger
	q *goquery.QueryServer
}

// DialContext ...
func (r RakNetQuery) DialContext(ctx context.Context, address string) (net.Conn, error) {
	return raknet.Dialer{ErrorLog: r.l.With("net origin", "raknet")}.DialContext(ctx, address)
}

// PingContext ...
func (r RakNetQuery) PingContext(ctx context.Context, address string) (response []byte, err error) {
	return raknet.Dialer{ErrorLog: r.l.With("net origin", "raknet")}.PingContext(ctx, address)
}

// Listen ...
func (r RakNetQuery) Listen(address string) (minecraft.NetworkListener, error) {
	return raknet.ListenConfig{ErrorLog: r.l.With("net origin", "raknet"), UpstreamPacketListener: NewQueryUpstreamPacketListener(r.q)}.Listen(address)
}
