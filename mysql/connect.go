package mysql

import (
	"github.com/lvchengchang/canalysis/log"
	"io"
	"net"
)

var connlog = log.Logger.With()

type Connection struct {
	conn *net.Conn

	Read  *io.Reader
	Write *io.Writer
}

func NewConnection() *Connection {
	return &Connection{}
}

func (conn *Connection) Connect() {

}
