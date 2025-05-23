// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/healthcare/ConsentStore.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package healthcare

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const HealthcareConsentStoreAssetType string = "healthcare.googleapis.com/ConsentStore"

func ResourceConverterHealthcareConsentStore() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: HealthcareConsentStoreAssetType,
		Convert:   GetHealthcareConsentStoreCaiObject,
	}
}

func GetHealthcareConsentStoreCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//healthcare.googleapis.com/{{dataset}}/consentStores/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetHealthcareConsentStoreApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: HealthcareConsentStoreAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/healthcare/v1beta1/rest",
				DiscoveryName:        "ConsentStore",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetHealthcareConsentStoreApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	defaultConsentTtlProp, err := expandHealthcareConsentStoreDefaultConsentTtl(d.Get("default_consent_ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_consent_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultConsentTtlProp)) && (ok || !reflect.DeepEqual(v, defaultConsentTtlProp)) {
		obj["defaultConsentTtl"] = defaultConsentTtlProp
	}
	enableConsentCreateOnUpdateProp, err := expandHealthcareConsentStoreEnableConsentCreateOnUpdate(d.Get("enable_consent_create_on_update"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_consent_create_on_update"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableConsentCreateOnUpdateProp)) && (ok || !reflect.DeepEqual(v, enableConsentCreateOnUpdateProp)) {
		obj["enableConsentCreateOnUpdate"] = enableConsentCreateOnUpdateProp
	}
	labelsProp, err := expandHealthcareConsentStoreEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandHealthcareConsentStoreDefaultConsentTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareConsentStoreEnableConsentCreateOnUpdate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareConsentStoreEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
