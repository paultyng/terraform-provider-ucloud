package ucloud

import (
	"fmt"

	"github.com/hashicorp/terraform/terraform"
)

func resourceUCloudDiskAttachmentMigrateState(
	v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		return migrateUCloudDiskAttachmentV0toV1(is)
	default:
		return is, fmt.Errorf("unexpected schema version: %d", v)
	}
}

func migrateUCloudDiskAttachmentV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() || is.Attributes == nil {
		return is, nil
	}

	ai, err := parseAssociationInfo(is.ID)
	if err != nil {
		return is, err
	}

	is.ID = fmt.Sprintf("%s:%s", ai.PrimaryId, ai.ResourceId)

	return is, nil
}
