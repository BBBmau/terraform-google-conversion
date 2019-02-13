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

func GetComputeSnapshotApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeSnapshotName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeSnapshotLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	sourceDiskProp, err := expandComputeSnapshotSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_disk"); !isEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	zoneProp, err := expandComputeSnapshotZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	snapshotEncryptionKeyProp, err := expandComputeSnapshotSnapshotEncryptionKey(d.Get("snapshot_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot_encryption_key"); !isEmptyValue(reflect.ValueOf(snapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, snapshotEncryptionKeyProp)) {
		obj["snapshotEncryptionKey"] = snapshotEncryptionKeyProp
	}
	sourceDiskEncryptionKeyProp, err := expandComputeSnapshotSourceDiskEncryptionKey(d.Get("source_disk_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_disk_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceDiskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceDiskEncryptionKeyProp)) {
		obj["sourceDiskEncryptionKey"] = sourceDiskEncryptionKeyProp
	}

	return obj, nil
}

func expandComputeSnapshotName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeSnapshotLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDisk(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotSnapshotEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	return transformed, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSourceDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	return transformed, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
