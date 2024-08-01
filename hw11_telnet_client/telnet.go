package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type client struct {
	address string
	timeout time.Duration
	conn    net.Conn
	in      io.ReadCloser
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) Close() error {
	errMsgs := []string{}
	if err := c.conn.Close(); err != nil {
		errMsgs = append(errMsgs, err.Error())
	}

	if err := c.in.Close(); err != nil {
		errMsgs = append(errMsgs, err.Error())
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf("failed to close client with errors: %s", strings.Join(errMsgs, ", "))
	}

	return nil
}

func (c *client) Send() error {
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *client) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	return err
}
