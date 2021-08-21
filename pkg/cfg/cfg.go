package cfg

// Conf is the configuration for accessing bitnodes endpoint
type Config struct {
	COMHTTP []string
	Port    map[string]string
	Path    string
	Out     string
	RPC     RPClogin
	CF      CloudFlare
	ApiKeys map[string]string
}

// Conf is the configuration for accessing bitnodes endpoint
type JDBservers map[string]string

type RPClogin struct {
	Username, Password string
}
type CloudFlare struct {
	CloudFlareAPI, CloudFlareEmail, CloudFlareAPIkey, CloudFlareAPItoken string
}
