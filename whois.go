package whois

// Whois is a third-party whois API client
type Whois interface {
	ExistsDomain(domain string) (exists string, err error)
}
