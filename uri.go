package goutils

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"regexp"
)

var (
	urlPattern = regexp.MustCompile(`^((?P<schema>[^:/?#]+):)?(//(?P<host>[^/?#]*))?(?P<path>[^?#]*)(\?(?P<query>[^#]*))?(#(?P<hash>.*))?$`)
)

type URL struct {
	Schema string
	Host   string
	Path   string
	Hash   string
	Query  url.Values
}

// ParseURL parses URL to parts
func ParseURL(urlStr string) *URL {
	res := URL{}
	match := urlPattern.FindStringSubmatch(urlStr)

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
				var err error
				res.Query, err = url.ParseQuery(match[i])
				if err != nil {
					continue
				}

				//if len(match[i]) == 0 {
				//	continue
				//}
				//
				//// TODO: support query arrays
				//pairs := strings.Split(match[i], "&")
				//for i := range pairs {
				//	kv := strings.Split(pairs[i], "=")
				//	if len(kv) != 2 {
				//		log.Warnf("kv len is not 2: %#v", kv)
				//		continue
				//	}
				//	res.Query[kv[0]] = kv[1]
				//}
			default:
				log.Warn("unknown name: ", name)
			}
		}
	}
	return &res
}
