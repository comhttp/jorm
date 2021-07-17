package nodes

import (
	"fmt"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"path/filepath"
)

type BitNode struct {
	IP   string `json:"ip"`
	Port int64  `json:"p"`
	Jrc  *utl.Endpoint
}

// BitNodes is array of bitnodes addresses
type BitNodes []BitNode

// BitNoded data
type BitNoded struct {
	Coin     string          `json:"coin"`
	BitNodes []BitNodeStatus `json:"bitnodes"`
}

// Coin stores identifying information about coins in the database
type NodeCoin struct {
	Rank   int      `json:"r"`
	Name   string   `json:"n"`
	Ticker string   `json:"t"`
	Slug   string   `json:"s"`
	Algo   string   `json:"a"`
	Nodes  BitNodes `json:"b"`
}

type NodeCoins struct {
	N int        `json:"n"`
	C []NodeCoin `json:"c"`
}

// NodeStatus stores current data for a node
type BitNodeStatus struct {
	Live           bool        `json:"live"`
	IP             string      `json:"ip"`
	GetInfo        interface{} `json:"getinfo"`
	GetPeerInfo    interface{} `json:"getpeerinfo"`
	GetRawMemPool  interface{} `json:"getrawmempool"`
	GetMiningInfo  interface{} `json:"getmininginfo"`
	GetNetworkInfo interface{} `json:"getnetworkinfo"`
	GeoIP          interface{} `json:"geoip"`
}

type Nodes []NodeInfo

// NodeInfo stores info retrieved via geoip about a node
type NodeInfo struct {
	IP            string  `json:"ip"`
	Port          int64   `json:"port"`
	Host          string  `json:"host"`
	Rdns          string  `json:"rdns"`
	ASN           string  `json:"asn"`
	ISP           string  `json:"isp"`
	CountryName   string  `json:"countryname"`
	CountryCode   string  `json:"countrycode"`
	RegionName    string  `json:"regionname"`
	RegionCode    string  `json:"regioncode"`
	City          string  `json:"city"`
	Postcode      string  `json:"postcode"`
	ContinentName string  `json:"continentname"`
	ContinentCode string  `json:"continentcode"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Zipcode       string  `json:"zipcode"`
	Timezone      string  `json:"timezone"`
	LastSeen      string  `json:"lastseen"`
	Live          bool    `json:"live"`
}

//// GetBitNodes updates the data about all of the coins in the database
func GetBitNodes(j *jdb.JDB, coins NodeCoins) {
	var b []string
	bns := make(map[string]BitNoded)
	for _, coin := range coins.C {
		var bn BitNoded

		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin.Slug)) {
			b = append(b, coin.Slug)
			//bitNodes := BitNodes{}
			//if err := cfg.CFG.Read("nodes", coin.Slug, &bitNodes); err != nil {
			//	fmt.Println("Error", err)
			//}
			for _, bitnode := range coin.Nodes {
				bitnode.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, bitnode.IP, bitnode.Port)

				s := bitnode.GetBitNodeStatus()
				//j.Write("nodes", coin.Slug+"_"+bitnode.IP, bitnode.GetBitNodeStatus())

				j.Write("info", coin.Slug+"_mempool", s.GetRawMemPool)
				j.Write("info", coin.Slug+"_mining", s.GetInfo)
				j.Write("info", coin.Slug+"_info", s.GetInfo)
				j.Write("info", coin.Slug+"_network", s.GetNetworkInfo)
				j.Write("info", coin.Slug+"_peers", s.GetPeerInfo)

				fmt.Println("GetBitNodeStatus: ", coin.Slug+"_"+bitnode.IP)
				//	nds := GetNodes(bitNode)
				//	for _, n := range nds {
				//		if n.IP[:3] == "10." {
				//			n.IP = "212.62.35.158"
				//		}
				//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/nodes/"+coin), n.IP, n)
				//	}
				//	if bitnode.IP[:3] == "10." {
				//		bitnode.IP = "212.62.35.158"
				//	}
				//	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/bitnodes/"+coin), bitnode.IP, bitNode)

				//	//bn.Coin = coin
				//	bn.BitNodes = append(bn.BitNodes, *bitNode)
			}
			bns[coin.Slug] = bn
			//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info/nodes/"+coin), "bitnodes", bn)

			//data, err := jdb.JDB.ReadAll(filepath.FromSlash(cfg.C.Out + "/nodes/" + coin))
			//utl.ErrorLog(err)
			//nodes := make([][]byte, len(data))
			//for i := range data {
			//	nodes[i] = []byte(data[i])
			//}

			//ns := make(Nodes, len(nodes))
			//
			//for i := range nodes {
			//	if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
			//		fmt.Println("Error", err)
			//	}
			//}
			//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info/nodes/"+coin), "nodes", ns)
		}
	}

	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info"), "bitnoded", b)
	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info"), "bitnodestat", bns)
}

func GetNode(j *jdb.JDB, c, ip string) map[string]interface{} {
	node := make(map[string]interface{})
	err := j.Read("nodes", c+"_"+ip, &node)
	utl.ErrorLog(err)
	return node
}
