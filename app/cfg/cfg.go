package cfg

// Conf is the configuration for accessing bitnodes endpoint
type config struct {
	COMHTTP []string
	Port    string
	Path    string
	Out     string
	RPC     RPClogin
	CF      CloudFlare
	ApiKeys map[string]string
}
type RPClogin struct {
	Username, Password string
}
type CloudFlare struct {
	CloudFlareAPI, CloudFlareEmail, CloudFlareAPIkey string
}

// configurations for jorm
var (
	Path = "/jorm/"
	C    *config
)
