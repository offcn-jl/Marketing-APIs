/*
   @Time : 2020/4/19 1:36 下午
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : main
   @Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/offcn-jl/chaos-go-scf"
	"github.com/offcn-jl/chaos-go-scf/fake-http"
	"io"
	"os"
)

// 测试配置日志文件
func init() {
	// 设置记录日志到文件。
	f, _ := os.Create("1.log")
	chaos.DefaultWriter = io.MultiWriter(f, os.Stdout) // 同时将日志写入文件和控制台
}

func main() {
	defer func() { // 在程序结束时关闭 ORM 的连接
		//orm.Mysql.Chaos.Close()
		//orm.Mysql.Events.Close()
	}()
	r := chaos.Default()

	r.Use(GetSurplusHandler)

	r.Run()
}

/*
 * 获取剩余量
 */
func GetSurplusHandler(c *chaos.Context) {
	//count := 408
	//orm.Mysql.Chaos.Model(structs.EventsGift{}).Where("Name = ? AND Phone IS NULL", c.Param("Name")).Count(&count)
	fmt.Println("666")

	fmt.Println(c.Request)
	fmt.Println(c.Response)

	// 测试设置 header
	c.Header("tttt-tes", "99999")
	c.Header("ttttte-s11", "99999")

	// 返回数据
	c.JSON(http.StatusOK, c.Request)
	//c.AbortWithStatus(count)
}
