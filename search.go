package ldap

import (
	"fmt"

	"github.com/go-ldap/ldap"
)

// DefaultSearchPageSize 查询数据默认页大小.
var DefaultSearchPageSize uint32

// Query 分页查询AD用户列表信息.
func Query(c *ldap.Conn, searchbase string, filter string, attributeList []string) (*ldap.SearchResult, error) {
	if c == nil {
		return nil, fmt.Errorf("connection is nil")
	}

	searchRequest := ldap.NewSearchRequest(
		searchbase,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributeList,
		nil,
	)

	sr, err := c.SearchWithPaging(searchRequest, DefaultSearchPageSize)
	if err != nil {
		return nil, err
	}

	return sr, nil
}
