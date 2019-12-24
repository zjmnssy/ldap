package ldap

import (
	"fmt"
	"time"

	"github.com/zjmnssy/zlog"

	"github.com/go-ldap/ldap"
)

// Connection 获取AD服务器的连接.
func Connection(domain string, timeout time.Duration) (*ldap.Conn, error) {
	c, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", domain, 389))
	if err != nil {
		return nil, err
	}

	c.SetTimeout(timeout * time.Millisecond)

	return c, nil
}

// Login 登录AD域.
func Login(c *ldap.Conn, account string, password string) (*ldap.Conn, error) {
	if c == nil {
		return nil, fmt.Errorf("connection is nil")
	}

	err := c.Bind(account, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Close 关闭连接.
func Close(c *ldap.Conn) {
	if c == nil {
		zlog.Prints(zlog.Notice, "Ldap", "connection is nil")
		return
	}

	c.Close()
}

func init() {
	ldap.DefaultTimeout = time.Duration(1000) * time.Millisecond
}
