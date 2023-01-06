package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/zaihui/go-hutils"
)

// Test 门店详细信息.
type Test struct {
	hutils.BaseSchema
}

type DeliveryConfig struct {
	PrepareMealTimes   int         `json:"prepare_meal_times"` // 备餐时间
	Level              int         `json:"level"`              // 目前有0(基础), 1(第一档), 2(第二档), 3(第三档) 与配送相关
	BasePrice          string      `json:"base_price"`         // 20 起送价
	FreePrice          string      `json:"free_price"`         // 免费配送价格
	LevelConfig        LevelConfig `json:"level_config"`
	DeliveryAreas      int         `json:"delivery_areas"`       // 配送区域半径 单位m
	DeliveryPartners   []string    `json:"delivery_partners"`    // 配送方式
	EnableAutoDelivery bool        `json:"enable_auto_delivery"` // 是否自动呼叫配送
	EnableBooking      bool        `json:"enable_booking"`       // 是否开启预约配送功能
	BookingConfigs     BookConfig  `json:"booking_configs"`      // 预约配置
	DeliveryCallTime   int         `json:"delivery_call_time"`   // 呼叫第三方配送时长, 单位分钟
	DelayCallTime      int         `json:"delay_call_time"`      // 延迟呼叫配送，单位分钟
}

type PayConfig struct {
	EnableInStore bool `json:"enable_in_store"` // 到店支付
	EnableAlipay  bool `json:"enable_alipay"`   // 支付宝支付
	EnableWechat  bool `json:"enable_wechat"`   // 微信支付
	EnableApple   bool `json:"enable_apple"`    // 苹果支付
	EnableGoogle  bool `json:"enable_google"`   // google支付
	EnablePayNow  bool `json:"enable_pay_now"`  // paynow支付
	EnableCard    bool `json:"enable_card"`     // 卡支付
}

type LevelConfig struct {
	Price              string `json:"price"`                // 1km 以内的配送费 8 新币
	Additional         string `json:"additional"`           // 额外配送费, 每 km 为 0.35
	Distance           int    `json:"distance"`             // 基础配送距离 1km
	DeliveryTime       int    `json:"delivery_time"`        // 基础配送时长
	FirstPrice         string `json:"first_price"`          // 第一档配送费
	FirstDistance      int    `json:"first_distance"`       // 第一档配送距离
	FirstDeliveryTime  int    `json:"first_delivery_time"`  // 第一档配送时长
	SecondPrice        string `json:"second_price"`         // 第二档配送费
	SecondDistance     int    `json:"second_distance"`      // 第二档配送距离
	SecondDeliveryTime int    `json:"second_delivery_time"` // 第二档配送时长
	ThirdPrice         string `json:"third_price"`          // 第三档配送费
	ThirdDistance      int    `json:"third_distance"`       // 第三档配送距离
	ThirdDeliveryTime  int    `json:"third_delivery_time"`  // 第三档配送时长
}

type BookConfig struct {
	BookingDays int `json:"booking_days"` // 最大预约天数
	NoticeTimes int `json:"notice_times"` // 预订单提醒时间
	PrintScene  int `json:"print_scene"`  // 打印时机
}

type RingConfig struct {
	NewOrder                  string `json:"new_order"`                     // 新订单(即时配送，预约配送，即时自提)
	AutoAcceptOrder           string `json:"auto_accept_order"`             // 自动接单
	OverdueOrder              string `json:"overdue_order"`                 // 订单过期
	DeliverException          string `json:"deliver_exception"`             // 配送异常
	PrinterDisconnect         string `json:"printer_disconnect"`            // 打印机断开
	NetworkDisconnect         string `json:"network_disconnect"`            // 网络断开
	BookingNotice             string `json:"booking_notice"`                // 预订单配送提醒
	SelfPickupNewOrder        string `json:"self_pickup_new_order"`         // 自提新订单
	SelfPickupAutoAcceptOrder string `json:"self_pickup_auto_accept_order"` // 自提自动接单
	SelfPickupBookingNotice   string `json:"self_pickup_booking_notice"`    // 自提预订单提醒
}

