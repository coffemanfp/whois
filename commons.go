package whois

import "strings"

func cutHostname(domain string) (resume string) {
	if domain == "" || strings.Count(domain, ".") <= 1 {
		resume = domain
		return
	}

	resume = domain[strings.Index(domain, ".")+1:]
	return
}
