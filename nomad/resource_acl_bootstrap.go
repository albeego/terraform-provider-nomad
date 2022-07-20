package nomad

import (
	"fmt"
	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceACLBootstrap() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLBootstrapWrite,
		Update: resourceACLBootstrapWrite,
		Delete: resourceACLBootstrapDelete,
		Read:   resourceACLBootstrapRead,
		Exists: resourceACLBootstrapExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"accessor_id": {
				Description: "Token accessor id",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},

			"secret_id": {
				Description: "Token secret id",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},

			"name": {
				Description: "Token name",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},

			"type": {
				Description: "Token type",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},

			"global": {
				Description: "Token is global",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
			},

			"policies": {
				Description: "Token policies",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"create_time": {
				Description: "Token creation time",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},

			"create_index": {
				Description: "Token creation index",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
			},

			"modify_index": {
				Description: "Token modification index",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
			},
		},
	}
}

func resourceACLBootstrapWrite(d *schema.ResourceData, meta interface{}) error {
	providerConfig := meta.(ProviderConfig)
	client := providerConfig.client

	log.Printf("[DEBUG] Bootrapping ACL")
	opts := &api.WriteOptions{
		Namespace: d.Get("namespace").(string),
	}
	if opts.Namespace == "" {
		opts.Namespace = "default"
	}
	token, _, err := client.ACLTokens().Bootstrap(opts)
	if err != nil {
		return fmt.Errorf("error bootstrapping ACL %s", err.Error())
	}
	log.Printf("[DEBUG] Bootstrapped ACL")
	d.SetId(token.AccessorID)

	err = d.Set("accessor_id", token.AccessorID)
	err = d.Set("secret_id", token.SecretID)
	err = d.Set("name", token.Name)
	err = d.Set("type", token.Type)
	err = d.Set("create_index", token.CreateIndex)
	err = d.Set("create_time", token.CreateTime)
	err = d.Set("global", token.Global)
	err = d.Set("modify_index", token.ModifyIndex)
	err = d.Set("policies", token.Policies)

	return nil
}

func resourceACLBootstrapDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceACLBootstrapRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceACLBootstrapExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	return false, nil
}
