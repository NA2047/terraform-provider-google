// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package backupdr

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleCloudBackupDRDataSource() *schema.Resource {
	dsSchema := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
			Description: `Name of the datasource to create.
			It must have the format "projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}".
			'{datasource}' cannot be changed after creation. It must be between 3-63 characters long and must be unique within the backup vault.`,
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The DataSource resource instance state.`,
		},
		"labels": {
			Type:        schema.TypeMap,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: `Resource labels to represent user provided metadata.`,
		},
		"create_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The time when the instance was created.`,
		},
		"update_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The time when the instance was updated.`,
		},
		"backup_count": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `Number of backups in the data source.`,
		},
		"etag": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.`,
		},
		"total_stored_bytes": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The number of bytes (metadata and data) stored in this datasource.`,
		},
		"config_state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The backup configuration state.`,
		},
		"backup_config_info": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"last_backup_state": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `LastBackupstate tracks whether the last backup was not yet started, successful, failed, or could not be run because of the lack of permissions.`,
					},
					"last_successful_backup_consistency_time": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `If the last backup were successful, this field has the consistency date.`,
					},
					"last_backup_error": {
						Type:        schema.TypeMap,
						Computed:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: `If the last backup failed, this field has the error message.`,
					},
					"gcp_backup_config": {
						Type:     schema.TypeList,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"backup_plan": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the backup plan.`,
								},
								"backup_plan_description": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The description of the backup plan.`,
								},
								"backup_plan_association": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the backup plan association.`,
								},
								"backup_plan_rules": {
									Type:        schema.TypeList,
									Computed:    true,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Description: `The names of the backup plan rules which point to this backupvault`,
								},
							},
						},
						Description: `Configuration for a Google Cloud resource.`,
					},
					"backup_appliance_backup_config": {
						Type:     schema.TypeList,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"backup_appliance_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the backup appliance.`,
								},
								"backup_appliance_id": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The ID of the backup appliance.`,
								},
								"sla_id": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The ID of the SLA of this application.`,
								},
								"application_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the application.`,
								},
								"host_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the host where the application is running.`,
								},
								"slt_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the SLT associated with the application.`,
								},
								"slp_name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The name of the SLP associated with the application.`,
								},
							},
						},
						Description: `Configuration for an application backed up by a Backup Appliance.`,
					},
				},
			},
			Description: `Details of how the resource is configured for backup.`,
		},
		"data_source_gcp_resource": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"gcp_resourcename": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `Full resource pathname URL of the source Google Cloud resource.`,
					},
					"location": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `Location of the resource: <region>/<zone>/"global"/"unspecified".`,
					},
					"type": {
						Type:     schema.TypeString,
						Computed: true,
						Description: `The type of the Google Cloud resource. Use the Unified Resource Type,
						eg. compute.googleapis.com/Instance.`,
					},
					"compute_instance_data_source_properties": {
						Type:     schema.TypeList,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `Name of the compute instance backed up by the datasource.`,
								},
								"description": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The description of the Compute Engine instance.`,
								},
								"machine_type": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The machine type of the instance.`,
								},
								"total_disk_count": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The total number of disks attached to the Instance.`,
								},
								"total_disk_size_gb": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: `The sum of all the disk sizes.`,
								},
							},
						},
						Description: `ComputeInstanceDataSourceProperties has a subset of Compute Instance properties that are useful at the Datasource level.`,
					},
				},
			},
			Description: `The backed up resource is a Google Cloud resource.
			The word 'DataSource' was included in the names to indicate that this is
			the representation of the Google Cloud resource used within the
			DataSource object.`,
		},
		"data_source_backup_appliance_application": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"application_name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `The name of the Application as known to the Backup Appliance.`,
					},
					"backup_appliance": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `Appliance name.`,
					},
					"appliance_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `Appliance Id of the Backup Appliance.`,
					},
					"type": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `The type of the application. e.g. VMBackup`,
					},
					"application_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `The appid field of the application within the Backup Appliance.`,
					},
					"hostname": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `Hostname of the host where the application is running.`,
					},
					"host_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: `Hostid of the application host.`,
					},
				},
			},
			Description: `The backed up resource is a backup appliance application.`,
		},
		"location": {
			Type:     schema.TypeString,
			Required: true,
		},
		"project": {
			Type:     schema.TypeString,
			Required: true,
		},
		"data_source_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"backup_vault_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	return &schema.Resource{
		Read:   DataSourceGoogleCloudBackupDRDataSourceRead,
		Schema: dsSchema,
	}
}

func DataSourceGoogleCloudBackupDRDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	location, err := tpgresource.GetLocation(d, config)
	if err != nil {
		return err
	}
	if len(location) == 0 {
		return fmt.Errorf("Cannot determine location: set location in this data source or at provider-level")
	}

	billingProject := project
	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupVaults/{{backup_vault_id}}/dataSources/{{data_source_id}}")

	if err != nil {
		return err
	}

	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})

	if err != nil {
		return fmt.Errorf("Error reading BackupVault: %s", err)
	}

	if err := d.Set("name", flattenDataSourceBackupDRDataSourceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("create_time", flattenDataSourceBackupDRDataSourceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("update_time", flattenDataSourceBackupDRDataSourceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("backup_count", flattenDataSourceBackupDRDataSourceBackupCount(res["backupCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("etag", flattenDataSourceBackupDRDataSourceEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("state", flattenDataSourceBackupDRDataSourceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("total_stored_bytes", flattenDataSourceBackupDRDataSourceTotalStoredBytes(res["totalStoredBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("backup_config_info", flattenDataSourceBackupDRDataSourceBackupConfigInfo(res["backupConfigInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("data_source_gcp_resource", flattenBackupDRDataSourceDataSourceGCPResource(res["dataSourceGcpResource"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	if err := d.Set("data_source_backup_appliance_application", flattenBackupDRDataSourceDataSourceBackupApplianceApplication(res["dataSourceBackupApplianceApplication"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataSource: %s", err)
	}

	d.SetId(res["name"].(string))

	return nil
}

func flattenDataSourceBackupDRDataSourceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceBackupCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceTotalStoredBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceBackupConfigInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["last_backup_state"] = flattenDataSourceBackupDRDataSourceConfigInfoLastBackupState(original["lastBackupState"], d, config)
	transformed["last_successful_backup_consistency_time"] = flattenDataSourceBackupDRDataSourceConfigInfoLastSuccessfulBackupConsistencyTime(original["lastSuccessfulBackupConsistencyTime"], d, config)
	transformed["last_backup_error"] = flattenDataSourceBackupDRDataSourceConfigInfoLastBackupError(original["lastBackupError"], d, config)
	transformed["gcp_backup_config"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfig(original["gcpBackupConfig"], d, config)
	transformed["backup_appliance_backup_config"] = flattenBackupDRDataSourceConfigInfoBackupApplianceBackupConfig(original["backupApplianceBackupConfig"], d, config)

	return []interface{}{transformed}
}

func flattenDataSourceBackupDRDataSourceConfigInfoLastBackupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceConfigInfoLastSuccessfulBackupConsistencyTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataSourceBackupDRDataSourceConfigInfoLastBackupError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoBackupApplianceBackupConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["backup_appliance_name"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigBackupApplianceName(original["backupApplianceName"], d, config)
	transformed["backup_appliance_id"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigBackupApplianceId(original["backupApplianceId"], d, config)
	transformed["sla_id"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigSlaId(original["slaId"], d, config)
	transformed["application_name"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigApplicationName(original["applicationName"], d, config)
	transformed["slt_name"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigSltName(original["sltName"], d, config)
	transformed["slp_name"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigSlpName(original["slpName"], d, config)

	return []interface{}{transformed}
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["backup_plan"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlan(original["backupPlan"], d, config)
	transformed["backup_plan_description"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlanDescription(original["backupPlanDescription"], d, config)
	transformed["backup_plan_association"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlanAssociation(original["backupPlanAssociation"], d, config)
	transformed["backup_plan_rules"] = flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlanRules(original["backupPlanRules"], d, config)

	return []interface{}{transformed}
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlan(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlanDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlanAssociation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupPlanRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigBackupApplianceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigBackupApplianceId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigSlaId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigApplicationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigSltName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceConfigInfoGCPBackupConfigBackupApplianceBackupConfigSlpName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gcp_resourcename"] = flattenBackupDRDataSourceDataSourceGCPResourceGCPResourceName(original["gcpResourcename"], d, config)
	transformed["location"] = flattenBackupDRDataSourceDataSourceGCPResourceLocation(original["location"], d, config)
	transformed["compute_instance_data_source_properties"] = flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourceProperties(original["computeInstanceDatasourceProperties"], d, config)

	return []interface{}{transformed}
}

func flattenBackupDRDataSourceDataSourceGCPResourceGCPResourceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResourceLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourceProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] = flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesName(original["name"], d, config)
	transformed["description"] = flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesDescription(original["description"], d, config)
	transformed["machine_type"] = flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesMachineType(original["machineType"], d, config)
	transformed["total_disk_count"] = flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesTotalDiskCount(original["totalDiskCount"], d, config)
	transformed["total_disk_size_gb"] = flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesTotalDiskSizeGb(original["totalDiskSizeGb"], d, config)

	return []interface{}{transformed}
}

func flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesMachineType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesTotalDiskCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceGCPResourceComputeInstanceDataSourcePropertiesTotalDiskSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplication(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["application_name"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationName(original["applicationName"], d, config)
	transformed["backup_appliance"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationBackupAppliance(original["backupAppliance"], d, config)
	transformed["appliance_id"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplianceId(original["applianceId"], d, config)
	transformed["type"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationType(original["type"], d, config)
	transformed["application_id"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationId(original["applicationId"], d, config)
	transformed["hostname"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationHostname(original["hostname"], d, config)
	transformed["host_id"] = flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationHostId(original["hostId"], d, config)

	return []interface{}{transformed}
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationBackupAppliance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplianceId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationHostname(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRDataSourceDataSourceBackupApplianceApplicationApplicationHostId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
