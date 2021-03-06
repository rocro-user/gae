// Code generated by "stringer -type=PropertyType"; DO NOT EDIT.

package datastore

import "fmt"

const _PropertyType_name = "PTNullPTIntPTTimePTBoolPTBytesPTStringPTFloatPTGeoPointPTKeyPTBlobKeyPTUnknown"

var _PropertyType_index = [...]uint8{0, 6, 11, 17, 23, 30, 38, 45, 55, 60, 69, 78}

func (i PropertyType) String() string {
	if i >= PropertyType(len(_PropertyType_index)-1) {
		return fmt.Sprintf("PropertyType(%d)", i)
	}
	return _PropertyType_name[_PropertyType_index[i]:_PropertyType_index[i+1]]
}
