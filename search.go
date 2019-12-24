package ldap

import (
	"fmt"

	"github.com/go-ldap/ldap"
)

const (
	defaultSearchPageSize = 256
)

// Query 分页查询AD用户列表信息.
func Query(c *ldap.Conn, domain string, filter string, attributeList []string) (*ldap.SearchResult, error) {
	if c == nil {
		return nil, fmt.Errorf("connection is nil")
	}

	searchRequest := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributeList,
		nil,
	)

	sr, err := c.SearchWithPaging(searchRequest, defaultSearchPageSize)
	if err != nil {
		return nil, err
	}

	return sr, nil
}
