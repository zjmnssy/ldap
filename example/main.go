package main

import (
	"github.com/zjmnssy/ldap"
	"github.com/zjmnssy/zlog"
)

func main() {
	searchBase := "ou=数篷科技,dc=datacloak3,dc=com"
	userName := "test01@datacloak3.com"
	password := "start@2018"
	filter := "(&(SamAccountName=*))"
	attributes := make([]string, 0)
	attributes = append(attributes, "SamAccountName")
	attributes = append(attributes, "Mail")
	attributes = append(attributes, "Mobile")
	attributes = append(attributes, "Department")
	attributes = append(attributes, "Sn")
	attributes = append(attributes, "Givename")

	c, err := ldap.Connect("10.10.3.65:389", 2000)
	defer ldap.Close(c)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "connect ad error : %s", err)
		return
	}

	_, err = ldap.Login(c, userName, password)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "login ad error : %s", err)
		return
	}

	result, err := ldap.Query(c, searchBase, filter, attributes)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "query error : %s", err)
		return
	}

	for _, entry := range result.Entries {
		zlog.Prints(zlog.Info, "example", "SamAccountName = %s", entry.GetAttributeValue("sAMAccountName"))
		zlog.Prints(zlog.Info, "example", "Mail = %s", entry.GetAttributeValue("mail"))
		zlog.Prints(zlog.Info, "example", "Mobile = %s", entry.GetAttributeValue("mobile"))
		zlog.Prints(zlog.Info, "example", "Department = %s", entry.GetAttributeValue("department"))
		zlog.Prints(zlog.Info, "example", "Sn = %s", entry.GetAttributeValue("sn"))
		zlog.Prints(zlog.Info, "example", "Givename = %s", entry.GetAttributeValue("givename"))
	}

}
