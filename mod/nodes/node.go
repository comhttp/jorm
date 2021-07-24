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

type Nodes []Node

// Node stores info retrieved via geoip about a node
type Node struct {
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
	bns := make(map[string]*BitNoded)
	for _, coin := range coins.C {
		bn := &BitNoded{}

		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin.Slug)) {
			b = append(b, coin.Slug)
			bitNodes := BitNodes{}
			err := j.Read("nodes", coin.Slug, &bitNodes)
			utl.ErrorLog(err)

			for _, bitnode := range coin.Nodes {
				bitnode.getNode(j, bn, coin.Slug)
			}
			bns[coin.Slug] = bn
			j.Write("nodes", coin.Slug+"_"+"bitnodes", bn)

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

func (b *BitNode) getNode(j *jdb.JDB, bn *BitNoded, c string) {
	b.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, b.IP, b.Port)
	s := b.GetBitNodeStatus()
	j.Write("nodes", c+"_"+b.IP, s)

	j.Write("info", c+"_mempool", s.GetRawMemPool)
	j.Write("info", c+"_mining", s.GetInfo)
	j.Write("info", c+"_info", s.GetInfo)
	j.Write("info", c+"_network", s.GetNetworkInfo)
	j.Write("info", c+"_peers", s.GetPeerInfo)

	fmt.Println("GetBitNodeStatus: ", c+"_"+b.IP)
	nds := GetNodes(s)
	for _, n := range nds {
		j.Write("nodes", c+"_"+n.IP, n)
		fmt.Println("Node: ", c+"_"+n.IP)

	}

	bn.Coin = c
	bn.BitNodes = append(bn.BitNodes, *s)
	j.Write("nodes", c+"_"+b.IP, s)
	return
}
