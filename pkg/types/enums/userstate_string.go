// Code generated by "enumer -type=UserState -trimprefix=User -sql -output=userstate_string.go"; DO NOT EDIT.

package enums

import (
	"database/sql/driver"
	"fmt"
	"github.com/transerver/utils"
)

const _UserStateName = "UnverifiedNormal"

var _UserStateIndex = [...]uint8{0, 10, 16}

func (i UserState) String() string {
	if i >= UserState(len(_UserStateIndex)-1) {
		return fmt.Sprintf("UserState(%d)", i)
	}
	return _UserStateName[_UserStateIndex[i]:_UserStateIndex[i+1]]
}

var _UserStateValues = []UserState{0, 1}

var _UserStateNameToValueMap = map[string]UserState{
	_UserStateName[0:10]:  0,
	_UserStateName[10:16]: 1,
}

// UserStateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UserStateString(s string) (UserState, error) {
	if val, ok := _UserStateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to UserState values", s)
}

// UserStateValues returns all values of the enum
func UserStateValues() []UserState {
	return _UserStateValues
}

// IsAUserState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i UserState) IsAUserState() bool {
	for _, v := range _UserStateValues {
		if i == v {
			return true
		}
	}
	return false
}

func (i UserState) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *UserState) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case int:
		*i = UserState(v)
	case int8:
		*i = UserState(v)
	case int16:
		*i = UserState(v)
	case int32:
		*i = UserState(v)
	case int64:
		*i = UserState(v)
	case uint:
		*i = UserState(v)
	case uint8:
		*i = UserState(v)
	case uint16:
		*i = UserState(v)
	case uint32:
		*i = UserState(v)
	case uint64:
		*i = UserState(v)
	case bool:
		if v {
			*i = UserState(1)
		} else {
			*i = UserState(0)
		}
	case string:
		return s2s(i, v)
	case []byte:
		return s2s(i, utils.String(v))
	case fmt.Stringer:
		return s2s(i, v.String())
	}
	return nil
}

func s2s(i *UserState, src string) error {
	s, err := UserStateString(src)
	if err != nil {
		return err
	}
	*i = s
	return nil
}