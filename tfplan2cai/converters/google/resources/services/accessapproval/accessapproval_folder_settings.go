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

package accessapproval

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var accessApprovalCloudProductMapping = map[string]string{
	"appengine.googleapis.com": "App Engine",
	"bigquery.googleapis.com":  "BigQuery",
	"bigtable.googleapis.com":  "Cloud Bigtable",
	"cloudkms.googleapis.com":  "Cloud Key Management Service",
	"compute.googleapis.com":   "Compute Engine",
	"dataflow.googleapis.com":  "Cloud Dataflow",
	"iam.googleapis.com":       "Cloud Identity and Access Management",
	"pubsub.googleapis.com":    "Cloud Pub/Sub",
	"storage.googleapis.com":   "Cloud Storage",
}

func accessApprovalEnrolledServicesHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	cp := m["cloud_product"].(string)
	if n, ok := accessApprovalCloudProductMapping[cp]; ok {
		cp = n
	}
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(cp))) // ToLower just in case
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["enrollment_level"].(string))))
	return tpgresource.Hashcode(buf.String())
}

const AccessApprovalFolderSettingsAssetType string = "accessapproval.googleapis.com/FolderSettings"

func ResourceConverterAccessApprovalFolderSettings() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AccessApprovalFolderSettingsAssetType,
		Convert:   GetAccessApprovalFolderSettingsCaiObject,
	}
}

func GetAccessApprovalFolderSettingsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//accessapproval.googleapis.com/folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAccessApprovalFolderSettingsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AccessApprovalFolderSettingsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accessapproval/v1/rest",
				DiscoveryName:        "FolderSettings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAccessApprovalFolderSettingsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalFolderSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_emails"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalFolderSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enrolled_services"); !tpgresource.IsEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalFolderSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("active_key_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(activeKeyVersionProp)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}

	return obj, nil
}

func expandAccessApprovalFolderSettingsNotificationEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalFolderSettingsEnrolledServices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalFolderSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalFolderSettingsEnrolledServicesCloudProduct(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalFolderSettingsActiveKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
