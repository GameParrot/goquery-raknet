package goqueryraknet

import (
	"net"

	"github.com/gameparrot/goquery"
)

type QueryUpstreamPacketListener struct {
	q *goquery.QueryServer
}

func NewQueryUpstreamPacketListener(q *goquery.QueryServer) *QueryUpstreamPacketListener {
	return &QueryUpstreamPacketListener{q: q}
}

func (q *QueryUpstreamPacketListener) QueryHandler() *goquery.QueryServer {
	return q.q
}

func (q *QueryUpstreamPacketListener) ListenPacket(network, address string) (net.PacketConn, error) {
	conn, err := net.ListenPacket(network, address)
	if err != nil {
		return conn, err
	}
	return newQueryHandlerConn(conn, q.q), nil
}
