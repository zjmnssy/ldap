package main

import (
	"github.com/zjmnssy/ldap"
	"github.com/zjmnssy/zlog"
)

func ad() {
	searchBase := "ou=数篷科技,dc=datacloak3,dc=com"
	userName := "CN=test01,OU=数篷科技,DC=datacloak3,DC=com"
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
		zlog.Prints(zlog.Warn, "example", "connect ad error = %s", err)
		return
	}

	_, err = ldap.Login(c, userName, password)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "login ad error = %s", err)
		return
	}

	result, err := ldap.Query(c, searchBase, filter, attributes)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "query error = %s", err)
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

func openldap() {
	searchBase := "dc=zjmdomain,dc=com"
	userName := "uid=ztest1,dc=zjmdomain,dc=com"
	password := "qqq111"
	//filter := "(&(objectClass=inetOrgPerson))"
	filter := "(&(objectClass=*))"
	attributes := make([]string, 0)
	attributes = append(attributes, "cn")
	attributes = append(attributes, "sn")
	attributes = append(attributes, "givenName")
	attributes = append(attributes, "mail")
	attributes = append(attributes, "entryDN")
	//attributes = append(attributes, "entryUUID")
	attributes = append(attributes, "uid")
	attributes = append(attributes, "mobile")

	c, err := ldap.Connect("10.10.17.166:389", 2000)
	defer ldap.Close(c)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "connect ad error = %s", err)
		return
	}

	_, err = ldap.Login(c, userName, password)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "login ad error = %s", err)
		return
	}

	result, err := ldap.Query(c, searchBase, filter, attributes)
	if err != nil {
		zlog.Prints(zlog.Warn, "example", "query error = %s", err)
		return
	}

	for _, entry := range result.Entries {
		if entry.GetAttributeValue("cn") == "" {
			continue
		}

		zlog.Prints(zlog.Info, "example", "cn = %s", entry.GetAttributeValue("cn"))
		zlog.Prints(zlog.Info, "example", "sn = %s", entry.GetAttributeValue("sn"))
		zlog.Prints(zlog.Info, "example", "givenName = %s", entry.GetAttributeValue("givenName"))
		zlog.Prints(zlog.Info, "example", "mail = %s", entry.GetAttributeValue("mail"))
		zlog.Prints(zlog.Info, "example", "entryDN = %s", entry.GetAttributeValue("entryDN"))
		//zlog.Prints(zlog.Info, "example", "entryUUID = %s", entry.GetAttributeValue("entryUUID"))
		zlog.Prints(zlog.Info, "example", "uid = %s", entry.GetAttributeValue("uid"))
		zlog.Prints(zlog.Info, "example", "mobile = %s\n", entry.GetAttributeValue("mobile"))
	}
}

func main() {
	openldap()
}
