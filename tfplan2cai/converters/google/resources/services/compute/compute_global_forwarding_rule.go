// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/GlobalForwardingRule.yaml
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

const ComputeGlobalForwardingRuleAssetType string = "compute.googleapis.com/GlobalForwardingRule"

func ResourceConverterComputeGlobalForwardingRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeGlobalForwardingRuleAssetType,
		Convert:   GetComputeGlobalForwardingRuleCaiObject,
	}
}

func GetComputeGlobalForwardingRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeGlobalForwardingRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeGlobalForwardingRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "GlobalForwardingRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeGlobalForwardingRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeGlobalForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeGlobalForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_address"); !tpgresource.IsEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeGlobalForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	ipVersionProp, err := expandComputeGlobalForwardingRuleIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	labelFingerprintProp, err := expandComputeGlobalForwardingRuleLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	loadBalancingSchemeProp, err := expandComputeGlobalForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !tpgresource.IsEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	metadataFiltersProp, err := expandComputeGlobalForwardingRuleMetadataFilters(d.Get("metadata_filters"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata_filters"); !tpgresource.IsEmptyValue(reflect.ValueOf(metadataFiltersProp)) && (ok || !reflect.DeepEqual(v, metadataFiltersProp)) {
		obj["metadataFilters"] = metadataFiltersProp
	}
	nameProp, err := expandComputeGlobalForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeGlobalForwardingRuleNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	portRangeProp, err := expandComputeGlobalForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	subnetworkProp, err := expandComputeGlobalForwardingRuleSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	targetProp, err := expandComputeGlobalForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	networkTierProp, err := expandComputeGlobalForwardingRuleNetworkTier(d.Get("network_tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_tier"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkTierProp)) && (ok || !reflect.DeepEqual(v, networkTierProp)) {
		obj["networkTier"] = networkTierProp
	}
	externalManagedBackendBucketMigrationStateProp, err := expandComputeGlobalForwardingRuleExternalManagedBackendBucketMigrationState(d.Get("external_managed_backend_bucket_migration_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_managed_backend_bucket_migration_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalManagedBackendBucketMigrationStateProp)) && (ok || !reflect.DeepEqual(v, externalManagedBackendBucketMigrationStateProp)) {
		obj["externalManagedBackendBucketMigrationState"] = externalManagedBackendBucketMigrationStateProp
	}
	externalManagedBackendBucketMigrationTestingPercentageProp, err := expandComputeGlobalForwardingRuleExternalManagedBackendBucketMigrationTestingPercentage(d.Get("external_managed_backend_bucket_migration_testing_percentage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_managed_backend_bucket_migration_testing_percentage"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalManagedBackendBucketMigrationTestingPercentageProp)) && (ok || !reflect.DeepEqual(v, externalManagedBackendBucketMigrationTestingPercentageProp)) {
		obj["externalManagedBackendBucketMigrationTestingPercentage"] = externalManagedBackendBucketMigrationTestingPercentageProp
	}
	serviceDirectoryRegistrationsProp, err := expandComputeGlobalForwardingRuleServiceDirectoryRegistrations(d.Get("service_directory_registrations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_directory_registrations"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceDirectoryRegistrationsProp)) && (ok || !reflect.DeepEqual(v, serviceDirectoryRegistrationsProp)) {
		obj["serviceDirectoryRegistrations"] = serviceDirectoryRegistrationsProp
	}
	sourceIpRangesProp, err := expandComputeGlobalForwardingRuleSourceIpRanges(d.Get("source_ip_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_ip_ranges"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceIpRangesProp)) && (ok || !reflect.DeepEqual(v, sourceIpRangesProp)) {
		obj["sourceIpRanges"] = sourceIpRangesProp
	}
	allowPscGlobalAccessProp, err := expandComputeGlobalForwardingRuleAllowPscGlobalAccess(d.Get("allow_psc_global_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_psc_global_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowPscGlobalAccessProp)) && (ok || !reflect.DeepEqual(v, allowPscGlobalAccessProp)) {
		obj["allowPscGlobalAccess"] = allowPscGlobalAccessProp
	}
	noAutomateDnsZoneProp, err := expandComputeGlobalForwardingRuleNoAutomateDnsZone(d.Get("no_automate_dns_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_automate_dns_zone"); ok || !reflect.DeepEqual(v, noAutomateDnsZoneProp) {
		obj["noAutomateDnsZone"] = noAutomateDnsZoneProp
	}
	labelsProp, err := expandComputeGlobalForwardingRuleEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandComputeGlobalForwardingRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIpVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFilters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFilterMatchCriteria, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(original["filter_match_criteria"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilterMatchCriteria); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["filterMatchCriteria"] = transformedFilterMatchCriteria
		}

		transformedFilterLabels, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabels(original["filter_labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilterLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["filterLabels"] = transformedFilterLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeGlobalForwardingRulePortRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeGlobalForwardingRuleTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleNetworkTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleExternalManagedBackendBucketMigrationState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleExternalManagedBackendBucketMigrationTestingPercentage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleServiceDirectoryRegistrations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamespace, err := expandComputeGlobalForwardingRuleServiceDirectoryRegistrationsNamespace(original["namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["namespace"] = transformedNamespace
		}

		transformedServiceDirectoryRegion, err := expandComputeGlobalForwardingRuleServiceDirectoryRegistrationsServiceDirectoryRegion(original["service_directory_region"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceDirectoryRegion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["serviceDirectoryRegion"] = transformedServiceDirectoryRegion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleServiceDirectoryRegistrationsNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleServiceDirectoryRegistrationsServiceDirectoryRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleSourceIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleAllowPscGlobalAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleNoAutomateDnsZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
