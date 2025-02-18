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

package beyondcorp

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const BeyondcorpApplicationIAMAssetType string = "beyondcorp.googleapis.com/Application"

func ResourceConverterBeyondcorpApplicationIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BeyondcorpApplicationIAMAssetType,
		Convert:           GetBeyondcorpApplicationIamPolicyCaiObject,
		MergeCreateUpdate: MergeBeyondcorpApplicationIamPolicy,
	}
}

func ResourceConverterBeyondcorpApplicationIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BeyondcorpApplicationIAMAssetType,
		Convert:           GetBeyondcorpApplicationIamBindingCaiObject,
		FetchFullResource: FetchBeyondcorpApplicationIamPolicy,
		MergeCreateUpdate: MergeBeyondcorpApplicationIamBinding,
		MergeDelete:       MergeBeyondcorpApplicationIamBindingDelete,
	}
}

func ResourceConverterBeyondcorpApplicationIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BeyondcorpApplicationIAMAssetType,
		Convert:           GetBeyondcorpApplicationIamMemberCaiObject,
		FetchFullResource: FetchBeyondcorpApplicationIamPolicy,
		MergeCreateUpdate: MergeBeyondcorpApplicationIamMember,
		MergeDelete:       MergeBeyondcorpApplicationIamMemberDelete,
	}
}

func GetBeyondcorpApplicationIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBeyondcorpApplicationIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetBeyondcorpApplicationIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBeyondcorpApplicationIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetBeyondcorpApplicationIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBeyondcorpApplicationIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeBeyondcorpApplicationIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeBeyondcorpApplicationIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeBeyondcorpApplicationIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeBeyondcorpApplicationIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeBeyondcorpApplicationIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newBeyondcorpApplicationIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//beyondcorp.googleapis.com/projects/{{project}}/locations/global/securityGateways/{{security_gateways_id}}/applications/{{application_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: BeyondcorpApplicationIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchBeyondcorpApplicationIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("security_gateways_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("application_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		BeyondcorpApplicationIamUpdaterProducer,
		d,
		config,
		"//beyondcorp.googleapis.com/projects/{{project}}/locations/global/securityGateways/{{security_gateways_id}}/applications/{{application_id}}",
		BeyondcorpApplicationIAMAssetType,
	)
}
