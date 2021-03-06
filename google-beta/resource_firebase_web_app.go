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
)

func resourceFirebaseWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseWebAppCreate,
		Read:   resourceFirebaseWebAppRead,
		Update: resourceFirebaseWebAppUpdate,
		Delete: resourceFirebaseWebAppDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseWebAppImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user-assigned display name of the App.`,
			},
			"app_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Immutable. The globally unique, Firebase-assigned identifier of the App.

This identifier should be treated as an opaque token, as the data format is not specified.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully qualified resource name of the App, for example:

projects/projectId/webApps/appId`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFirebaseWebAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/webApps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WebApp: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating WebApp: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = firebaseOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating WebApp",
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create WebApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseWebAppName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating WebApp %q: %#v", d.Id(), res)

	return resourceFirebaseWebAppRead(d, meta)
}

func resourceFirebaseWebAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseWebApp %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseWebAppName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseWebAppDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseWebAppAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}

	return nil
}

func resourceFirebaseWebAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WebApp %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating WebApp %q: %s", d.Id(), err)
	}

	return resourceFirebaseWebAppRead(d, meta)
}

func resourceFirebaseWebAppDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Firebase WebApp resources"+
		" cannot be deleted from GCP. The resource %s will be removed from Terraform"+
		" state, but will still be present on the server.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirebaseWebAppImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseWebAppName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseWebAppDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseWebAppAppId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirebaseWebAppDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
