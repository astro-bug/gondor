package handlers

import (
	"strings"

	"github.com/astro-bug/gondor/webapi/fakes"
	"github.com/gofiber/fiber"
)

// 订单列表
func OrderListHandler(ctx *fiber.Ctx) {
	result := fakes.ReduceBlanks(`{"code":200, "total":20, "data":[` + fakes.GenOrder() + `]}`)
	ctx.Type("json").SendBytes([]byte(result))
}

// 查找用户名
func SearchUserHandler(ctx *fiber.Ctx) {
	var names []string
	match := strings.ToLower(ctx.Query("name"))
	for _, name := range fakes.FakeUsers {
		lowerName := strings.ToLower(name)
		if strings.Contains(lowerName, match) {
			names = append(names, name)
		}
	}
	result := fiber.Map{
		"code": 200,
		"data": fiber.Map{
			"items": names,
		},
	}
	ctx.JSON(result)
}
