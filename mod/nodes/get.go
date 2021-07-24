package nodes

import (
	"fmt"
	cfg "github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/utl"
)

// GetBitNodeStatus returns the full set of information about a node
func (b *BitNode) GetBitNodeStatus() (bitnodeStatus *BitNodeStatus) {
	b.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, b.IP, b.Port)
	if b.Jrc != nil {
		getInfo := b.APIGetInfo()
		getPeerInfo := b.APIGetPeerInfo()
		getRawMemPool := b.APIGetRawMemPool()
		getMiningInfo := b.APIGetMiningInfo()
		getNetworkInfo := b.APIGetNetworkInfo()
		if getInfo == nil && getPeerInfo == nil && getRawMemPool == nil && getMiningInfo == nil && getNetworkInfo == nil {
			bitnodeStatus = &BitNodeStatus{
				Live: false,
			}
		} else {
			bitnodeStatus = &BitNodeStatus{
				Live:           true,
				IP:             b.IP,
				GetInfo:        getInfo,
				GetPeerInfo:    getPeerInfo,
				GetRawMemPool:  getRawMemPool,
				GetMiningInfo:  getMiningInfo,
				GetNetworkInfo: getNetworkInfo,
				//GeoIP:          GetGeoIP(b.IP),
			}
		}
	}
	return
}

// GetNodes returns the peers connected to a
func GetNodes(n *BitNodeStatus) (nodes []Node) {
	fmt.Println("GetNodes", n.IP)
	switch p := n.GetPeerInfo.(type) {
	case []interface{}:
		for _, nodeRaw := range p {
			nod := nodeRaw.(map[string]interface{})
			ip, _ := utl.GetIP(nod["addr"].(string))
			// p, _ := strconv.ParseInt(port, 10, 64)
			// n.IP = ip
			// n.Port = p
			node := GetGeoIP(ip)
			// if node != nil {

			nodes = append(nodes, node)
			// }
		}
	}
	return
}
