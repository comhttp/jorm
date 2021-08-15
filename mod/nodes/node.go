package nodes

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
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
func GetBitNodes(j *jdb.JDB, bitNodeCoins []string) {
	var b []string
	bns := make(map[string]*BitNoded)
	//log.Print("bitNodeCoinsbitNodeCoinsbitNodeCoinsbitNodeCoins", bitNodeCoins)
	for _, coin := range bitNodeCoins {
		bn := &BitNoded{}
		b = append(b, coin)
		//c := coins.GetCoin(j, coin)

		//for _, bitnode := range c.Nodes {
		//	bitnode.getNode(j, bn, coin)
		//}
		bns[coin] = bn
		j.Write("nodes", coin+"_"+"bitnodes", bn)
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

func (b *BitNode) getNode(j *jdb.JDB, bn *BitNoded, username, password, coin string) {
	b.Jrc = utl.NewClient(username, password, b.IP, b.Port)
	s := b.GetBitNodeStatus(username, password)
	j.Write("nodes", coin+"_"+b.IP, s)

	j.Write("info", coin+"_mempool", s.GetRawMemPool)
	j.Write("info", coin+"_mining", s.GetInfo)
	j.Write("info", coin+"_info", s.GetInfo)
	j.Write("info", coin+"_network", s.GetNetworkInfo)
	j.Write("info", coin+"_peers", s.GetPeerInfo)

	log.Print("GetBitNodeStatus: ", coin+"_"+b.IP)
	nds := GetNodes(s)
	for _, n := range nds {
		j.Write("nodes", coin+"_"+n.IP, n)
		log.Print("Node: ", coin+"_"+n.IP)

	}

	bn.Coin = coin
	bn.BitNodes = append(bn.BitNodes, *s)
	j.Write("nodes", coin+"_"+b.IP, s)
	return
}
