package main

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {

	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"occm_endpoint": {
				Type:        schema.TypeString,
				required:    false,
				Description: "OCCM Endpoint",
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("OCCM_ENDPOINT"); v != "" {
						return v, nil
					}

					return "", nil
				},
			},
			"occm_username": {
				Type:        schema.TypeString,
				required:    false,
				Description: "OCCM Username",
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("OCCM_USERNAME"); v != "" {
						return v, nil
					}

					return "", nil
				},
			},
			"occm_password": {
				Type:        schema.TypeString,
				required:    true,
				Description: "OCCM Password",
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("OCCM_PASSWORD"); v != "" {
						return v, nil
					}

					return "", nil
				},
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"occm_instance": resourceOCCMInstance(),
		},
	}
}
