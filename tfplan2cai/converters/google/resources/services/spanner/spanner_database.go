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

package spanner

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// customizeDiff func for additional checks on google_spanner_database properties:
func resourceSpannerDBDdlCustomDiffFunc(diff tpgresource.TerraformResourceDiff) error {
	old, new := diff.GetChange("ddl")
	oldDdls := old.([]interface{})
	newDdls := new.([]interface{})
	var err error

	if len(newDdls) < len(oldDdls) {
		err = diff.ForceNew("ddl")
		if err != nil {
			return fmt.Errorf("ForceNew failed for ddl, old - %v and new - %v", oldDdls, newDdls)
		}
		return nil
	}

	for i := range oldDdls {
		if newDdls[i].(string) != oldDdls[i].(string) {
			err = diff.ForceNew("ddl")
			if err != nil {
				return fmt.Errorf("ForceNew failed for ddl, old - %v and new - %v", oldDdls, newDdls)
			}
			return nil
		}
	}
	return nil
}

func resourceSpannerDBDdlCustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	// separate func to allow unit testing
	return resourceSpannerDBDdlCustomDiffFunc(diff)
}

func ValidateDatabaseRetentionPeriod(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	valueError := fmt.Errorf("version_retention_period should be in range [1h, 7d], in a format resembling 1d, 24h, 1440m, or 86400s")

	r := regexp.MustCompile("^(\\d{1}d|\\d{1,3}h|\\d{2,5}m|\\d{4,6}s)$")
	if !r.MatchString(value) {
		errors = append(errors, valueError)
		return
	}

	unit := value[len(value)-1:]
	multiple := value[:len(value)-1]
	num, err := strconv.Atoi(multiple)
	if err != nil {
		errors = append(errors, valueError)
		return
	}

	if unit == "d" && (num < 1 || num > 7) {
		errors = append(errors, valueError)
		return
	}
	if unit == "h" && (num < 1 || num > 7*24) {
		errors = append(errors, valueError)
		return
	}
	if unit == "m" && (num < 1*60 || num > 7*24*60) {
		errors = append(errors, valueError)
		return
	}
	if unit == "s" && (num < 1*60*60 || num > 7*24*60*60) {
		errors = append(errors, valueError)
		return
	}

	return
}

func resourceSpannerDBVirtualUpdate(d *schema.ResourceData, resourceSchema map[string]*schema.Schema) bool {
	// deletion_protection is the only virtual field
	if d.HasChange("deletion_protection") {
		for field := range resourceSchema {
			if field == "deletion_protection" {
				continue
			}
			if d.HasChange(field) {
				return false
			}
		}
		return true
	}
	return false
}

const SpannerDatabaseAssetType string = "spanner.googleapis.com/Database"

func ResourceConverterSpannerDatabase() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SpannerDatabaseAssetType,
		Convert:   GetSpannerDatabaseCaiObject,
	}
}

func GetSpannerDatabaseCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//spanner.googleapis.com/projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSpannerDatabaseApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SpannerDatabaseAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/spanner/v1/rest",
				DiscoveryName:        "Database",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSpannerDatabaseApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSpannerDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	versionRetentionPeriodProp, err := expandSpannerDatabaseVersionRetentionPeriod(d.Get("version_retention_period"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_retention_period"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionRetentionPeriodProp)) && (ok || !reflect.DeepEqual(v, versionRetentionPeriodProp)) {
		obj["versionRetentionPeriod"] = versionRetentionPeriodProp
	}
	extraStatementsProp, err := expandSpannerDatabaseDdl(d.Get("ddl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ddl"); !tpgresource.IsEmptyValue(reflect.ValueOf(extraStatementsProp)) && (ok || !reflect.DeepEqual(v, extraStatementsProp)) {
		obj["extraStatements"] = extraStatementsProp
	}
	encryptionConfigProp, err := expandSpannerDatabaseEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	databaseDialectProp, err := expandSpannerDatabaseDatabaseDialect(d.Get("database_dialect"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database_dialect"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseDialectProp)) && (ok || !reflect.DeepEqual(v, databaseDialectProp)) {
		obj["databaseDialect"] = databaseDialectProp
	}
	enableDropProtectionProp, err := expandSpannerDatabaseEnableDropProtection(d.Get("enable_drop_protection"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_drop_protection"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableDropProtectionProp)) && (ok || !reflect.DeepEqual(v, enableDropProtectionProp)) {
		obj["enableDropProtection"] = enableDropProtectionProp
	}
	instanceProp, err := expandSpannerDatabaseInstance(d.Get("instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	return resourceSpannerDatabaseEncoder(d, config, obj)
}

func resourceSpannerDatabaseEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["createStatement"] = fmt.Sprintf("CREATE DATABASE `%s`", obj["name"])
	if dialect, ok := obj["databaseDialect"]; ok && dialect == "POSTGRESQL" {
		obj["createStatement"] = fmt.Sprintf("CREATE DATABASE \"%s\"", obj["name"])
	}

	// Extra DDL statements are removed from the create request and instead applied to the database in
	// a post-create action, to accommodate retrictions when creating PostgreSQL-enabled databases.
	// https://cloud.google.com/spanner/docs/create-manage-databases#create_a_database
	log.Printf("[DEBUG] Preparing to create new Database. Any extra DDL statements will be applied to the Database in a separate API call")

	delete(obj, "name")
	delete(obj, "instance")

	return obj, nil
}

func expandSpannerDatabaseName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseVersionRetentionPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseDdl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandSpannerDatabaseEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandSpannerDatabaseEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseDatabaseDialect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseEnableDropProtection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("instances", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for instance: %s", err)
	}
	return f.RelativeLink(), nil
}
