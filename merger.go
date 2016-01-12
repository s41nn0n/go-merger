package merger

import (
	"encoding/json"
	"fmt"
)

/*
MergeError error,
potential helper for error handling
*/
type MergeError struct {
	What int
}

/*
Error required for MergeError struct
*/
func (e MergeError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

/*
MergeMaster merges two maps
  This Takes the into, and merges the data from into it.
  This function assumes into has the primary data, from will not overwrite the data from into

  This will itterate through the values in into, and only if empty, will try fetch the data from
*/
func MergeMaster(into *map[string]interface{}, from map[string]interface{}) {
	/*
	   This can be changed (possibly).
	     To add concurrecncy.
	     To speed up this process
	*/
	for key, value := range *into {
		if (value == nil || value == "") && (from[key] != nil || from[key] != "") {
			(*into)[key] = from[key]
		}
	}
}

/*
MergeMasterJSON takes two json strings,
   converts to map[string]interface{}
   and returns json string of the interfaces merged
*/
func MergeMasterJSON(into, from string) (string, error) {
	var intoStruct, fromStruct map[string]interface{}
	var err error

	err = json.Unmarshal([]byte(into), &intoStruct)
	if err != nil {
		fmt.Println(err)
		return "", MergeError{1}
	}
	err = json.Unmarshal([]byte(from), &fromStruct)

	MergeMaster(&intoStruct, fromStruct)

	ret, err := json.Marshal(&intoStruct)
	return string(ret[:]), err
}

/*
MergeMasterInterface client side function to pass
any two interfaces{} and merges them.
Returns type interface with error
*/
func MergeMasterInterface(into, from interface{}) (interface{}, error) {
	var intoString, fromString []byte
	var err error

	intoString, err = json.Marshal(into)
	if err != nil {
		return nil, MergeError{2}
	}

	fromString, err = json.Marshal(from)
	if err != nil {
		return nil, MergeError{3}
	}

	var retString string

	retString, err = MergeMasterJSON(string(intoString[:]), string(fromString[:]))
	if err != nil {
		fmt.Println(retString)
		return nil, MergeError{4}
	}

	err = json.Unmarshal([]byte(retString), into)
	return into, err
}
