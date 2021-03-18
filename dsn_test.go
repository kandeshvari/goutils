package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dsnSuit = [][]interface{}{
	{
		"postgres://user:password@localhost:15432/tilka_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     StrPtrEmptyNil("user"),
			Password: StrPtrEmptyNil("password"),
			Host:     "localhost",
			Port:     StrPtrEmptyNil("15432"),
			DbName:   "tilka_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://user@localhost:15432/tilka_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     StrPtrEmptyNil("user"),
			Password: nil,
			Host:     "localhost",
			Port:     StrPtrEmptyNil("15432"),
			DbName:   "tilka_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://localhost:15432/tilka_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     StrPtrEmptyNil("15432"),
			DbName:   "tilka_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://localhost/tilka_dev_db?sslmode=disable&opt=2&x=d",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     nil,
			DbName:   "tilka_dev_db",
			Opts:     StrPtrEmptyNil("sslmode=disable&opt=2&x=d"),
		},
	},
	{
		"postgres://localhost/tilka_dev_db",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     nil,
			DbName:   "tilka_dev_db",
			Opts:     nil,
		},
	},
	{
		"postgres://localhost/tilka_dev_db/scope",
		&DSN{
			Protocol: "postgres",
			User:     nil,
			Password: nil,
			Host:     "localhost",
			Port:     nil,
			DbName:   "tilka_dev_db/scope",
			Opts:     nil,
		},
	},
}

func TestParseDSN(t *testing.T) {
	for _, s := range dsnSuit {
		dsn := ParseDSN(s[0].(string))
		assert.Equal(t, s[1].(*DSN), dsn)
	}
}
