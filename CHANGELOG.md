## 1.8.0 (Unreleased)

ENHANCEMENTS:

* datasource/ucloud_disks: Add attribute `availability_zone`[GH-30]
* datasource/ucloud_disks: Add possible value `rssd_data_disk` of `disk_type` [GH-30]
* resource/ucloud_disk: Add possible value `rssd_data_disk` of `disk_type` [GH-30]
* resource/ucloud_instance: Optimization of validate about `instance_type` [GH-30]
* resource/ucloud_instance: Add Outstanding `instance_type` [GH-30]

## 1.7.0 (May 10, 2019)

ENHANCEMENTS:

* datasource/ucloud_instances: Add attribute `vpc_id `, `subnet_id`, `private_ip` ([#27](https://github.com/terraform-providers/terraform-provider-ucloud/issues/27))
* resource/ucloud_instance: Add attribute `private_ip` ([#27](https://github.com/terraform-providers/terraform-provider-ucloud/issues/27))
* resource/ucloud_lb_listener: Update customdiff precheck about `protocol` and `listen_type` ([#27](https://github.com/terraform-providers/terraform-provider-ucloud/issues/27))

BUG FIXES:

* resource/ucloud_vpc: Fix `cidr_blocks` validate func ([#28](https://github.com/terraform-providers/terraform-provider-ucloud/issues/28))
* resource/ucloud_subnet: Fix `cidr_block` validate func ([#28](https://github.com/terraform-providers/terraform-provider-ucloud/issues/28))
* resource/ucloud_instance: Fix `image_id` read ([#28](https://github.com/terraform-providers/terraform-provider-ucloud/issues/28))

## 1.6.0 (April 12, 2019)

FEATURES:

* **New Datasource:** `ucloud_db_instances`([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* **New Datasource:** `ucloud_lb_ssls`([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* **New Datasource:** `ucloud_security_groups`([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* **New Datasource:** `ucloud_vpcs`([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* **New Datasource:** `ucloud_subnets`([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))

ENHANCEMENTS:

* datasource/ucloud_lbs: Add attribute `internal` ([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* datasource/ucloud_instances: Add argument `name_regex` ([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* datasource/ucloud_eips: Add argument `name_regex` ([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* datasource/ucloud_projects: Add argument `name_regex` ([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))
* datasource/ucloud_zones: Add attribute `total_count` ([#24](https://github.com/terraform-providers/terraform-provider-ucloud/issues/24))

## 1.5.0 (April 01, 2019)

FEATURES:

* **New Datasource:** `ucloud_disks`([#23](https://github.com/terraform-providers/terraform-provider-ucloud/issues/23))
* **New Datasource:** `ucloud_lbs`([#23](https://github.com/terraform-providers/terraform-provider-ucloud/issues/23))
* **New Datasource:** `ucloud_lb_listeners`([#23](https://github.com/terraform-providers/terraform-provider-ucloud/issues/23))
* **New Datasource:** `ucloud_lb_rules`([#23](https://github.com/terraform-providers/terraform-provider-ucloud/issues/23))
* **New Datasource:** `ucloud_lb_attachments`([#23](https://github.com/terraform-providers/terraform-provider-ucloud/issues/23))

ENHANCEMENTS:

* resource/ucloud_lb: Deprecated attribute `expire_time` for optimizing outputs ([#23](https://github.com/terraform-providers/terraform-provider-ucloud/issues/23))

## 1.4.0 (March 18, 2019)

ENHANCEMENTS:

* resource/ucloud_db_instance: Shorten the waiting time ([#22](https://github.com/terraform-providers/terraform-provider-ucloud/issues/22))
* resource/ucloud_disk: Shorten the waiting time and update states ([#22](https://github.com/terraform-providers/terraform-provider-ucloud/issues/22))
* resource/ucloud_disk_attachment: Shorten the waiting time ([#22](https://github.com/terraform-providers/terraform-provider-ucloud/issues/22))
* resource/ucloud_eip: Shorten the waiting time and update states ([#22](https://github.com/terraform-providers/terraform-provider-ucloud/issues/22))
* resource/ucloud_lb_listener: Shorten the waiting time and update states ([#22](https://github.com/terraform-providers/terraform-provider-ucloud/issues/22))
* resource/ucloud_lb_attachment: Shorten the waiting time and update states ([#22](https://github.com/terraform-providers/terraform-provider-ucloud/issues/22))

## 1.3.1 (March 15, 2019)

BUG FIXES:

* resource/ucloud_lb_listener: Fix lb listener import ([#21](https://github.com/terraform-providers/terraform-provider-ucloud/issues/21))
* resource/ucloud_lb_attachment: Fix lb attachment import ([#21](https://github.com/terraform-providers/terraform-provider-ucloud/issues/21))
* resource/ucloud_lb_rule: Fix lb rule import ([#21](https://github.com/terraform-providers/terraform-provider-ucloud/issues/21))

## 1.3.0 (March 12, 2019)

ENHANCEMENTS:

* resource/ucloud_db_instance: Add default password ([#18](https://github.com/ucloud/terraform-provider-ucloud/issues/18))
* resource/ucloud_lb: Deprecated `charge_type` ([#18](https://github.com/ucloud/terraform-provider-ucloud/issues/18))

BUG FIXES:

* resource/ucloud_lb: Fix lb import about `charge_type` and `internal` ([#18](https://github.com/ucloud/terraform-provider-ucloud/issues/18))

## 1.2.1 (March 06, 2019)

ENHANCEMENTS:

* resource/ucloud_instance: Add default root password ([#15](https://github.com/terraform-providers/terraform-provider-ucloud/issues/15))

BUG FIXES:

* resource/ucloud_instance: Fix validate cloud disk import ([#15](https://github.com/terraform-providers/terraform-provider-ucloud/issues/15))

## 1.2.0 (March 05, 2019)

FEATURES:

* **New Resource:** `ucloud_db_instance` ([#12](https://github.com/terraform-providers/terraform-provider-ucloud/issues/12))
* **New Resource:** `ucloud_lb_ssl` ([#12](https://github.com/terraform-providers/terraform-provider-ucloud/issues/12))
* **New Resource:** `ucloud_lb_ssl_attachment` ([#12](https://github.com/terraform-providers/terraform-provider-ucloud/issues/12))
* **New Datasource:** `ucloud_instances` ([#12](https://github.com/terraform-providers/terraform-provider-ucloud/issues/12))
* **New Resource:** `ucloud_udpn_connection` ([#7](https://github.com/terraform-providers/terraform-provider-ucloud/issues/7))

ENHANCEMENTS:

* resource/ucloud_disk_attachment: Update schema version for disk attachment ID ([#12](https://github.com/terraform-providers/terraform-provider-ucloud/issues/12))
* resource/ucloud_vpc: Add update logic to `cidr_blocks` ([#9](https://github.com/terraform-providers/terraform-provider-ucloud/issues/9))
* provider: Support shared credential file and named profile ([#11](https://github.com/terraform-providers/terraform-provider-ucloud/issues/11))
* provider: Support customize endpoint url ([#11](https://github.com/terraform-providers/terraform-provider-ucloud/issues/11))

BUG FIXES:

* resource/ucloud_instance: Fix read of `image_id` and `instance_type` ([#12](https://github.com/terraform-providers/terraform-provider-ucloud/issues/12))
* resource/ucloud_instance: Check and create default firewall for new account ([#9](https://github.com/terraform-providers/terraform-provider-ucloud/issues/9))
* resource/ucloud_vpc: Fix cannot add multi value to `cidr_blocks` ([#9](https://github.com/terraform-providers/terraform-provider-ucloud/issues/9))

## 1.1.0 (January 09, 2019)

ENHANCEMENTS:

* resource/ucloud_eip_association: Update schema version for eip association `ID` ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_eip_association: Deprecated `resource_type` ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_lb_attachment: Deprecated `resource_type` ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_eip: Add `public_ip` attribute ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_instance: Update `instance_type` about customized ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* provider: Add `UserAgent` to external API ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))

BUG FIXES:

* resource/ucloud_disk: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_eip: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_instance: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_lb_listener: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_lb: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_security_group: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_subnet: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))
* resource/ucloud_vpc: Fix default of `name` argument ([#2](https://github.com/terraform-providers/terraform-provider-ucloud/issues/2))

## 1.0.0 (December 19, 2018)

FEATURES:

* **New Resource:** `ucloud_instance`
* **New Resource:** `ucloud_disk`
* **New Resource:** `ucloud_disk_attachment`
* **New Resource:** `ucloud_eip`
* **New Resource:** `ucloud_eip_association`
* **New Resource:** `ucloud_security_group`
* **New Resource:** `ucloud_vpc`
* **New Resource:** `ucloud_subnet`
* **New Resource:** `ucloud_vpc_peering_connection`
* **New Resource:** `ucloud_lb`
* **New Resource:** `ucloud_lb_listener`
* **New Resource:** `ucloud_lb_attachment`
* **New Resource:** `ucloud_lb_rule`
* **New Datasource:** `ucloud_eips`
* **New Datasource:** `ucloud_images`
* **New Datasource:** `ucloud_projects`
* **New Datasource:** `ucloud_zones`
