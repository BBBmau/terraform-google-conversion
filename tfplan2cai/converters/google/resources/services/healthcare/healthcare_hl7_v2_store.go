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

package healthcare

import (
	"encoding/json"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const HealthcareHl7V2StoreAssetType string = "healthcare.googleapis.com/Hl7V2Store"

func ResourceConverterHealthcareHl7V2Store() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: HealthcareHl7V2StoreAssetType,
		Convert:   GetHealthcareHl7V2StoreCaiObject,
	}
}

func GetHealthcareHl7V2StoreCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//healthcare.googleapis.com/{{dataset}}/hl7V2Stores/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetHealthcareHl7V2StoreApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: HealthcareHl7V2StoreAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/healthcare/v1beta1/rest",
				DiscoveryName:        "Hl7V2Store",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetHealthcareHl7V2StoreApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareHl7V2StoreName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	parserConfigProp, err := expandHealthcareHl7V2StoreParserConfig(d.Get("parser_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parser_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(parserConfigProp)) && (ok || !reflect.DeepEqual(v, parserConfigProp)) {
		obj["parserConfig"] = parserConfigProp
	}
	notificationConfigsProp, err := expandHealthcareHl7V2StoreNotificationConfigs(d.Get("notification_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationConfigsProp)) && (ok || !reflect.DeepEqual(v, notificationConfigsProp)) {
		obj["notificationConfigs"] = notificationConfigsProp
	}
	notificationConfigProp, err := expandHealthcareHl7V2StoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationConfigProp)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}
	labelsProp, err := expandHealthcareHl7V2StoreEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandHealthcareHl7V2StoreName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreParserConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowNullHeader, err := expandHealthcareHl7V2StoreParserConfigAllowNullHeader(original["allow_null_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowNullHeader); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowNullHeader"] = transformedAllowNullHeader
	}

	transformedSegmentTerminator, err := expandHealthcareHl7V2StoreParserConfigSegmentTerminator(original["segment_terminator"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSegmentTerminator); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["segmentTerminator"] = transformedSegmentTerminator
	}

	transformedSchema, err := expandHealthcareHl7V2StoreParserConfigSchema(original["schema"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchema); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schema"] = transformedSchema
	}

	transformedVersion, err := expandHealthcareHl7V2StoreParserConfigVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	return transformed, nil
}

func expandHealthcareHl7V2StoreParserConfigAllowNullHeader(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreParserConfigSegmentTerminator(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreParserConfigSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandHealthcareHl7V2StoreParserConfigVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreNotificationConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPubsubTopic, err := expandHealthcareHl7V2StoreNotificationConfigsPubsubTopic(original["pubsub_topic"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pubsubTopic"] = transformedPubsubTopic
		}

		transformedFilter, err := expandHealthcareHl7V2StoreNotificationConfigsFilter(original["filter"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["filter"] = transformedFilter
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandHealthcareHl7V2StoreNotificationConfigsPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreNotificationConfigsFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreNotificationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandHealthcareHl7V2StoreNotificationConfigPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	return transformed, nil
}

func expandHealthcareHl7V2StoreNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
