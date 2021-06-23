package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dsnSuit = [][]interface{}{
	{
		"clickhouse://user_name:pas-!#$%^&*()_+=sword@clickhouse:9000?database=dev_db",
		&DSN{
			Protocol: "clickhouse",
			User:     StrPtrEmptyNil("user_name"),
			Password: StrPtrEmptyNil("pas-!#$%^&*()_+=sword"),
			Host:     "clickhouse",
			Port:     StrPtrEmptyNil("9000"),
			DbName:   "",
			Opts:     StrPtrEmptyNil("database=dev_db"),
		},
	},
	{
		"postgres://user_name:pas!#$%^&*()_+=sword@localhost:15432/app_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     StrPtrEmptyNil("user_name"),
			Password: StrPtrEmptyNil("pas!#$%^&*()_+=sword"),
			Host:     "localhost",
			Port:     StrPtrEmptyNil("15432"),
			DbName:   "app_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://user_name@localhost:15432/app_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     StrPtrEmptyNil("user_name"),
			Password: nil,
			Host:     "localhost",
			Port:     StrPtrEmptyNil("15432"),
			DbName:   "app_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://localhost:15432/app_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     StrPtrEmptyNil("15432"),
			DbName:   "app_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://localhost/app_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     nil,
			DbName:   "app_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://localhost/app_dev_db",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     nil,
			DbName:   "app_dev_db",
			Opts:     nil,
		},
	},
	{
		"postgres://host-with-dash.and.dots/app_dev_db",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "host-with-dash.and.dots",
			Port:     nil,
			DbName:   "app_dev_db",
			Opts:     nil,
		},
	},
	{
		"postgres://192.168.0.1:5432/app_dev_db",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "192.168.0.1",
			Port:     StrPtrEmptyNil("5432"),
			DbName:   "app_dev_db",
			Opts:     nil,
		},
	},
}

var dsnSuitBAD = []string{
	"",
	"invalid",
	"some-invalid-string",
	"postgres://192.168.0.1.0:5432/app_dev_db/scope",
	"postgres://192.168.0.1:54a32/app_dev_db/scope",
	"postgres://192.168.0.1:54-32/app_dev_db/scope",
	"postgres://192.168.0.1:5432/d-ev_db/scope",
	"postgres://192.168.0.1:5432/dev_db/sc-ope",
	"pos-tgres://192.168.0.1:5432/dev_db/sc-ope",
	"postgres://asdfa_asdfas:5432/dev_db/sc-ope",
}

func TestParseDSN(t *testing.T) {
	for _, s := range dsnSuit {
		dsn := ParseDSN(s[0].(string))
		assert.Equal(t, s[1].(*DSN), dsn)
	}
}

func TestInvalidParseDSN(t *testing.T) {
	for _, s := range dsnSuitBAD {
		dsn := ParseDSN(s)
		assert.Nil(t, dsn)
	}
}
