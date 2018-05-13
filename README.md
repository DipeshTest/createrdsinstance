---
title: Create AWS RDS Instance
weight: 1
---

# Create AWS RDS Instance
This activity allows you to create AWS RDS DB Instance of your choice.This activity is built by Team AllStars.

## Installation
### Flogo CLI
```bash
flogo install github.com/DipeshTest/createrdsinstance
```

### Third-party libraries used
- #### aws-sdk-go:
	Package aws provides the core SDK's utilities and shared types. This package's utilities are used to simplify setting and reading API operations parameters. For more details, check - https://docs.aws.amazon.com/sdk-for-go/api/aws/
	
	
## Schema
Inputs and Outputs:

```json
 "inputs": [{
		"name": "accessKeyId",
		"type": "string",
		"required": true
	},
	{
		"name": "secretAccessKey",
		"type": "string",
		"required": true
	},
	{
		"name": "region",
		"type": "string",
		"required": true
	},
	{
		"name": "createInstanceInputJson",
		"type": "any",
		"required": true
	}],
 "outputs": [{
		"name": "statusCode",
		"type": "string"
	},
	{
		"name": "createInstanceOutput",
		"type": "any"
	}]
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accessKeyId | True     | Access Key ID for your AWS acount, Use link :  https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html |         
| secretAccessKey   | True    |Secret Access Key for your AWS acount, Use link :  https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html |
| region    | True     | Regions and Availability Zones for AWS, Check link : https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html|  
| createInstanceInputJson   | True     | Input json to create rds db instance, Please refer below json schema |

Note - 
1. While creating the IAM user for AWS account, Please add the IAM user to the group which has privilige to create RDS DB instance 
2. Please refer link : https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBCluster.html for more input and output details

```json
{
	"AllocatedStorage": null,
	"AutoMinorVersionUpgrade": null,
	"AvailabilityZone": null,
	"BackupRetentionPeriod": null,
	"CharacterSetName": null,
	"CopyTagsToSnapshot": null,
	"DBClusterIdentifier": null,
	"DBInstanceClass": null,
	"DBInstanceIdentifier": null,
	"DBName": null,
	"DBParameterGroupName": null,
	"DBSecurityGroups": null,
	"DBSubnetGroupName": null,
	"Domain": null,
	"DomainIAMRoleName": null,
	"EnableCloudwatchLogsExports": null,
	"EnableIAMDatabaseAuthentication": null,
	"EnablePerformanceInsights": null,
	"Engine": null,
	"EngineVersion": null,
	"Iops": null,
	"KmsKeyId": null,
	"LicenseModel": null,
	"MasterUserPassword": null,
	"MasterUsername": null,
	"MonitoringInterval": null,
	"MonitoringRoleArn": null,
	"MultiAZ": null,
	"OptionGroupName": null,
	"PerformanceInsightsKMSKeyId": null,
	"Port": null,
	"PreferredBackupWindow": null,
	"PreferredMaintenanceWindow": null,
	"PromotionTier": null,
	"PubliclyAccessible": null,
	"StorageEncrypted": null,
	"StorageType": null,
	"Tags": null,
	"TdeCredentialArn": null,
	"TdeCredentialPassword": null,
	"Timezone": null,
	"VpcSecurityGroupIds": null
}
```

Min. Required Fields : 
"AllocatedStorage","DBInstanceClass","DBInstanceIdentifier","Engine","MasterUserPassword","MasterUsername"

## Examples
Please refer activity_test.go 

## Response Codes

| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful and the response body contains the json output for operation requested.|
|101 |ACCESS KEYID BLANK ERROR| Access KeyId field is blank.|
|102 |SECRET ACCESSKEY BLANK ERROR| SecretAccessKey  field is blank.|
|103 |REGION BLANK ERROR| Region  field is blank.|
|104 |CreateInstanceInputJson BLANK ERROR| CreateInstanceInputJson  field is blank.|
|105 |CreateInstanceInputJson PARSE ERROR| CreateInstanceInputJson is an invalid json.|
|106 |INVALID REGION ERROR| When Region is invalid you will get this error.|

Note - Please refer link - https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Operations.html for additional response codes
