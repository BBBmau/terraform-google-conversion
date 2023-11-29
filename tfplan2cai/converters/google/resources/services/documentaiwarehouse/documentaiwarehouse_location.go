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

package documentaiwarehouse

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DocumentAIWarehouseLocationAssetType string = "contentwarehouse.googleapis.com/Location"

func ResourceConverterDocumentAIWarehouseLocation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DocumentAIWarehouseLocationAssetType,
		Convert:   GetDocumentAIWarehouseLocationCaiObject,
	}
}

func GetDocumentAIWarehouseLocationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//contentwarehouse.googleapis.com/projects/{{project_number}}/locations/{{location}}:initialize/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDocumentAIWarehouseLocationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DocumentAIWarehouseLocationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/contentwarehouse/v1/rest",
				DiscoveryName:        "Location",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDocumentAIWarehouseLocationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	databaseTypeProp, err := expandDocumentAIWarehouseLocationDatabaseType(d.Get("database_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseTypeProp)) && (ok || !reflect.DeepEqual(v, databaseTypeProp)) {
		obj["databaseType"] = databaseTypeProp
	}
	accessControlModeProp, err := expandDocumentAIWarehouseLocationAccessControlMode(d.Get("access_control_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("access_control_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(accessControlModeProp)) && (ok || !reflect.DeepEqual(v, accessControlModeProp)) {
		obj["accessControlMode"] = accessControlModeProp
	}
	kmsKeyProp, err := expandDocumentAIWarehouseLocationKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	documentCreatorDefaultRoleProp, err := expandDocumentAIWarehouseLocationDocumentCreatorDefaultRole(d.Get("document_creator_default_role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("document_creator_default_role"); !tpgresource.IsEmptyValue(reflect.ValueOf(documentCreatorDefaultRoleProp)) && (ok || !reflect.DeepEqual(v, documentCreatorDefaultRoleProp)) {
		obj["documentCreatorDefaultRole"] = documentCreatorDefaultRoleProp
	}

	return obj, nil
}

func expandDocumentAIWarehouseLocationDatabaseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIWarehouseLocationAccessControlMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIWarehouseLocationKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIWarehouseLocationDocumentCreatorDefaultRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
