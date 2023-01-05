// Code generated by ent, DO NOT EDIT.

package entschema

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/zaihui/ent-factory/gen/entschema/test"
	"github.com/zaihui/ent-factory/spec/schema"
)

// Test is the model entity for the Test schema.
type Test struct {
	config `json:"-"`
	// ID of the ent.
	// 门店id
	ID int `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeactivatedAt holds the value of the "deactivated_at" field.
	DeactivatedAt *time.Time `json:"deactivated_at,omitempty"`
	// 门店名称
	Name string `json:"name,omitempty"`
	// 是否开启外卖
	EnableTakeaway bool `json:"enable_takeaway,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 门店公告
	Announcement string `json:"announcement,omitempty"`
	// 是否开启自动接单
	EnableAutoAccept bool `json:"enable_auto_accept,omitempty"`
	// 外卖配置
	DeliveryConfigs schema.DeliveryConfig `json:"delivery_configs,omitempty"`
	// 支付配置
	PayConfigs schema.PayConfig `json:"pay_configs,omitempty"`
	// 打印次数
	PrintTimes int `json:"print_times,omitempty"`
	// 铃声配置
	RingConfigs schema.RingConfig `json:"ring_configs,omitempty"`
	// 商户id
	BusinessGroupID int `json:"business_group_id,omitempty"`
	// 商户uid
	BusinessGroupUID string `json:"business_group_uid,omitempty"`
	// 菜品品类
	DishCategories []string `json:"dish_categories,omitempty"`
	// 外卖订购有效期
	EndOfTakeaway time.Time `json:"end_of_takeaway,omitempty"`
	// 可选模式
	Mode int `json:"mode,omitempty"`
	// 自取配置
	SelfPickupConfigs schema.SelfPickupConfig `json:"self_pickup_configs,omitempty"`
	// shiva的门店 ID
	ServerID int `json:"server_id,omitempty"`
	// 门店图片
	Image string `json:"image,omitempty"`
	// 门店地址
	Address string `json:"address,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty"`
	// 门店广告位
	Banners schema.Banners `json:"banners,omitempty"`
	// 自定义排序
	Sort int `json:"sort,omitempty"`
	// 支付模式
	PayMode string `json:"pay_mode,omitempty"`
	// 堂食配置
	DineInConfigs schema.DineInConfigs `json:"dine_in_configs,omitempty"`
	// 套餐配置
	DealsConfig schema.DealsConfig `json:"deals_config,omitempty"`
	// 小票打印相关字段
	PrintConfig schema.PrintConfig `json:"print_config,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Test) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case test.FieldDeliveryConfigs, test.FieldPayConfigs, test.FieldRingConfigs, test.FieldDishCategories, test.FieldSelfPickupConfigs, test.FieldBanners, test.FieldDineInConfigs, test.FieldDealsConfig, test.FieldPrintConfig:
			values[i] = new([]byte)
		case test.FieldEnableTakeaway, test.FieldEnableAutoAccept:
			values[i] = new(sql.NullBool)
		case test.FieldID, test.FieldPrintTimes, test.FieldBusinessGroupID, test.FieldMode, test.FieldServerID, test.FieldSort:
			values[i] = new(sql.NullInt64)
		case test.FieldUID, test.FieldName, test.FieldPhone, test.FieldAnnouncement, test.FieldBusinessGroupUID, test.FieldImage, test.FieldAddress, test.FieldLatitude, test.FieldLongitude, test.FieldPayMode:
			values[i] = new(sql.NullString)
		case test.FieldCreatedAt, test.FieldUpdatedAt, test.FieldDeactivatedAt, test.FieldEndOfTakeaway:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Test", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Test fields.
func (t *Test) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case test.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case test.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				t.UID = value.String
			}
		case test.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case test.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case test.FieldDeactivatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deactivated_at", values[i])
			} else if value.Valid {
				t.DeactivatedAt = new(time.Time)
				*t.DeactivatedAt = value.Time
			}
		case test.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case test.FieldEnableTakeaway:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable_takeaway", values[i])
			} else if value.Valid {
				t.EnableTakeaway = value.Bool
			}
		case test.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				t.Phone = value.String
			}
		case test.FieldAnnouncement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field announcement", values[i])
			} else if value.Valid {
				t.Announcement = value.String
			}
		case test.FieldEnableAutoAccept:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable_auto_accept", values[i])
			} else if value.Valid {
				t.EnableAutoAccept = value.Bool
			}
		case test.FieldDeliveryConfigs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_configs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.DeliveryConfigs); err != nil {
					return fmt.Errorf("unmarshal field delivery_configs: %w", err)
				}
			}
		case test.FieldPayConfigs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field pay_configs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.PayConfigs); err != nil {
					return fmt.Errorf("unmarshal field pay_configs: %w", err)
				}
			}
		case test.FieldPrintTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field print_times", values[i])
			} else if value.Valid {
				t.PrintTimes = int(value.Int64)
			}
		case test.FieldRingConfigs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ring_configs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.RingConfigs); err != nil {
					return fmt.Errorf("unmarshal field ring_configs: %w", err)
				}
			}
		case test.FieldBusinessGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field business_group_id", values[i])
			} else if value.Valid {
				t.BusinessGroupID = int(value.Int64)
			}
		case test.FieldBusinessGroupUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_group_uid", values[i])
			} else if value.Valid {
				t.BusinessGroupUID = value.String
			}
		case test.FieldDishCategories:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field dish_categories", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.DishCategories); err != nil {
					return fmt.Errorf("unmarshal field dish_categories: %w", err)
				}
			}
		case test.FieldEndOfTakeaway:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_of_takeaway", values[i])
			} else if value.Valid {
				t.EndOfTakeaway = value.Time
			}
		case test.FieldMode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				t.Mode = int(value.Int64)
			}
		case test.FieldSelfPickupConfigs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field self_pickup_configs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.SelfPickupConfigs); err != nil {
					return fmt.Errorf("unmarshal field self_pickup_configs: %w", err)
				}
			}
		case test.FieldServerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field server_id", values[i])
			} else if value.Valid {
				t.ServerID = int(value.Int64)
			}
		case test.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				t.Image = value.String
			}
		case test.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				t.Address = value.String
			}
		case test.FieldLatitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				t.Latitude = value.String
			}
		case test.FieldLongitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				t.Longitude = value.String
			}
		case test.FieldBanners:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field banners", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Banners); err != nil {
					return fmt.Errorf("unmarshal field banners: %w", err)
				}
			}
		case test.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				t.Sort = int(value.Int64)
			}
		case test.FieldPayMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pay_mode", values[i])
			} else if value.Valid {
				t.PayMode = value.String
			}
		case test.FieldDineInConfigs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field dine_in_configs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.DineInConfigs); err != nil {
					return fmt.Errorf("unmarshal field dine_in_configs: %w", err)
				}
			}
		case test.FieldDealsConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field deals_config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.DealsConfig); err != nil {
					return fmt.Errorf("unmarshal field deals_config: %w", err)
				}
			}
		case test.FieldPrintConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field print_config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.PrintConfig); err != nil {
					return fmt.Errorf("unmarshal field print_config: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Test.
// Note that you need to call Test.Unwrap() before calling this method if this Test
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Test) Update() *TestUpdateOne {
	return (&TestClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Test entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Test) Unwrap() *Test {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("entschema: Test is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Test) String() string {
	var builder strings.Builder
	builder.WriteString("Test(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("uid=")
	builder.WriteString(t.UID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := t.DeactivatedAt; v != nil {
		builder.WriteString("deactivated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("enable_takeaway=")
	builder.WriteString(fmt.Sprintf("%v", t.EnableTakeaway))
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(t.Phone)
	builder.WriteString(", ")
	builder.WriteString("announcement=")
	builder.WriteString(t.Announcement)
	builder.WriteString(", ")
	builder.WriteString("enable_auto_accept=")
	builder.WriteString(fmt.Sprintf("%v", t.EnableAutoAccept))
	builder.WriteString(", ")
	builder.WriteString("delivery_configs=")
	builder.WriteString(fmt.Sprintf("%v", t.DeliveryConfigs))
	builder.WriteString(", ")
	builder.WriteString("pay_configs=")
	builder.WriteString(fmt.Sprintf("%v", t.PayConfigs))
	builder.WriteString(", ")
	builder.WriteString("print_times=")
	builder.WriteString(fmt.Sprintf("%v", t.PrintTimes))
	builder.WriteString(", ")
	builder.WriteString("ring_configs=")
	builder.WriteString(fmt.Sprintf("%v", t.RingConfigs))
	builder.WriteString(", ")
	builder.WriteString("business_group_id=")
	builder.WriteString(fmt.Sprintf("%v", t.BusinessGroupID))
	builder.WriteString(", ")
	builder.WriteString("business_group_uid=")
	builder.WriteString(t.BusinessGroupUID)
	builder.WriteString(", ")
	builder.WriteString("dish_categories=")
	builder.WriteString(fmt.Sprintf("%v", t.DishCategories))
	builder.WriteString(", ")
	builder.WriteString("end_of_takeaway=")
	builder.WriteString(t.EndOfTakeaway.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mode=")
	builder.WriteString(fmt.Sprintf("%v", t.Mode))
	builder.WriteString(", ")
	builder.WriteString("self_pickup_configs=")
	builder.WriteString(fmt.Sprintf("%v", t.SelfPickupConfigs))
	builder.WriteString(", ")
	builder.WriteString("server_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ServerID))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(t.Image)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(t.Address)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(t.Latitude)
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(t.Longitude)
	builder.WriteString(", ")
	builder.WriteString("banners=")
	builder.WriteString(fmt.Sprintf("%v", t.Banners))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", t.Sort))
	builder.WriteString(", ")
	builder.WriteString("pay_mode=")
	builder.WriteString(t.PayMode)
	builder.WriteString(", ")
	builder.WriteString("dine_in_configs=")
	builder.WriteString(fmt.Sprintf("%v", t.DineInConfigs))
	builder.WriteString(", ")
	builder.WriteString("deals_config=")
	builder.WriteString(fmt.Sprintf("%v", t.DealsConfig))
	builder.WriteString(", ")
	builder.WriteString("print_config=")
	builder.WriteString(fmt.Sprintf("%v", t.PrintConfig))
	builder.WriteByte(')')
	return builder.String()
}

// Tests is a parsable slice of Test.
type Tests []*Test

func (t Tests) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}