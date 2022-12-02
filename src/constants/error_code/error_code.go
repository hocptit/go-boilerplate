package error_code

// xxyyzz: xx: MODULE; yy: FUNCTION; zz: error code in FUNCTION
const (
	/* COMMON ERROR*/

	INTERNAL_SERVER = "000001"
	INVALID_PARAMS  = "000002"
	UNKNOWN         = "999999"

	/* COMMON ERROR*/

)

var MsgFlags = map[string]string{
	/* COMMON ERROR*/
	INTERNAL_SERVER: "INTERNAL_SERVER",
	INVALID_PARAMS:  "INVALID_PARAMS",
	UNKNOWN:         "UNKNOWN",
	/* COMMON ERROR*/

	/* AUTHOR ERROR*/

	NOT_FOUND_AUTHOR: "NOT_FOUND_AUTHOR",

	/* AUTHOR ERROR*/

	/* BOOK ERROR*/

	NOT_FOUND_BOOK: "NOT_FOUND_BOOK",

	/* BOOK ERROR*/

}

// GetMsg get error information based on Code
func GetMsg(code string) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[UNKNOWN]
}