type SelfPickupConfig struct {
	BasePrice        string     `json:"base_price"`         // 起定价, 对等于配送的起送价
	PrepareMealTimes int        `json:"prepare_meal_times"` // 备餐时间
	AutoAcceptOrder  bool       `json:"auto_accept_order"`  // 自动接单
	EnableBooking    bool       `json:"enable_booking"`     // 是否预约配置
	BookingConfigs   BookConfig `json:"booking_configs"`    //  预约配置
}

type Banners struct {
	Thumb []string `json:"thumb"`
	Raw   []string `json:"raw"`
}

type DineInConfigs struct {
	PayFirst PayFirst `json:"pay_first"` // 先付
	PayLater PayLater `json:"pay_later"` // 后付
}

type DealsConfig struct {
	EnableCashier bool `json:"enable_cashier"` // 是否作为收银
}

type PrintConfig struct {
	ScanDisplayStyle int `json:"scan_display_style"` // 小票展示形式 1：为默认，2：暂时为food_hub样式
}

type PayFirst struct {
	CleanTableTimes  int `json:"clean_table_times"`  // 清台时间，单位m
	PrepareMealTimes int `json:"prepare_meal_times"` // 备餐时间
}

type PayLater struct {
	CleanTableTimes  int `json:"clean_table_times"`  // 清台时间，单位m
	PrepareMealTimes int `json:"prepare_meal_times"` // 备餐时间
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("门店id"),
		field.String("name").Validate(MaxRuneCount(LenNormal)).Comment("门店名称"),
		field.Bool("enable_takeaway").Default(FALSE).Comment("是否开启外卖"), // 整个门店全局业务关闭
		field.String("phone").Validate(MaxRuneCount(LenPhone)).Comment("手机号"),
		field.String("announcement").Validate(MaxRuneCount(LenDesc)).Comment("门店公告"),
		field.Bool("enable_auto_accept").Default(FALSE).Comment("是否开启自动接单"),
		field.JSON("delivery_configs", DeliveryConfig{}).Comment("外卖配置"),
		field.JSON("pay_configs", PayConfig{}).Comment("支付配置"),
		field.Int("print_times").Nillable().Optional().Comment("打印次数"),
		field.JSON("ring_configs", RingConfig{}).Comment("铃声配置"),
		field.Int("business_group_id").Comment("商户id"),
		field.String("business_group_uid").Comment("商户uid"),
		field.JSON("dish_categories", []string{}).Comment("菜品品类"),
		field.Time("end_of_takeaway").Optional().Comment("外卖订购有效期"), // 外卖服务按门店收费
		field.Int("mode").Default(1).Comment("可选模式"),                // 只关闭某个 配送、自提、业务
		field.JSON("self_pickup_configs", SelfPickupConfig{}).Comment("自取配置"),
		field.Int("server_id").Comment("shiva的门店 ID"),
		field.String("image").Validate(MaxRuneCount(LenURL)).Comment("门店图片"),
		field.String("address").Validate(MaxRuneCount(LenDesc)).Comment("门店地址"),
		field.String("latitude").Validate(MaxRuneCount(LenTwenty)).Comment("纬度"),
		field.String("longitude").Validate(MaxRuneCount(LenTwenty)).Comment("经度"),
		field.JSON("banners", Banners{}).Comment("门店广告位"),
		field.Int("sort").Default(0).Comment("自定义排序"),
		field.String("pay_mode").Default("pay_first").Comment("支付模式"), // 先付pay_first，后付模式pay_later
		field.JSON("dine_in_configs", DineInConfigs{}).Optional().Comment("堂食配置"),
		field.JSON("deals_config", DealsConfig{}).Optional().Comment("套餐配置"),
		field.JSON("print_config", PrintConfig{}).Default(PrintConfig{ScanDisplayStyle: 1}).Optional().Comment("小票打印相关字段"),
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the Test.
func (Test) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "intl_business"},
	}
}
