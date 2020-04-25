/*
   @Time : 2020/4/24 9:48 上午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : 2020-gift
   @Software: GoLand
*/

package structs

import "github.com/jinzhu/gorm"

/**
 * Events Gift 活动 奖品发放
 */
type EventsGift struct {
	gorm.Model
	Name        string `gorm:"not null"` // 奖品名称
	Detail      string // 奖品详情
	Phone       string // 领取人手机号
	ReceiveInfo string // 领取信息
	// 平台日志
	SourceIP   string
	ApiVersion string
}
