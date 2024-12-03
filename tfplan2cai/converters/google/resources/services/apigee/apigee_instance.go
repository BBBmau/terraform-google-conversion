// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Suppress diffs when the lists of project have the same number of entries to handle the case that
// API does not return what the user originally provided. Instead, API does some transformation.
// For example, user provides a list of project number, but API returns a list of project Id.
func projectListDiffSuppress(_, _, _ string, d *schema.ResourceData) bool {
	return ProjectListDiffSuppressFunc(d)
}

func ProjectListDiffSuppressFunc(d tpgresource.TerraformResourceDataChange) bool {
	kLength := "consumer_accept_list.#"
	oldLength, newLength := d.GetChange(kLength)

	oldInt, ok := oldLength.(int)
	if !ok {
		return false
	}

	newInt, ok := newLength.(int)
	if !ok {
		return false
	}
	log.Printf("[DEBUG] - suppressing diff with oldInt %d, newInt %d", oldInt, newInt)

	return oldInt == newInt
}

const ApigeeInstanceAssetType string = "apigee.googleapis.com/Instance"

func ResourceConverterApigeeInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeInstanceAssetType,
		Convert:   GetApigeeInstanceCaiObject,
	}
}

func GetApigeeInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{org_id}}/instances/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationProp, err := expandApigeeInstanceLocation(d.Get("location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	peeringCidrRangeProp, err := expandApigeeInstancePeeringCidrRange(d.Get("peering_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(peeringCidrRangeProp)) && (ok || !reflect.DeepEqual(v, peeringCidrRangeProp)) {
		obj["peeringCidrRange"] = peeringCidrRangeProp
	}
	ipRangeProp, err := expandApigeeInstanceIpRange(d.Get("ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipRangeProp)) && (ok || !reflect.DeepEqual(v, ipRangeProp)) {
		obj["ipRange"] = ipRangeProp
	}
	descriptionProp, err := expandApigeeInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandApigeeInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	diskEncryptionKeyNameProp, err := expandApigeeInstanceDiskEncryptionKeyName(d.Get("disk_encryption_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskEncryptionKeyNameProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyNameProp)) {
		obj["diskEncryptionKeyName"] = diskEncryptionKeyNameProp
	}
	consumerAcceptListProp, err := expandApigeeInstanceConsumerAcceptList(d.Get("consumer_accept_list"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_accept_list"); !tpgresource.IsEmptyValue(reflect.ValueOf(consumerAcceptListProp)) && (ok || !reflect.DeepEqual(v, consumerAcceptListProp)) {
		obj["consumerAcceptList"] = consumerAcceptListProp
	}

	return obj, nil
}

func expandApigeeInstanceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstancePeeringCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceIpRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDiskEncryptionKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceConsumerAcceptList(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
