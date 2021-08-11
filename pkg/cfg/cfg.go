package cfg

// Conf is the configuration for accessing bitnodes endpoint
type config struct {
	COMHTTP    []string
	Port       map[string]string
	Out        string
	RPC        RPClogin
	CF         CloudFlare
	ApiKeys    map[string]string
	JDBservers map[string]string
}
type RPClogin struct {
	Username, Password string
}
type CloudFlare struct {
	CloudFlareAPI, CloudFlareEmail, CloudFlareAPIkey, CloudFlareAPItoken string
}

// configurations for jorm
var (
	Path   = "/var/db/jorm/"
	C      *config
	CFG, _ = NewCFG(Path, nil)
)
