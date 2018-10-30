package main

import (
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	resty "gopkg.in/resty.v1"
)

func resourceOCCMInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Delete: resourceServerDelete,
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	auth_token, err := getAuthToken()

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken(auth_token).
		Get("/")

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
https: //services.cloud.netapp.com/occm/list-occms
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func registerCloudManager() bool {

}

func getAuthToken() (string, error) {
	return "", nil
}
