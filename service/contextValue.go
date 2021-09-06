package service

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"strconv"
	"strings"
)

func GetContextValue(ctx *fiber.Ctx, name string, unescape ...bool) (ContextValue, bool) {

	names := strings.Fields(name)

	var val string

	for _, valName := range names {
		val = ctx.Params(valName)
		if len(val) > 0 {
			break
		}

		val = ctx.Query(valName)
		if len(val) > 0 {
			break
		}

	}

	if len(val) == 0 {
		return "", false
	}

	urlUnescape := true

	if len(unescape) > 0 {
		urlUnescape = unescape[0]
	}

	if urlUnescape {
		val, _ = url.QueryUnescape(val)
	}

	return ContextValue(val), len(val) > 0
}

func GetContextValueOrNil(ctx *fiber.Ctx, name string, unescape ...bool) ContextValue {

	val, _ := GetContextValue(ctx, name, unescape...)

	return val
}

func GetClusterID(ctx *fiber.Ctx) (ContextValue, bool) {

	return GetContextValue(ctx, "cluster cluster-id", true)

}

func GetInfobaseID(ctx *fiber.Ctx) (ContextValue, bool) {

	return GetContextValue(ctx, "infobase infobase-id", true)

}

type ContextValue string

func (val ContextValue) Empty() bool {
	return len(val) == 0
}

func (val ContextValue) NotEmpty() bool {
	return !val.Empty()
}

func (val ContextValue) String() string {
	return string(val)
}

func (val ContextValue) Bool(defaultVal ...bool) bool {

	var defVal bool

	if len(defaultVal) > 0 {
		defVal = defaultVal[0]
	}

	valB, err := strconv.ParseBool(val.String())

	if err != nil {
		return defVal
	}

	return valB
}

func (val ContextValue) UUID() (uuid.UUID, error) {
	return uuid.FromString(val.String())
}

func (val ContextValue) NilUUID() bool {
	return len(val) > 0 &&
		uuid.FromStringOrNil(val.String()) == uuid.Nil
}

func (val ContextValue) NotNilUUID() bool {
	return !val.NilUUID()
}
