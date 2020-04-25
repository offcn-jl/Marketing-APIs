/*
   @Time : 2020/4/24 10:27 上午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : main
   @Software: GoLand
   @Package: consume
*/

package main

import (
	"github.com/offcn-jl/chaos-go-scf"
	"github.com/offcn-jl/chaos-go-scf/fake-http"
	"serverless/go-common/database/orm"
	"serverless/go-common/database/orm/structs"
	"serverless/go-common/handler"
)

var apiVersion = "0.1.0"

func main() {
	r := chaos.Default()
	r.Use(handler.AddVersions(apiVersion), MainHandler)
	r.Run()
}

func MainHandler(c *chaos.Context) {
	o := orm.New()

	giftCheckInfo := structs.EventsGift{}
	o.PostgreSQL.Marketing.Where("Name = ? AND Phone = ?", c.Param("Name"), c.Param("Phone")).First(&giftCheckInfo)
	if giftCheckInfo.ReceiveInfo != "" {
		c.JSON(http.StatusForbidden, chaos.H{"Code": -1, "Error": "物料已被使用，详情 : " + giftCheckInfo.ReceiveInfo})
	} else {
		o.PostgreSQL.Marketing.Model(&structs.EventsGift{}).Where("Name = ? AND Phone = ?", c.Param("Name"), c.Param("Phone")).Update("receive_info", c.Param("ReceiveInfo"))
	}

	defer func() {
		o.Close() // 在程序结束时关闭 ORM 的连接
	}()
}
