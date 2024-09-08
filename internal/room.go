package internal

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) boradcast(sender *client, msg string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
