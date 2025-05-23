// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/ResizeRequest.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeResizeRequestAssetType string = "compute.googleapis.com/ResizeRequest"

func ResourceConverterComputeResizeRequest() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeResizeRequestAssetType,
		Convert:   GetComputeResizeRequestCaiObject,
	}
}

func GetComputeResizeRequestCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/resizeRequests/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeResizeRequestApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeResizeRequestAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "ResizeRequest",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeResizeRequestApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeResizeRequestName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeResizeRequestDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	resizeByProp, err := expandComputeResizeRequestResizeBy(d.Get("resize_by"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resize_by"); !tpgresource.IsEmptyValue(reflect.ValueOf(resizeByProp)) && (ok || !reflect.DeepEqual(v, resizeByProp)) {
		obj["resizeBy"] = resizeByProp
	}
	requestedRunDurationProp, err := expandComputeResizeRequestRequestedRunDuration(d.Get("requested_run_duration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("requested_run_duration"); !tpgresource.IsEmptyValue(reflect.ValueOf(requestedRunDurationProp)) && (ok || !reflect.DeepEqual(v, requestedRunDurationProp)) {
		obj["requestedRunDuration"] = requestedRunDurationProp
	}
	zoneProp, err := expandComputeResizeRequestZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	return obj, nil
}

func expandComputeResizeRequestName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResizeRequestDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResizeRequestResizeBy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResizeRequestRequestedRunDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSeconds, err := expandComputeResizeRequestRequestedRunDurationSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandComputeResizeRequestRequestedRunDurationNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandComputeResizeRequestRequestedRunDurationSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResizeRequestRequestedRunDurationNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResizeRequestZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
