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

package vertexai

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VertexAIIndexEndpointDeployedIndexAssetType string = "aiplatform.googleapis.com/IndexEndpointDeployedIndex"

func ResourceConverterVertexAIIndexEndpointDeployedIndex() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VertexAIIndexEndpointDeployedIndexAssetType,
		Convert:   GetVertexAIIndexEndpointDeployedIndexCaiObject,
	}
}

func GetVertexAIIndexEndpointDeployedIndexCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/{{index_endpoint}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVertexAIIndexEndpointDeployedIndexApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VertexAIIndexEndpointDeployedIndexAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/aiplatform/v1beta1/rest",
				DiscoveryName:        "IndexEndpointDeployedIndex",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVertexAIIndexEndpointDeployedIndexApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	deployedIndexIdProp, err := expandVertexAIIndexEndpointDeployedIndexDeployedIndexId(d.Get("deployed_index_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployed_index_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(deployedIndexIdProp)) && (ok || !reflect.DeepEqual(v, deployedIndexIdProp)) {
		obj["deployedIndexId"] = deployedIndexIdProp
	}
	indexProp, err := expandVertexAIIndexEndpointDeployedIndexIndex(d.Get("index"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("index"); !tpgresource.IsEmptyValue(reflect.ValueOf(indexProp)) && (ok || !reflect.DeepEqual(v, indexProp)) {
		obj["index"] = indexProp
	}
	displayNameProp, err := expandVertexAIIndexEndpointDeployedIndexDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	automaticResourcesProp, err := expandVertexAIIndexEndpointDeployedIndexAutomaticResources(d.Get("automatic_resources"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("automatic_resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(automaticResourcesProp)) && (ok || !reflect.DeepEqual(v, automaticResourcesProp)) {
		obj["automaticResources"] = automaticResourcesProp
	}
	dedicatedResourcesProp, err := expandVertexAIIndexEndpointDeployedIndexDedicatedResources(d.Get("dedicated_resources"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dedicated_resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(dedicatedResourcesProp)) && (ok || !reflect.DeepEqual(v, dedicatedResourcesProp)) {
		obj["dedicatedResources"] = dedicatedResourcesProp
	}
	enableAccessLoggingProp, err := expandVertexAIIndexEndpointDeployedIndexEnableAccessLogging(d.Get("enable_access_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_access_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableAccessLoggingProp)) && (ok || !reflect.DeepEqual(v, enableAccessLoggingProp)) {
		obj["enableAccessLogging"] = enableAccessLoggingProp
	}
	deployedIndexAuthConfigProp, err := expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfig(d.Get("deployed_index_auth_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployed_index_auth_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(deployedIndexAuthConfigProp)) && (ok || !reflect.DeepEqual(v, deployedIndexAuthConfigProp)) {
		obj["deployedIndexAuthConfig"] = deployedIndexAuthConfigProp
	}
	reservedIpRangesProp, err := expandVertexAIIndexEndpointDeployedIndexReservedIpRanges(d.Get("reserved_ip_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_ip_ranges"); !tpgresource.IsEmptyValue(reflect.ValueOf(reservedIpRangesProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangesProp)) {
		obj["reservedIpRanges"] = reservedIpRangesProp
	}
	deploymentGroupProp, err := expandVertexAIIndexEndpointDeployedIndexDeploymentGroup(d.Get("deployment_group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployment_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(deploymentGroupProp)) && (ok || !reflect.DeepEqual(v, deploymentGroupProp)) {
		obj["deploymentGroup"] = deploymentGroupProp
	}

	return resourceVertexAIIndexEndpointDeployedIndexEncoder(d, config, obj)
}

func resourceVertexAIIndexEndpointDeployedIndexEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	req := make(map[string]interface{})
	obj["id"] = d.Get("deployed_index_id")
	delete(obj, "deployedIndexId")
	delete(obj, "name")
	delete(obj, "indexEndpoint")
	req["deployedIndex"] = obj
	return req, nil
}

func expandVertexAIIndexEndpointDeployedIndexDeployedIndexId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexIndex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexAutomaticResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicaCount, err := expandVertexAIIndexEndpointDeployedIndexAutomaticResourcesMinReplicaCount(original["min_replica_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinReplicaCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minReplicaCount"] = transformedMinReplicaCount
	}

	transformedMaxReplicaCount, err := expandVertexAIIndexEndpointDeployedIndexAutomaticResourcesMaxReplicaCount(original["max_replica_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxReplicaCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxReplicaCount"] = transformedMaxReplicaCount
	}

	return transformed, nil
}

func expandVertexAIIndexEndpointDeployedIndexAutomaticResourcesMinReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexAutomaticResourcesMaxReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDedicatedResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMachineSpec, err := expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMachineSpec(original["machine_spec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineSpec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["machineSpec"] = transformedMachineSpec
	}

	transformedMinReplicaCount, err := expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMinReplicaCount(original["min_replica_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinReplicaCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minReplicaCount"] = transformedMinReplicaCount
	}

	transformedMaxReplicaCount, err := expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMaxReplicaCount(original["max_replica_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxReplicaCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxReplicaCount"] = transformedMaxReplicaCount
	}

	return transformed, nil
}

func expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMachineSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMachineType, err := expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMachineSpecMachineType(original["machine_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["machineType"] = transformedMachineType
	}

	return transformed, nil
}

func expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMachineSpecMachineType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMinReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDedicatedResourcesMaxReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexEnableAccessLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAuthProvider, err := expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider(original["auth_provider"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthProvider); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authProvider"] = transformedAuthProvider
	}

	return transformed, nil
}

func expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAudiences, err := expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderAudiences(original["audiences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudiences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audiences"] = transformedAudiences
	}

	transformedAllowedIssuers, err := expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderAllowedIssuers(original["allowed_issuers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedIssuers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedIssuers"] = transformedAllowedIssuers
	}

	return transformed, nil
}

func expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderAudiences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderAllowedIssuers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexReservedIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDeployedIndexDeploymentGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
