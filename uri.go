package goutils

import (
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

var (
	urlPattern = regexp.MustCompile(`^((?P<schema>[^:/?#]+):)?(//(?P<host>[^/?#]*))?(?P<path>[^?#]*)(\?(?P<query>[^#]*))?(#(?P<hash>.*))?$`)
)

type URL struct {
	Schema string
	Host   string
	Path   string
	Hash   string
	Query  map[string]string
}

// ParseURL parses URL to parts
func ParseURL(url string) *URL {
	res := URL{Query: make(map[string]string)}
	match := urlPattern.FindStringSubmatch(url)

	for i, name := range urlPattern.SubexpNames() {
		if i != 0 && name != "" {
			switch name {
			case "schema":
				res.Schema = match[i]
			case "host":
				res.Host = match[i]
			case "path":
				res.Path = match[i]
			case "hash":
				res.Hash = match[i]
			case "query":
				if len(match[i]) == 0 {
					continue
				}
				// TODO: support query arrays
				pairs := strings.Split(match[i], "&")
				for i := range pairs {
					kv := strings.Split(pairs[i], "=")
					if len(kv) != 2 {
						log.Warnf("kv len is not 2: %#v", kv)
						continue
					}
					res.Query[kv[0]] = kv[1]
				}
			default:
				log.Warn("unknown name: ", name)
			}
		}
	}
	return &res
}
