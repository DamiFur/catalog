package validators

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/damifur/catalog/utils/errors"
	"github.com/damifur/catalog/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	intErr = logger.Sprintf("Can't parse string '%s' to int")
)

type Validator struct {
	err *errors.Error
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Error() *errors.Error {
	if v.err != nil {
		return v.err
	}
	return nil
}

func Body(c *gin.Context, format string) {

	if c.Request.Body != nil {

		buf, bodyErr := ioutil.ReadAll(c.Request.Body)
		if bodyErr != nil {
			logger.Warnf(format, bodyErr.Error())
			return
		}
		c.Request.Body.Close()

		// Use `buf` as required
		logger.Infof(format, buf)

		// Pass `buf` (possibly modified!) on to the "next" handler.
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf))
	}
}

func Int(s string) (int, *errors.Error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.BadRequest(intErr(s))
	}
	return i, nil
}

// SafeBody checks if it is possible to unmarshall request body into the given struct
func (v *Validator) SafeBody(c *gin.Context, instance interface{}) *Validator {

	if v.err != nil {
		return nil
	}

	Body(c, fmt.Sprintf("%T :%%s", instance))

	if c.Request.Body == nil {
		v.err = errors.BadRequest("Empty body")
		return nil
	}
	if err := c.BindJSON(instance); err != nil {
		v.err = errors.BadRequest(fmt.Sprintf("Can't parse body: %s", err.Error()))
		return nil
	}
	return v
}
