package common

import "encoding/json"

type Response struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Status  bool            `json:"status"`
	TTL     int             `json:"ttl,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}


type NavInfo struct {
	EmailVerified      int                    `json:"email_verified"`       // 是否验证邮箱地址 0:未验证 1:已验证
	Face               string                 `json:"face"`                 // 用户头像url
	LevelInfo          *NavInfoLevel          `json:"level_info"`           // 等级信息
	MID                int64                  `json:"mid"`                  // 用户mid
	MobileVerified     int                    `json:"mobile_verified"`      // 是否验证手机号 0:未验证 1:已验证
	Money              float64                `json:"money"`                // 拥有硬币数
	Moral              int                    `json:"moral"`                // 当前节操值 上限为70
	Official           *NavInfoOfficial       `json:"official"`             // 认证信息
	OfficialVerify     *NavInfoOfficialVerify `json:"officialVerify"`       // 认证信息2
	Pendant            *NavInfoPendant        `json:"pendant"`              // 头像框信息
	Scores             int                    `json:"scores"`               // 0 作用尚不明确
	Uname              string                 `json:"uname"`                // 用户昵称
	VipDueDate         int64                  `json:"vipDueDate"`           // 会员到期时间 毫秒 时间戳(东八区)
	VipStatus          int                    `json:"vipStatus"`            // 会员开通状态 0:无 1:有
	VipType            int                    `json:"vipType"`              // 会员类型 0:无 1:月度大会员 2:年度及以上大会员
	VipPayType         int                    `json:"vip_pay_type"`         // 会员开通状态	0:无 1:有
	VipThemeType       int                    `json:"vip_theme_type"`       // 0 作用尚不明确
	VipLabel           *NavInfoVipLabel       `json:"vip_label"`            // 会员标签
	VipAvatarSubscript int                    `json:"vip_avatar_subscript"` // 是否显示会员图标 0:不显示 1:显示
	VipNicknameColor   string                 `json:"vip_nickname_color"`   // 会员昵称颜色	颜色码 如#FFFFFF
	Wallet             *NavInfoWallet         `json:"wallet"`               // B币钱包信息
	HasShop            bool                   `json:"has_shop"`             // 是否拥有推广商品 false:无 true:有
	ShopURL            string                 `json:"shop_url"`             // 商品推广页面url
	AllowanceCount     int                    `json:"allowance_count"`      // 0 作用尚不明确
	AnswerStatus       int                    `json:"answer_status"`        // 0 作用尚不明确
}
type NavInfoLevel struct {
	CurrentLevel int `json:"current_level"` // 当前等级
	CurrentMin   int `json:"current_min"`   // 当前等级经验最低值
	CurrentExp   int `json:"current_exp"`   // 当前经验
	NextExp      int `json:"next_exp"`      // 升级下一等级需达到的经验
}
type NavInfoOfficial struct {
	// 认证类型
	//
	// 0:无
	//
	// 1 2 7:个人认证
	//
	// 3 4 5 6:机构认证
	Role  int    `json:"role"`
	Title string `json:"title"` // 认证信息 无为空
	Desc  string `json:"desc"`  // 认证备注 无为空
	Type  int    `json:"type"`  // 是否认证 -1:无 0:认证
}
type NavInfoOfficialVerify struct {
	Type int    `json:"type"` // 是否认证 -1:无 0:认证
	Desc string `json:"desc"` // 认证信息 无为空
}
type NavInfoPendant struct {
	PID    int64  // 挂件id
	Name   string // 挂件名称
	Image  string // 挂件图片url
	Expire int    // 0 作用尚不明确
}
type NavInfoVipLabel struct {
	Path string `json:"path"` // 空 作用尚不明确
	Text string `json:"text"` // 会员名称
	// 会员标签
	//
	// vip:大会员
	//
	// annual_vip:年度大会员
	//
	// ten_annual_vip:十年大会员
	//
	// hundred_annual_vip:百年大会员
	LabelTheme string `json:"label_theme"`
}
type NavInfoWallet struct {
	MID           int64   `json:"mid"`             // 登录用户mid
	BcoinBalance  float64 `json:"bcoin_balance"`   // 拥有B币数
	CouponBalance float64 `json:"coupon_balance"`  // 每月奖励B币数
	CouponDueTime int     `json:"coupon_due_time"` // 0 作用尚不明确
}

type NavStat struct {
	Following    int `json:"following"`     // 关注数
	Follower     int `json:"follower"`      // 粉丝数
	DynamicCount int `json:"dynamic_count"` // 发布动态数
}
