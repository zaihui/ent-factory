// Code generated by ent, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/zaihui/ent-factory/gen/entschema/predicate"
	"github.com/zaihui/ent-factory/gen/entschema/test"
	"github.com/zaihui/ent-factory/spec/schema"
)

// TestUpdate is the builder for updating Test entities.
type TestUpdate struct {
	config
	hooks    []Hook
	mutation *TestMutation
}

// Where appends a list predicates to the TestUpdate builder.
func (tu *TestUpdate) Where(ps ...predicate.Test) *TestUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TestUpdate) SetUpdatedAt(t time.Time) *TestUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (tu *TestUpdate) SetDeactivatedAt(t time.Time) *TestUpdate {
	tu.mutation.SetDeactivatedAt(t)
	return tu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (tu *TestUpdate) SetNillableDeactivatedAt(t *time.Time) *TestUpdate {
	if t != nil {
		tu.SetDeactivatedAt(*t)
	}
	return tu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (tu *TestUpdate) ClearDeactivatedAt() *TestUpdate {
	tu.mutation.ClearDeactivatedAt()
	return tu
}

// SetName sets the "name" field.
func (tu *TestUpdate) SetName(s string) *TestUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetEnableTakeaway sets the "enable_takeaway" field.
func (tu *TestUpdate) SetEnableTakeaway(b bool) *TestUpdate {
	tu.mutation.SetEnableTakeaway(b)
	return tu
}

// SetNillableEnableTakeaway sets the "enable_takeaway" field if the given value is not nil.
func (tu *TestUpdate) SetNillableEnableTakeaway(b *bool) *TestUpdate {
	if b != nil {
		tu.SetEnableTakeaway(*b)
	}
	return tu
}

// SetPhone sets the "phone" field.
func (tu *TestUpdate) SetPhone(s string) *TestUpdate {
	tu.mutation.SetPhone(s)
	return tu
}

// SetAnnouncement sets the "announcement" field.
func (tu *TestUpdate) SetAnnouncement(s string) *TestUpdate {
	tu.mutation.SetAnnouncement(s)
	return tu
}

// SetEnableAutoAccept sets the "enable_auto_accept" field.
func (tu *TestUpdate) SetEnableAutoAccept(b bool) *TestUpdate {
	tu.mutation.SetEnableAutoAccept(b)
	return tu
}

// SetNillableEnableAutoAccept sets the "enable_auto_accept" field if the given value is not nil.
func (tu *TestUpdate) SetNillableEnableAutoAccept(b *bool) *TestUpdate {
	if b != nil {
		tu.SetEnableAutoAccept(*b)
	}
	return tu
}

// SetDeliveryConfigs sets the "delivery_configs" field.
func (tu *TestUpdate) SetDeliveryConfigs(sc schema.DeliveryConfig) *TestUpdate {
	tu.mutation.SetDeliveryConfigs(sc)
	return tu
}

// SetPayConfigs sets the "pay_configs" field.
func (tu *TestUpdate) SetPayConfigs(sc schema.PayConfig) *TestUpdate {
	tu.mutation.SetPayConfigs(sc)
	return tu
}

// SetPrintTimes sets the "print_times" field.
func (tu *TestUpdate) SetPrintTimes(i int) *TestUpdate {
	tu.mutation.ResetPrintTimes()
	tu.mutation.SetPrintTimes(i)
	return tu
}

// SetNillablePrintTimes sets the "print_times" field if the given value is not nil.
func (tu *TestUpdate) SetNillablePrintTimes(i *int) *TestUpdate {
	if i != nil {
		tu.SetPrintTimes(*i)
	}
	return tu
}

// AddPrintTimes adds i to the "print_times" field.
func (tu *TestUpdate) AddPrintTimes(i int) *TestUpdate {
	tu.mutation.AddPrintTimes(i)
	return tu
}

// ClearPrintTimes clears the value of the "print_times" field.
func (tu *TestUpdate) ClearPrintTimes() *TestUpdate {
	tu.mutation.ClearPrintTimes()
	return tu
}

// SetRingConfigs sets the "ring_configs" field.
func (tu *TestUpdate) SetRingConfigs(sc schema.RingConfig) *TestUpdate {
	tu.mutation.SetRingConfigs(sc)
	return tu
}

// SetBusinessGroupID sets the "business_group_id" field.
func (tu *TestUpdate) SetBusinessGroupID(i int) *TestUpdate {
	tu.mutation.ResetBusinessGroupID()
	tu.mutation.SetBusinessGroupID(i)
	return tu
}

// AddBusinessGroupID adds i to the "business_group_id" field.
func (tu *TestUpdate) AddBusinessGroupID(i int) *TestUpdate {
	tu.mutation.AddBusinessGroupID(i)
	return tu
}

// SetBusinessGroupUID sets the "business_group_uid" field.
func (tu *TestUpdate) SetBusinessGroupUID(s string) *TestUpdate {
	tu.mutation.SetBusinessGroupUID(s)
	return tu
}

// SetDishCategories sets the "dish_categories" field.
func (tu *TestUpdate) SetDishCategories(s []string) *TestUpdate {
	tu.mutation.SetDishCategories(s)
	return tu
}

// AppendDishCategories appends s to the "dish_categories" field.
func (tu *TestUpdate) AppendDishCategories(s []string) *TestUpdate {
	tu.mutation.AppendDishCategories(s)
	return tu
}

// SetEndOfTakeaway sets the "end_of_takeaway" field.
func (tu *TestUpdate) SetEndOfTakeaway(t time.Time) *TestUpdate {
	tu.mutation.SetEndOfTakeaway(t)
	return tu
}

// SetNillableEndOfTakeaway sets the "end_of_takeaway" field if the given value is not nil.
func (tu *TestUpdate) SetNillableEndOfTakeaway(t *time.Time) *TestUpdate {
	if t != nil {
		tu.SetEndOfTakeaway(*t)
	}
	return tu
}

// ClearEndOfTakeaway clears the value of the "end_of_takeaway" field.
func (tu *TestUpdate) ClearEndOfTakeaway() *TestUpdate {
	tu.mutation.ClearEndOfTakeaway()
	return tu
}

// SetMode sets the "mode" field.
func (tu *TestUpdate) SetMode(i int) *TestUpdate {
	tu.mutation.ResetMode()
	tu.mutation.SetMode(i)
	return tu
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tu *TestUpdate) SetNillableMode(i *int) *TestUpdate {
	if i != nil {
		tu.SetMode(*i)
	}
	return tu
}

// AddMode adds i to the "mode" field.
func (tu *TestUpdate) AddMode(i int) *TestUpdate {
	tu.mutation.AddMode(i)
	return tu
}

// SetSelfPickupConfigs sets the "self_pickup_configs" field.
func (tu *TestUpdate) SetSelfPickupConfigs(spc schema.SelfPickupConfig) *TestUpdate {
	tu.mutation.SetSelfPickupConfigs(spc)
	return tu
}

// SetServerID sets the "server_id" field.
func (tu *TestUpdate) SetServerID(i int) *TestUpdate {
	tu.mutation.ResetServerID()
	tu.mutation.SetServerID(i)
	return tu
}

// AddServerID adds i to the "server_id" field.
func (tu *TestUpdate) AddServerID(i int) *TestUpdate {
	tu.mutation.AddServerID(i)
	return tu
}

// SetImage sets the "image" field.
func (tu *TestUpdate) SetImage(s string) *TestUpdate {
	tu.mutation.SetImage(s)
	return tu
}

// SetAddress sets the "address" field.
func (tu *TestUpdate) SetAddress(s string) *TestUpdate {
	tu.mutation.SetAddress(s)
	return tu
}

// SetLatitude sets the "latitude" field.
func (tu *TestUpdate) SetLatitude(s string) *TestUpdate {
	tu.mutation.SetLatitude(s)
	return tu
}

// SetLongitude sets the "longitude" field.
func (tu *TestUpdate) SetLongitude(s string) *TestUpdate {
	tu.mutation.SetLongitude(s)
	return tu
}

// SetBanners sets the "banners" field.
func (tu *TestUpdate) SetBanners(s schema.Banners) *TestUpdate {
	tu.mutation.SetBanners(s)
	return tu
}

// SetSort sets the "sort" field.
func (tu *TestUpdate) SetSort(i int) *TestUpdate {
	tu.mutation.ResetSort()
	tu.mutation.SetSort(i)
	return tu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (tu *TestUpdate) SetNillableSort(i *int) *TestUpdate {
	if i != nil {
		tu.SetSort(*i)
	}
	return tu
}

// AddSort adds i to the "sort" field.
func (tu *TestUpdate) AddSort(i int) *TestUpdate {
	tu.mutation.AddSort(i)
	return tu
}

// SetPayMode sets the "pay_mode" field.
func (tu *TestUpdate) SetPayMode(s string) *TestUpdate {
	tu.mutation.SetPayMode(s)
	return tu
}

// SetNillablePayMode sets the "pay_mode" field if the given value is not nil.
func (tu *TestUpdate) SetNillablePayMode(s *string) *TestUpdate {
	if s != nil {
		tu.SetPayMode(*s)
	}
	return tu
}

// SetDineInConfigs sets the "dine_in_configs" field.
func (tu *TestUpdate) SetDineInConfigs(sic schema.DineInConfigs) *TestUpdate {
	tu.mutation.SetDineInConfigs(sic)
	return tu
}

// SetNillableDineInConfigs sets the "dine_in_configs" field if the given value is not nil.
func (tu *TestUpdate) SetNillableDineInConfigs(sic *schema.DineInConfigs) *TestUpdate {
	if sic != nil {
		tu.SetDineInConfigs(*sic)
	}
	return tu
}

// ClearDineInConfigs clears the value of the "dine_in_configs" field.
func (tu *TestUpdate) ClearDineInConfigs() *TestUpdate {
	tu.mutation.ClearDineInConfigs()
	return tu
}

// SetDealsConfig sets the "deals_config" field.
func (tu *TestUpdate) SetDealsConfig(sc schema.DealsConfig) *TestUpdate {
	tu.mutation.SetDealsConfig(sc)
	return tu
}

// SetNillableDealsConfig sets the "deals_config" field if the given value is not nil.
func (tu *TestUpdate) SetNillableDealsConfig(sc *schema.DealsConfig) *TestUpdate {
	if sc != nil {
		tu.SetDealsConfig(*sc)
	}
	return tu
}

// ClearDealsConfig clears the value of the "deals_config" field.
func (tu *TestUpdate) ClearDealsConfig() *TestUpdate {
	tu.mutation.ClearDealsConfig()
	return tu
}

// SetPrintConfig sets the "print_config" field.
func (tu *TestUpdate) SetPrintConfig(sc schema.PrintConfig) *TestUpdate {
	tu.mutation.SetPrintConfig(sc)
	return tu
}

// SetNillablePrintConfig sets the "print_config" field if the given value is not nil.
func (tu *TestUpdate) SetNillablePrintConfig(sc *schema.PrintConfig) *TestUpdate {
	if sc != nil {
		tu.SetPrintConfig(*sc)
	}
	return tu
}

// ClearPrintConfig clears the value of the "print_config" field.
func (tu *TestUpdate) ClearPrintConfig() *TestUpdate {
	tu.mutation.ClearPrintConfig()
	return tu
}

// Mutation returns the TestMutation object of the builder.
func (tu *TestUpdate) Mutation() *TestMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TestUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TestUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TestUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TestUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := test.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TestUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := test.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entschema: validator failed for field "Test.name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Phone(); ok {
		if err := test.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`entschema: validator failed for field "Test.phone": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Announcement(); ok {
		if err := test.AnnouncementValidator(v); err != nil {
			return &ValidationError{Name: "announcement", err: fmt.Errorf(`entschema: validator failed for field "Test.announcement": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Image(); ok {
		if err := test.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`entschema: validator failed for field "Test.image": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Address(); ok {
		if err := test.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`entschema: validator failed for field "Test.address": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Latitude(); ok {
		if err := test.LatitudeValidator(v); err != nil {
			return &ValidationError{Name: "latitude", err: fmt.Errorf(`entschema: validator failed for field "Test.latitude": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Longitude(); ok {
		if err := test.LongitudeValidator(v); err != nil {
			return &ValidationError{Name: "longitude", err: fmt.Errorf(`entschema: validator failed for field "Test.longitude": %w`, err)}
		}
	}
	return nil
}

func (tu *TestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   test.Table,
			Columns: test.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: test.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(test.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.DeactivatedAt(); ok {
		_spec.SetField(test.FieldDeactivatedAt, field.TypeTime, value)
	}
	if tu.mutation.DeactivatedAtCleared() {
		_spec.ClearField(test.FieldDeactivatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(test.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.EnableTakeaway(); ok {
		_spec.SetField(test.FieldEnableTakeaway, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Phone(); ok {
		_spec.SetField(test.FieldPhone, field.TypeString, value)
	}
	if value, ok := tu.mutation.Announcement(); ok {
		_spec.SetField(test.FieldAnnouncement, field.TypeString, value)
	}
	if value, ok := tu.mutation.EnableAutoAccept(); ok {
		_spec.SetField(test.FieldEnableAutoAccept, field.TypeBool, value)
	}
	if value, ok := tu.mutation.DeliveryConfigs(); ok {
		_spec.SetField(test.FieldDeliveryConfigs, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.PayConfigs(); ok {
		_spec.SetField(test.FieldPayConfigs, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.PrintTimes(); ok {
		_spec.SetField(test.FieldPrintTimes, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedPrintTimes(); ok {
		_spec.AddField(test.FieldPrintTimes, field.TypeInt, value)
	}
	if tu.mutation.PrintTimesCleared() {
		_spec.ClearField(test.FieldPrintTimes, field.TypeInt)
	}
	if value, ok := tu.mutation.RingConfigs(); ok {
		_spec.SetField(test.FieldRingConfigs, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.BusinessGroupID(); ok {
		_spec.SetField(test.FieldBusinessGroupID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedBusinessGroupID(); ok {
		_spec.AddField(test.FieldBusinessGroupID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.BusinessGroupUID(); ok {
		_spec.SetField(test.FieldBusinessGroupUID, field.TypeString, value)
	}
	if value, ok := tu.mutation.DishCategories(); ok {
		_spec.SetField(test.FieldDishCategories, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedDishCategories(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, test.FieldDishCategories, value)
		})
	}
	if value, ok := tu.mutation.EndOfTakeaway(); ok {
		_spec.SetField(test.FieldEndOfTakeaway, field.TypeTime, value)
	}
	if tu.mutation.EndOfTakeawayCleared() {
		_spec.ClearField(test.FieldEndOfTakeaway, field.TypeTime)
	}
	if value, ok := tu.mutation.Mode(); ok {
		_spec.SetField(test.FieldMode, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedMode(); ok {
		_spec.AddField(test.FieldMode, field.TypeInt, value)
	}
	if value, ok := tu.mutation.SelfPickupConfigs(); ok {
		_spec.SetField(test.FieldSelfPickupConfigs, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.ServerID(); ok {
		_spec.SetField(test.FieldServerID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedServerID(); ok {
		_spec.AddField(test.FieldServerID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Image(); ok {
		_spec.SetField(test.FieldImage, field.TypeString, value)
	}
	if value, ok := tu.mutation.Address(); ok {
		_spec.SetField(test.FieldAddress, field.TypeString, value)
	}
	if value, ok := tu.mutation.Latitude(); ok {
		_spec.SetField(test.FieldLatitude, field.TypeString, value)
	}
	if value, ok := tu.mutation.Longitude(); ok {
		_spec.SetField(test.FieldLongitude, field.TypeString, value)
	}
	if value, ok := tu.mutation.Banners(); ok {
		_spec.SetField(test.FieldBanners, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.Sort(); ok {
		_spec.SetField(test.FieldSort, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedSort(); ok {
		_spec.AddField(test.FieldSort, field.TypeInt, value)
	}
	if value, ok := tu.mutation.PayMode(); ok {
		_spec.SetField(test.FieldPayMode, field.TypeString, value)
	}
	if value, ok := tu.mutation.DineInConfigs(); ok {
		_spec.SetField(test.FieldDineInConfigs, field.TypeJSON, value)
	}
	if tu.mutation.DineInConfigsCleared() {
		_spec.ClearField(test.FieldDineInConfigs, field.TypeJSON)
	}
	if value, ok := tu.mutation.DealsConfig(); ok {
		_spec.SetField(test.FieldDealsConfig, field.TypeJSON, value)
	}
	if tu.mutation.DealsConfigCleared() {
		_spec.ClearField(test.FieldDealsConfig, field.TypeJSON)
	}
	if value, ok := tu.mutation.PrintConfig(); ok {
		_spec.SetField(test.FieldPrintConfig, field.TypeJSON, value)
	}
	if tu.mutation.PrintConfigCleared() {
		_spec.ClearField(test.FieldPrintConfig, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{test.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TestUpdateOne is the builder for updating a single Test entity.
type TestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TestUpdateOne) SetUpdatedAt(t time.Time) *TestUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (tuo *TestUpdateOne) SetDeactivatedAt(t time.Time) *TestUpdateOne {
	tuo.mutation.SetDeactivatedAt(t)
	return tuo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableDeactivatedAt(t *time.Time) *TestUpdateOne {
	if t != nil {
		tuo.SetDeactivatedAt(*t)
	}
	return tuo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (tuo *TestUpdateOne) ClearDeactivatedAt() *TestUpdateOne {
	tuo.mutation.ClearDeactivatedAt()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TestUpdateOne) SetName(s string) *TestUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetEnableTakeaway sets the "enable_takeaway" field.
func (tuo *TestUpdateOne) SetEnableTakeaway(b bool) *TestUpdateOne {
	tuo.mutation.SetEnableTakeaway(b)
	return tuo
}

// SetNillableEnableTakeaway sets the "enable_takeaway" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableEnableTakeaway(b *bool) *TestUpdateOne {
	if b != nil {
		tuo.SetEnableTakeaway(*b)
	}
	return tuo
}

// SetPhone sets the "phone" field.
func (tuo *TestUpdateOne) SetPhone(s string) *TestUpdateOne {
	tuo.mutation.SetPhone(s)
	return tuo
}

// SetAnnouncement sets the "announcement" field.
func (tuo *TestUpdateOne) SetAnnouncement(s string) *TestUpdateOne {
	tuo.mutation.SetAnnouncement(s)
	return tuo
}

// SetEnableAutoAccept sets the "enable_auto_accept" field.
func (tuo *TestUpdateOne) SetEnableAutoAccept(b bool) *TestUpdateOne {
	tuo.mutation.SetEnableAutoAccept(b)
	return tuo
}

// SetNillableEnableAutoAccept sets the "enable_auto_accept" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableEnableAutoAccept(b *bool) *TestUpdateOne {
	if b != nil {
		tuo.SetEnableAutoAccept(*b)
	}
	return tuo
}

// SetDeliveryConfigs sets the "delivery_configs" field.
func (tuo *TestUpdateOne) SetDeliveryConfigs(sc schema.DeliveryConfig) *TestUpdateOne {
	tuo.mutation.SetDeliveryConfigs(sc)
	return tuo
}

// SetPayConfigs sets the "pay_configs" field.
func (tuo *TestUpdateOne) SetPayConfigs(sc schema.PayConfig) *TestUpdateOne {
	tuo.mutation.SetPayConfigs(sc)
	return tuo
}

// SetPrintTimes sets the "print_times" field.
func (tuo *TestUpdateOne) SetPrintTimes(i int) *TestUpdateOne {
	tuo.mutation.ResetPrintTimes()
	tuo.mutation.SetPrintTimes(i)
	return tuo
}

// SetNillablePrintTimes sets the "print_times" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillablePrintTimes(i *int) *TestUpdateOne {
	if i != nil {
		tuo.SetPrintTimes(*i)
	}
	return tuo
}

// AddPrintTimes adds i to the "print_times" field.
func (tuo *TestUpdateOne) AddPrintTimes(i int) *TestUpdateOne {
	tuo.mutation.AddPrintTimes(i)
	return tuo
}

// ClearPrintTimes clears the value of the "print_times" field.
func (tuo *TestUpdateOne) ClearPrintTimes() *TestUpdateOne {
	tuo.mutation.ClearPrintTimes()
	return tuo
}

// SetRingConfigs sets the "ring_configs" field.
func (tuo *TestUpdateOne) SetRingConfigs(sc schema.RingConfig) *TestUpdateOne {
	tuo.mutation.SetRingConfigs(sc)
	return tuo
}

// SetBusinessGroupID sets the "business_group_id" field.
func (tuo *TestUpdateOne) SetBusinessGroupID(i int) *TestUpdateOne {
	tuo.mutation.ResetBusinessGroupID()
	tuo.mutation.SetBusinessGroupID(i)
	return tuo
}

// AddBusinessGroupID adds i to the "business_group_id" field.
func (tuo *TestUpdateOne) AddBusinessGroupID(i int) *TestUpdateOne {
	tuo.mutation.AddBusinessGroupID(i)
	return tuo
}

// SetBusinessGroupUID sets the "business_group_uid" field.
func (tuo *TestUpdateOne) SetBusinessGroupUID(s string) *TestUpdateOne {
	tuo.mutation.SetBusinessGroupUID(s)
	return tuo
}

// SetDishCategories sets the "dish_categories" field.
func (tuo *TestUpdateOne) SetDishCategories(s []string) *TestUpdateOne {
	tuo.mutation.SetDishCategories(s)
	return tuo
}

// AppendDishCategories appends s to the "dish_categories" field.
func (tuo *TestUpdateOne) AppendDishCategories(s []string) *TestUpdateOne {
	tuo.mutation.AppendDishCategories(s)
	return tuo
}

// SetEndOfTakeaway sets the "end_of_takeaway" field.
func (tuo *TestUpdateOne) SetEndOfTakeaway(t time.Time) *TestUpdateOne {
	tuo.mutation.SetEndOfTakeaway(t)
	return tuo
}

// SetNillableEndOfTakeaway sets the "end_of_takeaway" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableEndOfTakeaway(t *time.Time) *TestUpdateOne {
	if t != nil {
		tuo.SetEndOfTakeaway(*t)
	}
	return tuo
}

// ClearEndOfTakeaway clears the value of the "end_of_takeaway" field.
func (tuo *TestUpdateOne) ClearEndOfTakeaway() *TestUpdateOne {
	tuo.mutation.ClearEndOfTakeaway()
	return tuo
}

// SetMode sets the "mode" field.
func (tuo *TestUpdateOne) SetMode(i int) *TestUpdateOne {
	tuo.mutation.ResetMode()
	tuo.mutation.SetMode(i)
	return tuo
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableMode(i *int) *TestUpdateOne {
	if i != nil {
		tuo.SetMode(*i)
	}
	return tuo
}

// AddMode adds i to the "mode" field.
func (tuo *TestUpdateOne) AddMode(i int) *TestUpdateOne {
	tuo.mutation.AddMode(i)
	return tuo
}

// SetSelfPickupConfigs sets the "self_pickup_configs" field.
func (tuo *TestUpdateOne) SetSelfPickupConfigs(spc schema.SelfPickupConfig) *TestUpdateOne {
	tuo.mutation.SetSelfPickupConfigs(spc)
	return tuo
}

// SetServerID sets the "server_id" field.
func (tuo *TestUpdateOne) SetServerID(i int) *TestUpdateOne {
	tuo.mutation.ResetServerID()
	tuo.mutation.SetServerID(i)
	return tuo
}

// AddServerID adds i to the "server_id" field.
func (tuo *TestUpdateOne) AddServerID(i int) *TestUpdateOne {
	tuo.mutation.AddServerID(i)
	return tuo
}

// SetImage sets the "image" field.
func (tuo *TestUpdateOne) SetImage(s string) *TestUpdateOne {
	tuo.mutation.SetImage(s)
	return tuo
}

// SetAddress sets the "address" field.
func (tuo *TestUpdateOne) SetAddress(s string) *TestUpdateOne {
	tuo.mutation.SetAddress(s)
	return tuo
}

// SetLatitude sets the "latitude" field.
func (tuo *TestUpdateOne) SetLatitude(s string) *TestUpdateOne {
	tuo.mutation.SetLatitude(s)
	return tuo
}

// SetLongitude sets the "longitude" field.
func (tuo *TestUpdateOne) SetLongitude(s string) *TestUpdateOne {
	tuo.mutation.SetLongitude(s)
	return tuo
}

// SetBanners sets the "banners" field.
func (tuo *TestUpdateOne) SetBanners(s schema.Banners) *TestUpdateOne {
	tuo.mutation.SetBanners(s)
	return tuo
}

// SetSort sets the "sort" field.
func (tuo *TestUpdateOne) SetSort(i int) *TestUpdateOne {
	tuo.mutation.ResetSort()
	tuo.mutation.SetSort(i)
	return tuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableSort(i *int) *TestUpdateOne {
	if i != nil {
		tuo.SetSort(*i)
	}
	return tuo
}

// AddSort adds i to the "sort" field.
func (tuo *TestUpdateOne) AddSort(i int) *TestUpdateOne {
	tuo.mutation.AddSort(i)
	return tuo
}

// SetPayMode sets the "pay_mode" field.
func (tuo *TestUpdateOne) SetPayMode(s string) *TestUpdateOne {
	tuo.mutation.SetPayMode(s)
	return tuo
}

// SetNillablePayMode sets the "pay_mode" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillablePayMode(s *string) *TestUpdateOne {
	if s != nil {
		tuo.SetPayMode(*s)
	}
	return tuo
}

// SetDineInConfigs sets the "dine_in_configs" field.
func (tuo *TestUpdateOne) SetDineInConfigs(sic schema.DineInConfigs) *TestUpdateOne {
	tuo.mutation.SetDineInConfigs(sic)
	return tuo
}

// SetNillableDineInConfigs sets the "dine_in_configs" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableDineInConfigs(sic *schema.DineInConfigs) *TestUpdateOne {
	if sic != nil {
		tuo.SetDineInConfigs(*sic)
	}
	return tuo
}

// ClearDineInConfigs clears the value of the "dine_in_configs" field.
func (tuo *TestUpdateOne) ClearDineInConfigs() *TestUpdateOne {
	tuo.mutation.ClearDineInConfigs()
	return tuo
}

// SetDealsConfig sets the "deals_config" field.
func (tuo *TestUpdateOne) SetDealsConfig(sc schema.DealsConfig) *TestUpdateOne {
	tuo.mutation.SetDealsConfig(sc)
	return tuo
}

// SetNillableDealsConfig sets the "deals_config" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillableDealsConfig(sc *schema.DealsConfig) *TestUpdateOne {
	if sc != nil {
		tuo.SetDealsConfig(*sc)
	}
	return tuo
}

// ClearDealsConfig clears the value of the "deals_config" field.
func (tuo *TestUpdateOne) ClearDealsConfig() *TestUpdateOne {
	tuo.mutation.ClearDealsConfig()
	return tuo
}

// SetPrintConfig sets the "print_config" field.
func (tuo *TestUpdateOne) SetPrintConfig(sc schema.PrintConfig) *TestUpdateOne {
	tuo.mutation.SetPrintConfig(sc)
	return tuo
}

// SetNillablePrintConfig sets the "print_config" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillablePrintConfig(sc *schema.PrintConfig) *TestUpdateOne {
	if sc != nil {
		tuo.SetPrintConfig(*sc)
	}
	return tuo
}

// ClearPrintConfig clears the value of the "print_config" field.
func (tuo *TestUpdateOne) ClearPrintConfig() *TestUpdateOne {
	tuo.mutation.ClearPrintConfig()
	return tuo
}

// Mutation returns the TestMutation object of the builder.
func (tuo *TestUpdateOne) Mutation() *TestMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TestUpdateOne) Select(field string, fields ...string) *TestUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Test entity.
func (tuo *TestUpdateOne) Save(ctx context.Context) (*Test, error) {
	var (
		err  error
		node *Test
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Test)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TestUpdateOne) SaveX(ctx context.Context) *Test {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TestUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TestUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TestUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := test.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TestUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := test.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entschema: validator failed for field "Test.name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Phone(); ok {
		if err := test.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`entschema: validator failed for field "Test.phone": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Announcement(); ok {
		if err := test.AnnouncementValidator(v); err != nil {
			return &ValidationError{Name: "announcement", err: fmt.Errorf(`entschema: validator failed for field "Test.announcement": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Image(); ok {
		if err := test.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`entschema: validator failed for field "Test.image": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Address(); ok {
		if err := test.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`entschema: validator failed for field "Test.address": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Latitude(); ok {
		if err := test.LatitudeValidator(v); err != nil {
			return &ValidationError{Name: "latitude", err: fmt.Errorf(`entschema: validator failed for field "Test.latitude": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Longitude(); ok {
		if err := test.LongitudeValidator(v); err != nil {
			return &ValidationError{Name: "longitude", err: fmt.Errorf(`entschema: validator failed for field "Test.longitude": %w`, err)}
		}
	}
	return nil
}

func (tuo *TestUpdateOne) sqlSave(ctx context.Context) (_node *Test, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   test.Table,
			Columns: test.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: test.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "Test.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, test.FieldID)
		for _, f := range fields {
			if !test.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != test.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(test.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.DeactivatedAt(); ok {
		_spec.SetField(test.FieldDeactivatedAt, field.TypeTime, value)
	}
	if tuo.mutation.DeactivatedAtCleared() {
		_spec.ClearField(test.FieldDeactivatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(test.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.EnableTakeaway(); ok {
		_spec.SetField(test.FieldEnableTakeaway, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Phone(); ok {
		_spec.SetField(test.FieldPhone, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Announcement(); ok {
		_spec.SetField(test.FieldAnnouncement, field.TypeString, value)
	}
	if value, ok := tuo.mutation.EnableAutoAccept(); ok {
		_spec.SetField(test.FieldEnableAutoAccept, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.DeliveryConfigs(); ok {
		_spec.SetField(test.FieldDeliveryConfigs, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.PayConfigs(); ok {
		_spec.SetField(test.FieldPayConfigs, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.PrintTimes(); ok {
		_spec.SetField(test.FieldPrintTimes, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedPrintTimes(); ok {
		_spec.AddField(test.FieldPrintTimes, field.TypeInt, value)
	}
	if tuo.mutation.PrintTimesCleared() {
		_spec.ClearField(test.FieldPrintTimes, field.TypeInt)
	}
	if value, ok := tuo.mutation.RingConfigs(); ok {
		_spec.SetField(test.FieldRingConfigs, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.BusinessGroupID(); ok {
		_spec.SetField(test.FieldBusinessGroupID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedBusinessGroupID(); ok {
		_spec.AddField(test.FieldBusinessGroupID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.BusinessGroupUID(); ok {
		_spec.SetField(test.FieldBusinessGroupUID, field.TypeString, value)
	}
	if value, ok := tuo.mutation.DishCategories(); ok {
		_spec.SetField(test.FieldDishCategories, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedDishCategories(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, test.FieldDishCategories, value)
		})
	}
	if value, ok := tuo.mutation.EndOfTakeaway(); ok {
		_spec.SetField(test.FieldEndOfTakeaway, field.TypeTime, value)
	}
	if tuo.mutation.EndOfTakeawayCleared() {
		_spec.ClearField(test.FieldEndOfTakeaway, field.TypeTime)
	}
	if value, ok := tuo.mutation.Mode(); ok {
		_spec.SetField(test.FieldMode, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedMode(); ok {
		_spec.AddField(test.FieldMode, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.SelfPickupConfigs(); ok {
		_spec.SetField(test.FieldSelfPickupConfigs, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.ServerID(); ok {
		_spec.SetField(test.FieldServerID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedServerID(); ok {
		_spec.AddField(test.FieldServerID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Image(); ok {
		_spec.SetField(test.FieldImage, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Address(); ok {
		_spec.SetField(test.FieldAddress, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Latitude(); ok {
		_spec.SetField(test.FieldLatitude, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Longitude(); ok {
		_spec.SetField(test.FieldLongitude, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Banners(); ok {
		_spec.SetField(test.FieldBanners, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.Sort(); ok {
		_spec.SetField(test.FieldSort, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedSort(); ok {
		_spec.AddField(test.FieldSort, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.PayMode(); ok {
		_spec.SetField(test.FieldPayMode, field.TypeString, value)
	}
	if value, ok := tuo.mutation.DineInConfigs(); ok {
		_spec.SetField(test.FieldDineInConfigs, field.TypeJSON, value)
	}
	if tuo.mutation.DineInConfigsCleared() {
		_spec.ClearField(test.FieldDineInConfigs, field.TypeJSON)
	}
	if value, ok := tuo.mutation.DealsConfig(); ok {
		_spec.SetField(test.FieldDealsConfig, field.TypeJSON, value)
	}
	if tuo.mutation.DealsConfigCleared() {
		_spec.ClearField(test.FieldDealsConfig, field.TypeJSON)
	}
	if value, ok := tuo.mutation.PrintConfig(); ok {
		_spec.SetField(test.FieldPrintConfig, field.TypeJSON, value)
	}
	if tuo.mutation.PrintConfigCleared() {
		_spec.ClearField(test.FieldPrintConfig, field.TypeJSON)
	}
	_node = &Test{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{test.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
