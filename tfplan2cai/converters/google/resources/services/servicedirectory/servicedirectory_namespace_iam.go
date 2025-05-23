// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/servicedirectory/Namespace.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package servicedirectory

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ServiceDirectoryNamespaceIAMAssetType string = "servicedirectory.googleapis.com/Namespace"

func ResourceConverterServiceDirectoryNamespaceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ServiceDirectoryNamespaceIAMAssetType,
		Convert:           GetServiceDirectoryNamespaceIamPolicyCaiObject,
		MergeCreateUpdate: MergeServiceDirectoryNamespaceIamPolicy,
	}
}

func ResourceConverterServiceDirectoryNamespaceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ServiceDirectoryNamespaceIAMAssetType,
		Convert:           GetServiceDirectoryNamespaceIamBindingCaiObject,
		FetchFullResource: FetchServiceDirectoryNamespaceIamPolicy,
		MergeCreateUpdate: MergeServiceDirectoryNamespaceIamBinding,
		MergeDelete:       MergeServiceDirectoryNamespaceIamBindingDelete,
	}
}

func ResourceConverterServiceDirectoryNamespaceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ServiceDirectoryNamespaceIAMAssetType,
		Convert:           GetServiceDirectoryNamespaceIamMemberCaiObject,
		FetchFullResource: FetchServiceDirectoryNamespaceIamPolicy,
		MergeCreateUpdate: MergeServiceDirectoryNamespaceIamMember,
		MergeDelete:       MergeServiceDirectoryNamespaceIamMemberDelete,
	}
}

func GetServiceDirectoryNamespaceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newServiceDirectoryNamespaceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetServiceDirectoryNamespaceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newServiceDirectoryNamespaceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetServiceDirectoryNamespaceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newServiceDirectoryNamespaceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeServiceDirectoryNamespaceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeServiceDirectoryNamespaceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeServiceDirectoryNamespaceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeServiceDirectoryNamespaceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeServiceDirectoryNamespaceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newServiceDirectoryNamespaceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//servicedirectory.googleapis.com/projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ServiceDirectoryNamespaceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchServiceDirectoryNamespaceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ServiceDirectoryNamespaceIamUpdaterProducer,
		d,
		config,
		"//servicedirectory.googleapis.com/projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}",
		ServiceDirectoryNamespaceIAMAssetType,
	)
}
