package provider

import (
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
)

var mutex = &sync.Mutex{}

// resourceDNS :: dns provider resource
func resourceDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSRecordCreate,
		Read:   resourceDNSRecordRead,
		Update: resourceDNSRecordUpdate,
		Delete: resourceDNSRecordDelete,
		Exists: resourceDNSRecordExists,

		Schema: map[string]*schema.Schema{
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
		},
	}
}

// resourceDNSRecordExists :: check if dns resource exists
func resourceDNSRecordExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	mutex.Lock()

	client := meta.(*bluecat.ProteusApi)

	record := &client.Record{
		Name:  d.Get("name").(string),
		Zone:  d.Get("zone").(string),
		Type:  d.Get("type").(string),
		TTL:   d.Get("ttl").(string),
		Value: d.Get("value").(string),
	}

	switch(*record.type) {
		case "A":
			client.AddHostRecord(record)
		case "SRV":
			client.AddSRVRecord(record)
		case "SOA":
		client.AddStartOfAuthority(record)
		default:
			fmt.PrintLn("not an accepted type")
	}
	return true, nil
}

// resourceDNSRecordCreate :: create a new dns entry
func resourceDNSRecordCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceDNSRecordRead :: read a dns entry
func resourceDNSRecordRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceDNSRecordUpdate :: update a dns entry
func resourceDNSRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceDNSRecordDelete :: delete a dns entry
func resourceDNSRecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
