// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"reflect"
)

func GetComputeHealthCheckApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	nameProp, err := expandComputeHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	timeoutSecProp, err := expandComputeHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}
	httpHealthCheckProp, err := expandComputeHealthCheckHttpHealthCheck(d.Get("http_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_health_check"); !isEmptyValue(reflect.ValueOf(httpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpHealthCheckProp)) {
		obj["httpHealthCheck"] = httpHealthCheckProp
	}
	httpsHealthCheckProp, err := expandComputeHealthCheckHttpsHealthCheck(d.Get("https_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("https_health_check"); !isEmptyValue(reflect.ValueOf(httpsHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpsHealthCheckProp)) {
		obj["httpsHealthCheck"] = httpsHealthCheckProp
	}
	tcpHealthCheckProp, err := expandComputeHealthCheckTcpHealthCheck(d.Get("tcp_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tcp_health_check"); !isEmptyValue(reflect.ValueOf(tcpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, tcpHealthCheckProp)) {
		obj["tcpHealthCheck"] = tcpHealthCheckProp
	}
	sslHealthCheckProp, err := expandComputeHealthCheckSslHealthCheck(d.Get("ssl_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ssl_health_check"); !isEmptyValue(reflect.ValueOf(sslHealthCheckProp)) && (ok || !reflect.DeepEqual(v, sslHealthCheckProp)) {
		obj["sslHealthCheck"] = sslHealthCheckProp
	}

	return resourceComputeHealthCheckEncoder(d, config, obj)
}

func resourceComputeHealthCheckEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := d.GetOk("http_health_check"); ok {
		obj["type"] = "HTTP"
		return obj, nil
	}
	if _, ok := d.GetOk("https_health_check"); ok {
		obj["type"] = "HTTPS"
		return obj, nil
	}
	if _, ok := d.GetOk("tcp_health_check"); ok {
		obj["type"] = "TCP"
		return obj, nil
	}
	if _, ok := d.GetOk("ssl_health_check"); ok {
		obj["type"] = "SSL"
		return obj, nil
	}

	return nil, fmt.Errorf("error in HealthCheck %s: No health check block specified.", d.Get("name").(string))
}

func expandComputeHealthCheckCheckIntervalSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHealthyThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckUnhealthyThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttpHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttpHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckHttpHealthCheckHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckRequestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttpsHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttpsHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttpsHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttpsHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttpsHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckHttpsHealthCheckHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckRequestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeHealthCheckTcpHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeHealthCheckTcpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckTcpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckTcpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckTcpHealthCheckRequest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeHealthCheckSslHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeHealthCheckSslHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckSslHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckSslHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckSslHealthCheckRequest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
