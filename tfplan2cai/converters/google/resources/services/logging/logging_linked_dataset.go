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

package logging

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const LoggingLinkedDatasetAssetType string = "logging.googleapis.com/LinkedDataset"

func ResourceConverterLoggingLinkedDataset() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LoggingLinkedDatasetAssetType,
		Convert:   GetLoggingLinkedDatasetCaiObject,
	}
}

func GetLoggingLinkedDatasetCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//logging.googleapis.com/{{parent}}/locations/{{location}}/buckets/{{bucket}}/links/{{link_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetLoggingLinkedDatasetApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: LoggingLinkedDatasetAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/logging/v2/rest",
				DiscoveryName:        "LinkedDataset",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetLoggingLinkedDatasetApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandLoggingLinkedDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return resourceLoggingLinkedDatasetEncoder(d, config, obj)
}

func resourceLoggingLinkedDatasetEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Extract any empty fields from the bucket field.
	parent := d.Get("parent").(string)
	bucket := d.Get("bucket").(string)
	parent, err := ExtractFieldByPattern("parent", parent, bucket, "((projects|folders|organizations|billingAccounts)/[a-z0-9A-Z-]*)/locations/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting parent field: %s", err)
	}
	location := d.Get("location").(string)
	location, err = ExtractFieldByPattern("location", location, bucket, "[a-zA-Z]*/[a-z0-9A-Z-]*/locations/([a-z0-9-]*)/buckets/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting location field: %s", err)
	}
	// Set parent to the extracted value.
	d.Set("parent", parent)
	// Set all the other fields to their short forms before forming url and setting ID.
	bucket = tpgresource.GetResourceNameFromSelfLink(bucket)
	name := d.Get("name").(string)
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("location", location)
	d.Set("bucket", bucket)
	d.Set("name", name)
	return obj, nil
}

func expandLoggingLinkedDatasetDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
