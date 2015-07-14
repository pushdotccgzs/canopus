package canopus

import (
	. "github.com/zubairhamed/go-commons/network"
	"net"
	"strconv"
	"strings"
)

func NewRequest(messageType uint8, messageMethod CoapCode, messageId uint16) *CoapRequest {
	msg := NewMessage(messageType, messageMethod, messageId)
	msg.Token = []byte(GenerateToken(8))

	return &CoapRequest{
		msg: msg,
	}
}

func NewRequestFromMessage(msg *Message) *CoapRequest {
	return &CoapRequest{
		msg:   msg,
	}
}

func NewClientRequestFromMessage(msg *Message, attrs map[string]string, conn *net.UDPConn, addr *net.UDPAddr) *CoapRequest {
	return &CoapRequest{
		msg:   msg,
		attrs: attrs,
		conn:  conn,
		addr:  addr,
	}
}

type CoapRequest struct {
	msg   	*Message
	attrs 	map[string]string
	conn  	*net.UDPConn
	addr  	*net.UDPAddr
	server 	*CoapServer
}

func (c *CoapRequest) SetMediaType(mt MediaType) {
	c.msg.AddOption(OPTION_CONTENT_FORMAT, mt)
}

func (c *CoapRequest) GetConnection() *net.UDPConn {
	return c.conn
}

func (c *CoapRequest) GetAddress() *net.UDPAddr {
	return c.addr
}

func (c *CoapRequest) GetAttributes() map[string]string {
	return c.attrs
}

func (c *CoapRequest) GetAttribute(o string) string {
	return c.attrs[o]
}

func (c *CoapRequest) GetAttributeAsInt(o string) int {
	attr := c.GetAttribute(o)
	i, _ := strconv.Atoi(attr)

	return i
}

func (c *CoapRequest) GetMessage() *Message {
	return c.msg
}

func (c *CoapRequest) SetStringPayload(s string) {
	c.msg.Payload = NewPlainTextPayload(s)
}

func (c *CoapRequest) SetRequestURI(uri string) {
	c.msg.AddOptions(NewPathOptions(uri))
}

func (c *CoapRequest) SetConfirmable(con bool) {
	if con {
		c.msg.MessageType = TYPE_CONFIRMABLE
	} else {
		c.msg.MessageType = TYPE_NONCONFIRMABLE
	}
}

func (c *CoapRequest) SetToken(t string) {
	c.msg.Token = []byte(t)
}

func (c *CoapRequest) IncrementMessageId() {
	c.msg.MessageId = c.msg.MessageId + 1
}

func (c *CoapRequest) GetUriQuery(q string) string {
	qs := c.GetMessage().GetOptionsAsString(OPTION_URI_QUERY)

	for _, o := range qs {
		ps := strings.Split(o, "=")
		if len(ps) == 2 {
			if ps[0] == q {
				return ps[1]
			}
		}
	}
	return ""
}

func (c *CoapRequest) SetUriQuery(k string, v string) {
	c.GetMessage().AddOption(OPTION_URI_QUERY, k+"="+v)
}

//func (c *CoapRequest) Observe(seq int) {
//	if seq != 0 {
//		c.GetMessage().AddOption(OPTION_OBSERVE, strconv.Itoa(seq))
//	} else {
//		c.GetMessage().AddOption(OPTION_OBSERVE, "")
//	}
//}

func (c *CoapRequest) GetServer() *CoapServer {
	return c.server
}
