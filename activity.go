package createrdsinstance

import (
	"encoding/json"
	s "strings"

	"github.com/DipeshTest/allstarsshared/awsrds"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/service/rds"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-createrdsinstance")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	accessKeyId := s.TrimSpace(context.GetInput("accessKeyId").(string))
	secretAccessKey := s.TrimSpace(context.GetInput("secretAccessKey").(string))
	createInstanceInputJson := s.TrimSpace(context.GetInput("createInstanceInputJson").(string))
	region := s.TrimSpace(context.GetInput("region").(string))

	if len(accessKeyId) == 0 {

		context.SetOutput("statusCode", "101")

		context.SetOutput("createInstanceOutput", "Access KeyId field is blank")
	} else if len(secretAccessKey) == 0 {

		context.SetOutput("statusCode", "102")

		context.SetOutput("createInstanceOutput", "SecretAccessKey  field is blank")

	} else if len(region) == 0 {

		context.SetOutput("statusCode", "103")

		context.SetOutput("createInstanceOutput", "Region  field is blank")

	} else if len(createInstanceInputJson) == 0 {

		context.SetOutput("statusCode", "104")

		context.SetOutput("createInstanceOutput", "CreateInstanceInputJson  field is blank")

	} else {
		input := &rds.CreateDBInstanceInput{}
		err := json.Unmarshal([]byte(createInstanceInputJson), &input)
		if err == nil {
			code, message := awsrds.CreateRdsInstance(accessKeyId, secretAccessKey, region, input)
			context.SetOutput("statusCode", code)
			context.SetOutput("createInstanceOutput", message)
		} else {
			context.SetOutput("statusCode", "105")

			context.SetOutput("createInstanceOutput", "Error while parsing the input specified, "+err.Error())
		}
	}

	return true, err
}
