package errorcode

// xxyyzz: xx: MODULE; yy: FUNCTION; zz: error code in FUNCTION
const (
	/* COMMON ERROR*/

	InternalServer = "000001"
	InvalidParams  = "000002"
	UNKNOWN        = "999999"

	/* COMMON ERROR*/

)

var MsgFlags = map[string]string{
	/* COMMON ERROR*/
	InternalServer: "INTERNAL_SERVER",
	InvalidParams:  "INVALID_PARAMS",
	UNKNOWN:        "UNKNOWN",
	/* COMMON ERROR*/

	/* AUTHOR ERROR*/

	NotFoundAuthor: "NOT_FOUND_AUTHOR",

	/* AUTHOR ERROR*/

	/* BOOK ERROR*/

	NotFoundBook: "NOT_FOUND_BOOK",

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
