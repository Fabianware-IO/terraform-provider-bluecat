package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceDNS :: dns provider resource
func resourceDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSCreate,
		Read:   resourceDNSRead,
		Update: resourceDNSUpdate,
		Delete: resourceDNSDelete,
		Exists: resourceDNSExists,

		Schema: map[string]*schema.Schema{},
	}
}

// resourceDNSExists :: check if dns resource exists
func resourceDNSExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

// resourceDNSCreate :: create a new dns entry
func resourceDNSCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceDNSRead :: read a dns entry
func resourceDNSRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceDNSUpdate :: update a dns entry
func resourceDNSUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceDNSDelete :: delete a dns entry
func resourceDNSDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
