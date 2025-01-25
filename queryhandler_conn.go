package goqueryraknet

import (
	"bytes"
	"net"
	"time"

	"github.com/gameparrot/goquery"
)

type queryHandlerConn struct {
	q      *goquery.QueryServer
	parent net.PacketConn
}

func newQueryHandlerConn(parent net.PacketConn, q *goquery.QueryServer) *queryHandlerConn {
	return &queryHandlerConn{parent: parent, q: q}
}

func (q *queryHandlerConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	n, addr, err = q.parent.ReadFrom(p)
	if err != nil {
		return
	}
	if len(p) >= 2 && p[0] == 0xfe && p[1] == 0xfd {
		buf := bytes.NewBuffer(p[:n])
		if q.q.Handle(buf, addr) == nil {
			q.WriteTo(buf.Bytes(), addr)
		}
		n = 0
	}
	return
}

func (q *queryHandlerConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return q.parent.WriteTo(p, addr)
}

func (q *queryHandlerConn) Close() error {
	return q.parent.Close()
}

func (q *queryHandlerConn) LocalAddr() net.Addr {
	return q.parent.LocalAddr()
}

func (q *queryHandlerConn) SetDeadline(t time.Time) error {
	return q.parent.SetDeadline(t)
}

func (q *queryHandlerConn) SetReadDeadline(t time.Time) error {
	return q.parent.SetReadDeadline(t)
}

func (q *queryHandlerConn) SetWriteDeadline(t time.Time) error {
	return q.parent.SetWriteDeadline(t)
}
