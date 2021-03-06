// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-01 10:59:26 700B77                                   [zr/uuid.go]
// -----------------------------------------------------------------------------

package zr

// # Functions
//   IsUUID(val interface{}) bool
//   UUID() string

// import "crypto/rand" // used via mod.rand proxy
import "fmt" // standard

// -----------------------------------------------------------------------------
// # Functions

// IsUUID returns true if the given string is a well-formed UUID strng.
// It accepts both the standard UUID with '-' and the shortened 32-character
// UUID. If the type of 'val' is not a string, string pointer or Stringer
// interface, returns false.
func IsUUID(val interface{}) bool {
	switch val := val.(type) {
	case string:
		for str.Contains(val, "-") {
			val = str.Replace(val, "-", "", -1)
		}
		if len(val) != 32 {
			return false
		}
		for i, ch := range val {
			if (i == 12 && ch != '4') ||
				(i == 16 &&
					(ch != '8' && ch != '9' &&
						ch != 'A' && ch != 'B' &&
						ch != 'a' && ch != 'b')) {
				return false
			}
			if !((ch >= '0' && ch <= '9') ||
				(ch >= 'a' && ch <= 'f') ||
				(ch >= 'A' && ch <= 'F')) {
				return false
			}
		}
		return true
	case fmt.Stringer:
		return IsUUID(val.String())
	case *string:
		return IsUUID(*val)
	}
	return false
} //                                                                      IsUUID

// UUID generates and returns a new UUID (Universally Unique Identifier).
// The format is 'XXXXXXXX-XXXX-4XXX-ZXXX-XXXXXXXXXXXX' where every X is a
// random upper-case hex digit, and Z must be one of '8', '9', 'A' or 'B'.
func UUID() string {
	var b = make([]byte, 16)
	var _, err = mod.rand.Read(b)
	if err != nil {
		return ""
	}
	b[6] = (b[6] | 0x40) & 0x4F // 13th character (at [12]) must be '4'
	b[8] = (b[8] | 0x80) & 0xBF // 17th c. at [16] must be '8', '9', 'A', or 'B'
	return fmt.Sprintf("%X-%X-%X-%X-%X",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
} //                                                                        UUID

//end
