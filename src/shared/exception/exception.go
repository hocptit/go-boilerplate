package exception

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boilerplate/src/constants/error_code"
	"go-boilerplate/src/shared/response"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
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
			data, err := ioutil.ReadFile(file)
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
func RecoveryError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				// check validator class
				case response.Response:
					errData, _ := err.(response.Response)
					ReturnBadRequestError(c, errData.ErrorCode, errData.Errors)
					return
				default:
					stack := stack(3)
					// todo: logger
					fmt.Printf("[Recovery] %s panic recovered:\n%s\n%s%s",
						timeFormat(time.Now()), err, stack, "\033[0m")
					ReturnInternalServerError(c, error_code.INTERNAL_SERVER, err)
					return
				}
			}
		}()
		c.Next()
	}
}

func ReturnError(c *gin.Context, code int, errorCode string, errors any) {
	responseData := response.Response{
		Data:       nil,
		Message:    error_code.GetMsg(errorCode),
		Errors:     errors,
		ErrorCode:  errorCode,
		StatusCode: code,
		Success:    false,
	}
	// 200
	c.JSON(http.StatusOK, responseData)
	return
}

func BaseError(code int, errorCode string, errors any) response.Response {
	return response.Response{
		Data:       nil,
		Message:    error_code.GetMsg(errorCode),
		Errors:     errors,
		ErrorCode:  errorCode,
		StatusCode: code,
		Success:    false,
	}
}

func ReturnBadRequestError(c *gin.Context, errorCode string, errors any) {
	var code int = http.StatusBadRequest
	ReturnError(c, code, errorCode, errors)
}

func ReturnInternalServerError(c *gin.Context, errorCode string, errors any) {
	var code int = http.StatusInternalServerError
	ReturnError(c, code, errorCode, errors)
}

func ReturnUnauthorizedError(c *gin.Context, errorCode string, errors any) {
	var code int = http.StatusUnauthorized
	ReturnError(c, code, errorCode, errors)
}

func BadRequestError(errorCode string, errors any) response.Response {
	var code int = http.StatusBadRequest
	return BaseError(code, errorCode, errors)
}

func InternalServerError(errorCode string, errors any) response.Response {
	var code int = http.StatusInternalServerError
	return BaseError(code, errorCode, errors)
}

func UnauthorizedError(errorCode string, errors any) response.Response {
	var code int = http.StatusUnauthorized
	return BaseError(code, errorCode, errors)
}
