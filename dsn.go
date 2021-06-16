package goutils

import (
	"regexp"
)

type DSN struct {
	Protocol string
	User     *string
	Password *string
	Host     string
	Port     *string
	DbName   string
	Opts     *string
}

var (
	dsnPattern = regexp.MustCompile(`^(?P<proto>\w+):\/\/?((?P<username>\w+)?(:(?P<password>[^@]+))*@)*(?P<host>[[a-zA-Z0-9_.-]*[a-zA-Z]+|[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})?(:(?P<port>[0-9]+))*\/?(?P<db_name>([\w\d]+)+)(\?(?P<opts>[\w\d]+=[\w\d]+(\&[\w\d]+=[\w\d]+)*)*)*$`)
)

func StrPtrEmptyNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func ParseDSN(dsn string) *DSN {
	match := dsnPattern.FindStringSubmatch(dsn)
	if len(match) == 0 {
		return nil
	}
	result := make(map[string]string)
	for i, name := range dsnPattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	username := StrPtrEmptyNil(result["username"])
	password := StrPtrEmptyNil(result["password"])
	port := StrPtrEmptyNil(result["port"])
	opts := StrPtrEmptyNil(result["opts"])

	return &DSN{
		Protocol: result["proto"],
		User:     username,
		Password: password,
		Host:     result["host"],
		Port:     port,
		DbName:   result["db_name"],
		Opts:     opts,
	}
}
