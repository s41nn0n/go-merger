package merger

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	intoJSONStructConst = "{\"UserID\":\"\",\"UserName\":\"test@example.com\",\"Password\":\"PASS\"}"
	fromJSONStructConst = "{\"UserID\":\"19b90b42-786d-44bf-be74-6f6d54edfe02\",\"UserName\":\"test@example.com\",\"Password\":\"\"}"
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
		fmt.Println("UserID == NILL !1!")
		t.FailNow()
	}

	if x["UserName"] != "test@example.com" {
		fmt.Println("UserName Was changed")
		t.FailNow()
	}

	if x["Password"] != "PASS" {
		fmt.Println("Password Was Changed")
		t.FailNow()
	}

	fmt.Println(x["UserID"])

}

func BenchmarkMergeMasterJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var new string
		var err error
		intoJSONStruct = intoJSONStructConst
		fromJSONStruct = fromJSONStructConst
		if new, err = MergeMasterJSON(intoJSONStruct, fromJSONStruct); err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		var x map[string]interface{}
		if err = json.Unmarshal([]byte(new), &x); err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		if x["UserID"] == nil && x["UserID"] == "" {
			fmt.Println("UserID NILL !1!")
			b.FailNow()
		}
	}
}
