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

package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeSnapshotIAMAssetType string = "compute.googleapis.com/Snapshot"

func ResourceConverterComputeSnapshotIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeSnapshotIAMAssetType,
		Convert:           GetComputeSnapshotIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeSnapshotIamPolicy,
	}
}

func ResourceConverterComputeSnapshotIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeSnapshotIAMAssetType,
		Convert:           GetComputeSnapshotIamBindingCaiObject,
		FetchFullResource: FetchComputeSnapshotIamPolicy,
		MergeCreateUpdate: MergeComputeSnapshotIamBinding,
		MergeDelete:       MergeComputeSnapshotIamBindingDelete,
	}
}

func ResourceConverterComputeSnapshotIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeSnapshotIAMAssetType,
		Convert:           GetComputeSnapshotIamMemberCaiObject,
		FetchFullResource: FetchComputeSnapshotIamPolicy,
		MergeCreateUpdate: MergeComputeSnapshotIamMember,
		MergeDelete:       MergeComputeSnapshotIamMemberDelete,
	}
}

func GetComputeSnapshotIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeSnapshotIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetComputeSnapshotIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeSnapshotIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetComputeSnapshotIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeSnapshotIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeComputeSnapshotIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeSnapshotIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeComputeSnapshotIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeComputeSnapshotIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeComputeSnapshotIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newComputeSnapshotIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ComputeSnapshotIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeSnapshotIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ComputeSnapshotIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/global/snapshots/{{name}}",
		ComputeSnapshotIAMAssetType,
	)
}
