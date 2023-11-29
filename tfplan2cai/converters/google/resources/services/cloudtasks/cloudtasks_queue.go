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

package cloudtasks

import (
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func suppressOmittedMaxDuration(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "0s" {
		log.Printf("[INFO] max retry is 0s and api omitted field, suppressing diff")
		return true
	}
	return tpgresource.DurationDiffSuppress(k, old, new, d)
}

const CloudTasksQueueAssetType string = "cloudtasks.googleapis.com/Queue"

func ResourceConverterCloudTasksQueue() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudTasksQueueAssetType,
		Convert:   GetCloudTasksQueueCaiObject,
	}
}

func GetCloudTasksQueueCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudtasks.googleapis.com/projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudTasksQueueApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CloudTasksQueueAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudtasks/v2/rest",
				DiscoveryName:        "Queue",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudTasksQueueApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudTasksQueueName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	appEngineRoutingOverrideProp, err := expandCloudTasksQueueAppEngineRoutingOverride(d.Get("app_engine_routing_override"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine_routing_override"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineRoutingOverrideProp)) && (ok || !reflect.DeepEqual(v, appEngineRoutingOverrideProp)) {
		obj["appEngineRoutingOverride"] = appEngineRoutingOverrideProp
	}
	rateLimitsProp, err := expandCloudTasksQueueRateLimits(d.Get("rate_limits"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rate_limits"); !tpgresource.IsEmptyValue(reflect.ValueOf(rateLimitsProp)) && (ok || !reflect.DeepEqual(v, rateLimitsProp)) {
		obj["rateLimits"] = rateLimitsProp
	}
	retryConfigProp, err := expandCloudTasksQueueRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retry_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(retryConfigProp)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}
	stackdriverLoggingConfigProp, err := expandCloudTasksQueueStackdriverLoggingConfig(d.Get("stackdriver_logging_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("stackdriver_logging_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(stackdriverLoggingConfigProp)) && (ok || !reflect.DeepEqual(v, stackdriverLoggingConfigProp)) {
		obj["stackdriverLoggingConfig"] = stackdriverLoggingConfigProp
	}

	return obj, nil
}

func expandCloudTasksQueueName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
}

func expandCloudTasksQueueAppEngineRoutingOverride(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudTasksQueueAppEngineRoutingOverrideService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandCloudTasksQueueAppEngineRoutingOverrideVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedInstance, err := expandCloudTasksQueueAppEngineRoutingOverrideInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedHost, err := expandCloudTasksQueueAppEngineRoutingOverrideHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	return transformed, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxDispatchesPerSecond, err := expandCloudTasksQueueRateLimitsMaxDispatchesPerSecond(original["max_dispatches_per_second"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDispatchesPerSecond); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDispatchesPerSecond"] = transformedMaxDispatchesPerSecond
	}

	transformedMaxConcurrentDispatches, err := expandCloudTasksQueueRateLimitsMaxConcurrentDispatches(original["max_concurrent_dispatches"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxConcurrentDispatches); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxConcurrentDispatches"] = transformedMaxConcurrentDispatches
	}

	transformedMaxBurstSize, err := expandCloudTasksQueueRateLimitsMaxBurstSize(original["max_burst_size"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBurstSize); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxBurstSize"] = transformedMaxBurstSize
	}

	return transformed, nil
}

func expandCloudTasksQueueRateLimitsMaxDispatchesPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimitsMaxConcurrentDispatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimitsMaxBurstSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxAttempts, err := expandCloudTasksQueueRetryConfigMaxAttempts(original["max_attempts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxAttempts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxAttempts"] = transformedMaxAttempts
	}

	transformedMaxRetryDuration, err := expandCloudTasksQueueRetryConfigMaxRetryDuration(original["max_retry_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetryDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxRetryDuration"] = transformedMaxRetryDuration
	}

	transformedMinBackoff, err := expandCloudTasksQueueRetryConfigMinBackoff(original["min_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinBackoff); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minBackoff"] = transformedMinBackoff
	}

	transformedMaxBackoff, err := expandCloudTasksQueueRetryConfigMaxBackoff(original["max_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBackoff); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxBackoff"] = transformedMaxBackoff
	}

	transformedMaxDoublings, err := expandCloudTasksQueueRetryConfigMaxDoublings(original["max_doublings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDoublings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDoublings"] = transformedMaxDoublings
	}

	return transformed, nil
}

func expandCloudTasksQueueRetryConfigMaxAttempts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxRetryDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMinBackoff(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxBackoff(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxDoublings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueStackdriverLoggingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSamplingRatio, err := expandCloudTasksQueueStackdriverLoggingConfigSamplingRatio(original["sampling_ratio"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplingRatio); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["samplingRatio"] = transformedSamplingRatio
	}

	return transformed, nil
}

func expandCloudTasksQueueStackdriverLoggingConfigSamplingRatio(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
