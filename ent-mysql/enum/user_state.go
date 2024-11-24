package enum

type UserState int

// Define enum values as constants.
const (
	On  UserState = iota // 0
	Off                  // 1
)

// String representations for each role (optional, for debugging/logging).
var userStateNames = []string{"On", "Off"}

// String method to get the name of the role.
func (r UserState) String() string {
	if int(r) < 0 || int(r) >= len(userStateNames) {
		return "Unknown"
	}
	return userStateNames[r]
}
