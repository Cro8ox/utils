package status

import "github.com/cro8ox/gophercloud"

const resourcePath = "status"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func getURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}
