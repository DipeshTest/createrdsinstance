package createrdsinstance

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

// func TestRDSCreate_InvalidRegion(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us22")
// 	tc.SetInput("createInstanceInputJson", `{ "AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "106")
//
// }

// func TestRDSCreate_EmptyAccessKeyId(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("createInstanceInputJson", `{ "AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "101")
//
// }

// func TestRDSCreate_EmptySecretAccessKey(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us22")
// 	tc.SetInput("createInstanceInputJson", `{ "AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "102")
//
// }

// func TestRDSCreate_EmptyCreateInstanceInputJson(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us22")
// 	tc.SetInput("createInstanceInputJson", "")
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "104")
//
// }
// func TestRDSCreate_EmptyRegion(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "")
// 	tc.SetInput("createInstanceInputJson", `{ "AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "103")
//
// }

// func TestRDSCreate_InvalidAccessKeyId(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "123")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("createInstanceInputJson", `{ "AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "403")
//
// }

// func TestRDSCreate_InvalidSecretAccessKey(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "1234")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("createInstanceInputJson", `{ "AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "403")
//
// }

// func TestRDSCreate_InvalidCreateInstanceInputJson(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("createInstanceInputJson", `"AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "105")
//
// }

// func TestRDSCreate_ExistingDBInstanceIdentifier(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("createInstanceInputJson", ` {"AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "abt", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "400")
//
// }

// func TestRDSCreate_CreateRdsDBInstance(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("createInstanceInputJson", ` {"AllocatedStorage": 5, "DBInstanceClass": "db.t2.micro", "DBInstanceIdentifier": "ab7", "Engine": "MySQL", "MasterUserPassword": "MyPassword", "MasterUsername": "MyUser" }`)
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("createInstanceOutput")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }
