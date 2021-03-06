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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceHealthcareFhirStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareFhirStoreCreate,
		Read:   resourceHealthcareFhirStoreRead,
		Update: resourceHealthcareFhirStoreUpdate,
		Delete: resourceHealthcareFhirStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareFhirStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dataset": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `Identifies the dataset addressed by this request. Must be in the format
'projects/{project}/locations/{location}/datasets/{dataset}'`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name for the FhirStore.

** Changing this property may recreate the FHIR store (removing all data) **`,
			},
			"disable_referential_integrity": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
creation. The default value is false, meaning that the API will enforce referential integrity and fail the
requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
will skip referential integrity check. Consequently, operations that rely on references, such as
Patient.get$everything, will not return all the results if broken references exist.

** Changing this property may recreate the FHIR store (removing all data) **`,
			},
			"disable_resource_versioning": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
attempts to read the historical versions.

** Changing this property may recreate the FHIR store (removing all data) **`,
			},
			"enable_history_import": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Whether to allow the bulk import API to accept history bundles and directly insert historical resource
versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
will fail with an error.

** Changing this property may recreate the FHIR store (removing all data) **

** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **`,
			},
			"enable_update_create": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
notifications.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-supplied key-value pairs used to organize FHIR stores.

Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}

Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}

No more than 64 labels can be associated with a given store.

An object containing a list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"notification_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `A nested object resource`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_topic": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.`,
						},
					},
				},
			},
			"version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DSTU2", "STU3", "R4", ""}, false),
				Description:  `The FHIR specification version. Default value: "STU3" Possible values: ["DSTU2", "STU3", "R4"]`,
				Default:      "STU3",
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully qualified name of this dataset`,
			},
		},
	}
}

func resourceHealthcareFhirStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareFhirStoreName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	versionProp, err := expandHealthcareFhirStoreVersion(d.Get("version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version"); !isEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	enableUpdateCreateProp, err := expandHealthcareFhirStoreEnableUpdateCreate(d.Get("enable_update_create"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_update_create"); !isEmptyValue(reflect.ValueOf(enableUpdateCreateProp)) && (ok || !reflect.DeepEqual(v, enableUpdateCreateProp)) {
		obj["enableUpdateCreate"] = enableUpdateCreateProp
	}
	disableReferentialIntegrityProp, err := expandHealthcareFhirStoreDisableReferentialIntegrity(d.Get("disable_referential_integrity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_referential_integrity"); !isEmptyValue(reflect.ValueOf(disableReferentialIntegrityProp)) && (ok || !reflect.DeepEqual(v, disableReferentialIntegrityProp)) {
		obj["disableReferentialIntegrity"] = disableReferentialIntegrityProp
	}
	disableResourceVersioningProp, err := expandHealthcareFhirStoreDisableResourceVersioning(d.Get("disable_resource_versioning"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_resource_versioning"); !isEmptyValue(reflect.ValueOf(disableResourceVersioningProp)) && (ok || !reflect.DeepEqual(v, disableResourceVersioningProp)) {
		obj["disableResourceVersioning"] = disableResourceVersioningProp
	}
	enableHistoryImportProp, err := expandHealthcareFhirStoreEnableHistoryImport(d.Get("enable_history_import"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_history_import"); !isEmptyValue(reflect.ValueOf(enableHistoryImportProp)) && (ok || !reflect.DeepEqual(v, enableHistoryImportProp)) {
		obj["enableHistoryImport"] = enableHistoryImportProp
	}
	labelsProp, err := expandHealthcareFhirStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareFhirStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(notificationConfigProp)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores?fhirStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FhirStore: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FhirStore: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FhirStore %q: %#v", d.Id(), res)

	return resourceHealthcareFhirStoreRead(d, meta)
}

func resourceHealthcareFhirStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("HealthcareFhirStore %q", d.Id()))
	}

	res, err = resourceHealthcareFhirStoreDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing HealthcareFhirStore because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenHealthcareFhirStoreName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("version", flattenHealthcareFhirStoreVersion(res["version"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("enable_update_create", flattenHealthcareFhirStoreEnableUpdateCreate(res["enableUpdateCreate"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("disable_referential_integrity", flattenHealthcareFhirStoreDisableReferentialIntegrity(res["disableReferentialIntegrity"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("disable_resource_versioning", flattenHealthcareFhirStoreDisableResourceVersioning(res["disableResourceVersioning"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("enable_history_import", flattenHealthcareFhirStoreEnableHistoryImport(res["enableHistoryImport"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("labels", flattenHealthcareFhirStoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("notification_config", flattenHealthcareFhirStoreNotificationConfig(res["notificationConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}

	return nil
}

func resourceHealthcareFhirStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	enableUpdateCreateProp, err := expandHealthcareFhirStoreEnableUpdateCreate(d.Get("enable_update_create"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_update_create"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableUpdateCreateProp)) {
		obj["enableUpdateCreate"] = enableUpdateCreateProp
	}
	labelsProp, err := expandHealthcareFhirStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareFhirStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FhirStore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("enable_update_create") {
		updateMask = append(updateMask, "enableUpdateCreate")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("notification_config") {
		updateMask = append(updateMask, "notificationConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating FhirStore %q: %s", d.Id(), err)
	}

	return resourceHealthcareFhirStoreRead(d, meta)
}

func resourceHealthcareFhirStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting FhirStore %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "FhirStore")
	}

	log.Printf("[DEBUG] Finished deleting FhirStore %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareFhirStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	fhirStoreId, err := parseHealthcareFhirStoreId(d.Id(), config)
	if err != nil {
		return nil, err
	}

	d.Set("dataset", fhirStoreId.DatasetId.datasetId())
	d.Set("name", fhirStoreId.Name)

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareFhirStoreName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreEnableUpdateCreate(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreDisableReferentialIntegrity(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreDisableResourceVersioning(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreEnableHistoryImport(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareFhirStoreNotificationConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_topic"] =
		flattenHealthcareFhirStoreNotificationConfigPubsubTopic(original["pubsubTopic"], d, config)
	return []interface{}{transformed}
}
func flattenHealthcareFhirStoreNotificationConfigPubsubTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandHealthcareFhirStoreName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreEnableUpdateCreate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreDisableReferentialIntegrity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreDisableResourceVersioning(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreEnableHistoryImport(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandHealthcareFhirStoreNotificationConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandHealthcareFhirStoreNotificationConfigPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	return transformed, nil
}

func expandHealthcareFhirStoreNotificationConfigPubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceHealthcareFhirStoreDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	d.Set("self_link", res["name"].(string))
	res["name"] = d.Get("name").(string)
	return res, nil
}
