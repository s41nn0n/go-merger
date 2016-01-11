package merger

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	intoJSONStructConst = "{\"UserID\":\"\",\"UserName\":\"959dfbb6-f143-468c-85e0-997b753f79a5@iprg.co.za\",\"Password\":\"\""
	fromJSONStructConst = "{\"UserID\":\":19b90b42-786d-44bf-be74-6f6d54edfe02\",\"UserName\":\"959dfbb6-f143-468c-85e0-997b753f79a5@iprg.co.za\",\"Password\":\"\""
)

var intoJSONStruct, fromJSONStruct string

/*
TestMergeMasterJSON ...
*/
func TestMergeMasterJSON(t *testing.T) {
	var new string
	var err error
	intoJSONStruct = intoJSONStructConst
	fromJSONStruct = fromJSONStructConst
	if new, err = MergeMasterJSON(intoJSONStruct, fromJSONStruct); err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	var x map[string]interface{}
	if err = json.Unmarshal([]byte(new), &x); err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	if x["UserID"] == nil && x["UserID"] == "" {
		fmt.Println("UserID NILL !1!")
		t.FailNow()
	}

}
