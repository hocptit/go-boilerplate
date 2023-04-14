package exception

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"runtime"
	errorcode "server-go/src/constants/error_code"
	"server-go/src/share/constant"
	getLogger "server-go/src/share/logger"
	"server-go/src/share/response"
	"server-go/src/share/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contain dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		_, err := fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if err != nil {
			return nil
		}
		if file != lastFile {
			data, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		_, err = fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
		if err != nil {
			return nil
		}
	}
	return buf.Bytes()
}
func timeFormat(t time.Time) string {
	return t.Format("2006/01/02 - 15:04:05")
}
func RecoveryError(appIsReturnDetailErrors string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			logger := getLogger.GetLogger().Logging
			if err := recover(); err != nil {
				// nolint
				switch err.(type) {
				// check validator class
				case response.Response:
					// nolint
					errData, _ := err.(response.Response)
					if appIsReturnDetailErrors == "true" {
						ReturnBadRequestError(c, errData.ErrorCode, errData.Errors)
						return
					}
					ReturnBadRequestError(c, errData.ErrorCode, "")
					return
				default:
					stack := stack(3)
					logger.Errorf("[Recovery] %s %s panic recovered:\n%s\n%s%s",
						utils.TID(c), timeFormat(time.Now()), err, stack, "\033[0m")

					if appIsReturnDetailErrors == "true" {
						ReturnInternalServerError(c, errorcode.InternalServer, err)
						return
					}
					ReturnInternalServerError(c, errorcode.InternalServer, "")
					return
				}
			}
		}()
		c.Next()
	}
}

func NoRouteError(c *gin.Context) {
	ReturnBadRequestError(c, errorcode.NotFound, "Not found")
}

// ReturnError
// nolint
func ReturnError(c *gin.Context, code int, errorCode string, errors any) {
	traceID := c.Keys[constant.TraceID].(string)
	responseData := response.Response{
		Data:       nil,
		Message:    errorcode.GetMsg(errorCode),
		Errors:     errors,
		ErrorCode:  errorCode,
		StatusCode: code,
		TraceID:    traceID,
		Success:    false,
	}
	// 200
	c.JSON(http.StatusOK, responseData)
	c.Abort()
	return
}

// BaseError
// nolint
func BaseError(code int, errorCode string, errors any) response.Response {
	return response.Response{
		Data:       nil,
		Message:    errorcode.GetMsg(errorCode),
		Errors:     errors,
		ErrorCode:  errorCode,
		StatusCode: code,
		Success:    false,
	}
}

// ReturnBadRequestError Using for share
// nolint
func ReturnBadRequestError(c *gin.Context, errorCode string, errors any) {
	ReturnError(c, http.StatusBadRequest, errorCode, errors)
}

// ReturnInternalServerError
// nolint
func ReturnInternalServerError(c *gin.Context, errorCode string, errors any) {
	ReturnError(c, http.StatusInternalServerError, errorCode, errors)
}

// ReturnUnauthorizedError
// nolint
func ReturnUnauthorizedError(c *gin.Context, errorCode string, errors any) {
	ReturnError(c, http.StatusUnauthorized, errorCode, errors)
}

// BadRequestError Using with panic
// nolint
func BadRequestError(errorCode string, errors any) response.Response {
	return BaseError(http.StatusBadRequest, errorCode, errors)
}

// InternalServerError
// nolint
func InternalServerError(errorCode string, errors any) response.Response {
	return BaseError(http.StatusInternalServerError, errorCode, errors)
}

// UnauthorizedError
// nolint
func UnauthorizedError(errorCode string, errors any) response.Response {
	return BaseError(http.StatusUnauthorized, errorCode, errors)
}

// StatusNotFoundError
// nolint
func StatusNotFoundError(errorCode string, errors any) response.Response {
	return BaseError(http.StatusNotFound, errorCode, errors)
}
