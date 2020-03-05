package ucloud

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

func dataSourceUCloudDisks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUCloudDisksRead,

		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:      schema.HashString,
				Computed: true,
			},

			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.ValidateRegexp,
			},

			"disk_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"data_disk",
					"ssd_data_disk",
					"system_disk",
					"ssd_system_disk",
					"rssd_data_disk",
				}, false),
			},

			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"disks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability_zone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"disk_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"disk_type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"charge_type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"expire_time": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceUCloudDisksRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*UCloudClient).udiskconn
	var allDisks []udisk.UDiskDataSet
	var disks []udisk.UDiskDataSet
	var limit int = 100
	var offset int

	for {
		req := conn.NewDescribeUDiskRequest()
		req.Limit = ucloud.Int(limit)
		req.Offset = ucloud.Int(offset)
		if v, ok := d.GetOk("disk_type"); ok {
			req.DiskType = ucloud.String(diskTypeCvt.unconvert(v.(string)))
		}

		if v, ok := d.GetOk("availability_zone"); ok {
			req.Zone = ucloud.String(v.(string))
		}

		resp, err := conn.DescribeUDisk(req)
		if err != nil {
			return fmt.Errorf("error on reading disk list, %s", err)
		}

		if resp == nil || len(resp.DataSet) < 1 {
			break
		}

		allDisks = append(allDisks, resp.DataSet...)

		if len(resp.DataSet) < limit {
			break
		}

		offset = offset + limit
	}

	ids, idsOk := d.GetOk("ids")
	nameRegex, nameRegexOk := d.GetOk("name_regex")
	if idsOk || nameRegexOk {
		var r *regexp.Regexp
		if nameRegex != "" {
			r = regexp.MustCompile(nameRegex.(string))
		}
		for _, v := range allDisks {
			if r != nil && !r.MatchString(v.Name) {
				continue
			}

			if idsOk && !isStringIn(v.UDiskId, schemaSetToStringSlice(ids)) {
				continue
			}
			disks = append(disks, v)
		}
	} else {
		disks = allDisks
	}

	err := dataSourceUCloudDisksSave(d, disks)
	if err != nil {
		return fmt.Errorf("error on reading disk list, %s", err)
	}

	return nil
}

func dataSourceUCloudDisksSave(d *schema.ResourceData, disks []udisk.UDiskDataSet) error {
	ids := []string{}
	data := []map[string]interface{}{}

	for _, item := range disks {
		ids = append(ids, item.UDiskId)

		data = append(data, map[string]interface{}{
			"id":                item.UDiskId,
			"availability_zone": item.Zone,
			"disk_size":         item.Size,
			"disk_type":         diskTypeCvt.convert(item.DiskType),
			"charge_type":       upperCamelCvt.convert(item.ChargeType),
			"name":              item.Name,
			"tag":               item.Tag,
			"status":            item.Status,
			"create_time":       timestampToString(item.CreateTime),
			"expire_time":       timestampToString(item.ExpiredTime),
		})
	}

	d.SetId(hashStringArray(ids))
	d.Set("total_count", len(data))
	d.Set("ids", ids)
	if err := d.Set("disks", data); err != nil {
		return err
	}

	if outputFile, ok := d.GetOk("output_file"); ok && outputFile.(string) != "" {
		writeToFile(outputFile.(string), data)
	}

	return nil
}
