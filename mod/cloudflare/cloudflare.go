package cloudflare

import (
	"context"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"log"

	cf "github.com/cloudflare/cloudflare-go"
)

func CloudFlare(j *jdb.JDB) {
	//log.Println("CONFIGCONFIGCONFIGCONFIGCONFIGCONFIGCONFIG", cfg.C)
	ctx := context.Background()
	// Construct a new API object
	api, err := cf.NewWithAPIToken(cfg.C.CF.CloudFlareAPItoken)
	utl.ErrorLog(err)
	for _, tld := range cfg.C.COMHTTP {
		go createDNS(j, api, ctx, "com-http."+tld)
	}
	//createDNS(j,api, ctx, "com-http.us")
	//delAllCNameDNS(api, ctx, "com-http.us")
}

func createDNS(j *jdb.JDB, api *cf.API, ctx context.Context, domain string) {
	c := coins.GetAlgoCoins(j)
	// Fetch the zone ID
	id, err := api.ZoneIDByName(domain) // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}
	// Fetch all records for a zone
	recs, err := api.DNSRecords(context.Background(), id, cf.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}
	var registrated []string
	for _, r := range recs {
		if r.Type == "CNAME" {
			registrated = append(registrated, r.Name)
		}
	}
	for _, coin := range c.C {
		//_, err := http.Get("https://" + slug + "." + domain)
		//if err != nil {
		setDNS(api, ctx, registrated, domain, coin.Slug)
	}
}

func setDNS(api *cf.API, ctx context.Context, registrated []string, domain, slug string) {
	var exist bool
	for _, reg := range registrated {
		if slug+"."+domain == reg {
			log.Println("Ima:", slug+"."+domain)
			exist = true
		} else {
			exist = false
		}
	}
	if !exist {
		id, err := api.ZoneIDByName(domain)
		utl.ErrorLog(err)
		t := true
		_, err = api.CreateDNSRecord(ctx, id, cf.DNSRecord{
			Type:    "CNAME",
			Name:    slug,
			Content: domain,
			TTL:     1,
			Proxied: &t,
		})
		utl.ErrorLog(err)
		log.Println("Created subdomain: ", slug+"."+domain)
	}
}

func delAllCNameDNS(api *cf.API, ctx context.Context, domain string) {
	// Fetch the zone ID
	id, err := api.ZoneIDByName(domain) // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}
	// Fetch all records for a zone
	recs, err := api.DNSRecords(context.Background(), id, cf.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range recs {
		if r.Type == "CNAME" {
			go delDNS(api, ctx, id, r.ID)
		}
	}
}

func delDNS(api *cf.API, ctx context.Context, zoneId, id string) {
	err := api.DeleteDNSRecord(ctx, zoneId, id)
	utl.ErrorLog(err)
	log.Println("DeleteDNSRecord rrrrr:", id)
}
