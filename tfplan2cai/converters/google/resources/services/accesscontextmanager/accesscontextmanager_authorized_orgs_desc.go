// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accesscontextmanager/AuthorizedOrgsDesc.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package accesscontextmanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AccessContextManagerAuthorizedOrgsDescAssetType string = "accesscontextmanager.googleapis.com/AuthorizedOrgsDesc"

func ResourceConverterAccessContextManagerAuthorizedOrgsDesc() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AccessContextManagerAuthorizedOrgsDescAssetType,
		Convert:   GetAccessContextManagerAuthorizedOrgsDescCaiObject,
	}
}

func GetAccessContextManagerAuthorizedOrgsDescCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//accesscontextmanager.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAccessContextManagerAuthorizedOrgsDescApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AccessContextManagerAuthorizedOrgsDescAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accesscontextmanager/v1/rest",
				DiscoveryName:        "AuthorizedOrgsDesc",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAccessContextManagerAuthorizedOrgsDescApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	parentProp, err := expandAccessContextManagerAuthorizedOrgsDescParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerAuthorizedOrgsDescName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	orgsProp, err := expandAccessContextManagerAuthorizedOrgsDescOrgs(d.Get("orgs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("orgs"); !tpgresource.IsEmptyValue(reflect.ValueOf(orgsProp)) && (ok || !reflect.DeepEqual(v, orgsProp)) {
		obj["orgs"] = orgsProp
	}
	assetTypeProp, err := expandAccessContextManagerAuthorizedOrgsDescAssetType(d.Get("asset_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetTypeProp)) && (ok || !reflect.DeepEqual(v, assetTypeProp)) {
		obj["assetType"] = assetTypeProp
	}
	authorizationDirectionProp, err := expandAccessContextManagerAuthorizedOrgsDescAuthorizationDirection(d.Get("authorization_direction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorization_direction"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationDirectionProp)) && (ok || !reflect.DeepEqual(v, authorizationDirectionProp)) {
		obj["authorizationDirection"] = authorizationDirectionProp
	}
	authorizationTypeProp, err := expandAccessContextManagerAuthorizedOrgsDescAuthorizationType(d.Get("authorization_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorization_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationTypeProp)) && (ok || !reflect.DeepEqual(v, authorizationTypeProp)) {
		obj["authorizationType"] = authorizationTypeProp
	}

	return resourceAccessContextManagerAuthorizedOrgsDescEncoder(d, config, obj)
}

func resourceAccessContextManagerAuthorizedOrgsDescEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}

func expandAccessContextManagerAuthorizedOrgsDescParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescOrgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescAssetType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescAuthorizationDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescAuthorizationType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
