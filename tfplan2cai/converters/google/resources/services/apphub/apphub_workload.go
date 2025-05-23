// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apphub/Workload.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apphub

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApphubWorkloadAssetType string = "apphub.googleapis.com/Workload"

func ResourceConverterApphubWorkload() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApphubWorkloadAssetType,
		Convert:   GetApphubWorkloadCaiObject,
	}
}

func GetApphubWorkloadCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apphub.googleapis.com/projects/{{project}}/locations/{{location}}/applications/{{application_id}}/workloads/{{workload_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApphubWorkloadApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApphubWorkloadAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apphub/v1/rest",
				DiscoveryName:        "Workload",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApphubWorkloadApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandApphubWorkloadDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApphubWorkloadDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	discoveredWorkloadProp, err := expandApphubWorkloadDiscoveredWorkload(d.Get("discovered_workload"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("discovered_workload"); !tpgresource.IsEmptyValue(reflect.ValueOf(discoveredWorkloadProp)) && (ok || !reflect.DeepEqual(v, discoveredWorkloadProp)) {
		obj["discoveredWorkload"] = discoveredWorkloadProp
	}
	attributesProp, err := expandApphubWorkloadAttributes(d.Get("attributes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attributes"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributesProp)) && (ok || !reflect.DeepEqual(v, attributesProp)) {
		obj["attributes"] = attributesProp
	}

	return obj, nil
}

func expandApphubWorkloadDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadDiscoveredWorkload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCriticality, err := expandApphubWorkloadAttributesCriticality(original["criticality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCriticality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["criticality"] = transformedCriticality
	}

	transformedEnvironment, err := expandApphubWorkloadAttributesEnvironment(original["environment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnvironment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["environment"] = transformedEnvironment
	}

	transformedDeveloperOwners, err := expandApphubWorkloadAttributesDeveloperOwners(original["developer_owners"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeveloperOwners); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["developerOwners"] = transformedDeveloperOwners
	}

	transformedOperatorOwners, err := expandApphubWorkloadAttributesOperatorOwners(original["operator_owners"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOperatorOwners); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["operatorOwners"] = transformedOperatorOwners
	}

	transformedBusinessOwners, err := expandApphubWorkloadAttributesBusinessOwners(original["business_owners"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBusinessOwners); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["businessOwners"] = transformedBusinessOwners
	}

	return transformed, nil
}

func expandApphubWorkloadAttributesCriticality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandApphubWorkloadAttributesCriticalityType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandApphubWorkloadAttributesCriticalityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesEnvironment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandApphubWorkloadAttributesEnvironmentType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandApphubWorkloadAttributesEnvironmentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesDeveloperOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandApphubWorkloadAttributesDeveloperOwnersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedEmail, err := expandApphubWorkloadAttributesDeveloperOwnersEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApphubWorkloadAttributesDeveloperOwnersDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesDeveloperOwnersEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesOperatorOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandApphubWorkloadAttributesOperatorOwnersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedEmail, err := expandApphubWorkloadAttributesOperatorOwnersEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApphubWorkloadAttributesOperatorOwnersDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesOperatorOwnersEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesBusinessOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandApphubWorkloadAttributesBusinessOwnersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedEmail, err := expandApphubWorkloadAttributesBusinessOwnersEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApphubWorkloadAttributesBusinessOwnersDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubWorkloadAttributesBusinessOwnersEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
