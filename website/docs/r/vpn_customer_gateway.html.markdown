---
layout: "ucloud"
page_title: "UCloud: ucloud_vpn_customer_gateway"
sidebar_current: "docs-ucloud-resource-vpn-customer-gateway"
description: |-
  Provides a VPN Customer Gateway resource.
---

# ucloud_vpn_customer_gateway

Provides a VPN Customer Gateway resource.

## Example Usage

```hcl
resource "ucloud_vpn_customer_gateway" "foo" {
    ip_address  = "10.0.0.1"
	name 		= "tf-acc-vpn-customer-gateway-basic"
	tag         = "tf-acc"
}
```

## Argument Reference

The following arguments are supported:

* `ip_address` - (Required) The ip address of the VPN Customer Gateway. 

- - -

* `name` - (Optional) The name of the VPN Customer Gateway which contains 1-63 characters and only support Chinese, English, numbers, '-', '_' and '.'. If not specified, terraform will auto-generate a name beginning with `tf-vpn-customer-gateway-`.
* `remark` - (Optional) The remarks of the VPN Customer Gateway. (Default: `""`).
* `tag` - (Optional) A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
* ``
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `create_time` - The creation time for VPN Customer Gateway, formatted in RFC3339 time string.