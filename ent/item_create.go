// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khanakia/entautoslow/ent/item"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetTxTypeID sets the "tx_type_id" field.
func (ic *ItemCreate) SetTxTypeID(i int64) *ItemCreate {
	ic.mutation.SetTxTypeID(i)
	return ic
}

// SetNillableTxTypeID sets the "tx_type_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableTxTypeID(i *int64) *ItemCreate {
	if i != nil {
		ic.SetTxTypeID(*i)
	}
	return ic
}

// SetSku sets the "sku" field.
func (ic *ItemCreate) SetSku(s string) *ItemCreate {
	ic.mutation.SetSku(s)
	return ic
}

// SetNillableSku sets the "sku" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSku(s *string) *ItemCreate {
	if s != nil {
		ic.SetSku(*s)
	}
	return ic
}

// SetTitle sets the "title" field.
func (ic *ItemCreate) SetTitle(s string) *ItemCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ic *ItemCreate) SetNillableTitle(s *string) *ItemCreate {
	if s != nil {
		ic.SetTitle(*s)
	}
	return ic
}

// SetSubTitle sets the "sub_title" field.
func (ic *ItemCreate) SetSubTitle(s string) *ItemCreate {
	ic.mutation.SetSubTitle(s)
	return ic
}

// SetNillableSubTitle sets the "sub_title" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSubTitle(s *string) *ItemCreate {
	if s != nil {
		ic.SetSubTitle(*s)
	}
	return ic
}

// SetStatusID sets the "status_id" field.
func (ic *ItemCreate) SetStatusID(i int64) *ItemCreate {
	ic.mutation.SetStatusID(i)
	return ic
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableStatusID(i *int64) *ItemCreate {
	if i != nil {
		ic.SetStatusID(*i)
	}
	return ic
}

// SetNoIndex sets the "no_index" field.
func (ic *ItemCreate) SetNoIndex(b bool) *ItemCreate {
	ic.mutation.SetNoIndex(b)
	return ic
}

// SetNillableNoIndex sets the "no_index" field if the given value is not nil.
func (ic *ItemCreate) SetNillableNoIndex(b *bool) *ItemCreate {
	if b != nil {
		ic.SetNoIndex(*b)
	}
	return ic
}

// SetItemCategoryID sets the "item_category_id" field.
func (ic *ItemCreate) SetItemCategoryID(i int64) *ItemCreate {
	ic.mutation.SetItemCategoryID(i)
	return ic
}

// SetNillableItemCategoryID sets the "item_category_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableItemCategoryID(i *int64) *ItemCreate {
	if i != nil {
		ic.SetItemCategoryID(*i)
	}
	return ic
}

// SetBaseOnSalePrice sets the "base_on_sale_price" field.
func (ic *ItemCreate) SetBaseOnSalePrice(b bool) *ItemCreate {
	ic.mutation.SetBaseOnSalePrice(b)
	return ic
}

// SetNillableBaseOnSalePrice sets the "base_on_sale_price" field if the given value is not nil.
func (ic *ItemCreate) SetNillableBaseOnSalePrice(b *bool) *ItemCreate {
	if b != nil {
		ic.SetBaseOnSalePrice(*b)
	}
	return ic
}

// SetRetailPrice sets the "retail_price" field.
func (ic *ItemCreate) SetRetailPrice(f float64) *ItemCreate {
	ic.mutation.SetRetailPrice(f)
	return ic
}

// SetNillableRetailPrice sets the "retail_price" field if the given value is not nil.
func (ic *ItemCreate) SetNillableRetailPrice(f *float64) *ItemCreate {
	if f != nil {
		ic.SetRetailPrice(*f)
	}
	return ic
}

// SetSalePrice sets the "sale_price" field.
func (ic *ItemCreate) SetSalePrice(f float64) *ItemCreate {
	ic.mutation.SetSalePrice(f)
	return ic
}

// SetNillableSalePrice sets the "sale_price" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSalePrice(f *float64) *ItemCreate {
	if f != nil {
		ic.SetSalePrice(*f)
	}
	return ic
}

// SetShippingProfileIDNum sets the "shipping_profile_id_num" field.
func (ic *ItemCreate) SetShippingProfileIDNum(i int64) *ItemCreate {
	ic.mutation.SetShippingProfileIDNum(i)
	return ic
}

// SetNillableShippingProfileIDNum sets the "shipping_profile_id_num" field if the given value is not nil.
func (ic *ItemCreate) SetNillableShippingProfileIDNum(i *int64) *ItemCreate {
	if i != nil {
		ic.SetShippingProfileIDNum(*i)
	}
	return ic
}

// SetOptimizationIDNum sets the "optimization_id_num" field.
func (ic *ItemCreate) SetOptimizationIDNum(i int64) *ItemCreate {
	ic.mutation.SetOptimizationIDNum(i)
	return ic
}

// SetNillableOptimizationIDNum sets the "optimization_id_num" field if the given value is not nil.
func (ic *ItemCreate) SetNillableOptimizationIDNum(i *int64) *ItemCreate {
	if i != nil {
		ic.SetOptimizationIDNum(*i)
	}
	return ic
}

// SetQty sets the "qty" field.
func (ic *ItemCreate) SetQty(i int64) *ItemCreate {
	ic.mutation.SetQty(i)
	return ic
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (ic *ItemCreate) SetNillableQty(i *int64) *ItemCreate {
	if i != nil {
		ic.SetQty(*i)
	}
	return ic
}

// SetMaximumOrderQty sets the "maximum_order_qty" field.
func (ic *ItemCreate) SetMaximumOrderQty(i int64) *ItemCreate {
	ic.mutation.SetMaximumOrderQty(i)
	return ic
}

// SetNillableMaximumOrderQty sets the "maximum_order_qty" field if the given value is not nil.
func (ic *ItemCreate) SetNillableMaximumOrderQty(i *int64) *ItemCreate {
	if i != nil {
		ic.SetMaximumOrderQty(*i)
	}
	return ic
}

// SetPkgWidth sets the "pkg_width" field.
func (ic *ItemCreate) SetPkgWidth(f float64) *ItemCreate {
	ic.mutation.SetPkgWidth(f)
	return ic
}

// SetNillablePkgWidth sets the "pkg_width" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePkgWidth(f *float64) *ItemCreate {
	if f != nil {
		ic.SetPkgWidth(*f)
	}
	return ic
}

// SetPkgHeight sets the "pkg_height" field.
func (ic *ItemCreate) SetPkgHeight(f float64) *ItemCreate {
	ic.mutation.SetPkgHeight(f)
	return ic
}

// SetNillablePkgHeight sets the "pkg_height" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePkgHeight(f *float64) *ItemCreate {
	if f != nil {
		ic.SetPkgHeight(*f)
	}
	return ic
}

// SetPkgLength sets the "pkg_length" field.
func (ic *ItemCreate) SetPkgLength(f float64) *ItemCreate {
	ic.mutation.SetPkgLength(f)
	return ic
}

// SetNillablePkgLength sets the "pkg_length" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePkgLength(f *float64) *ItemCreate {
	if f != nil {
		ic.SetPkgLength(*f)
	}
	return ic
}

// SetPkgWeight sets the "pkg_weight" field.
func (ic *ItemCreate) SetPkgWeight(f float64) *ItemCreate {
	ic.mutation.SetPkgWeight(f)
	return ic
}

// SetNillablePkgWeight sets the "pkg_weight" field if the given value is not nil.
func (ic *ItemCreate) SetNillablePkgWeight(f *float64) *ItemCreate {
	if f != nil {
		ic.SetPkgWeight(*f)
	}
	return ic
}

// SetShortDesc sets the "short_desc" field.
func (ic *ItemCreate) SetShortDesc(s string) *ItemCreate {
	ic.mutation.SetShortDesc(s)
	return ic
}

// SetNillableShortDesc sets the "short_desc" field if the given value is not nil.
func (ic *ItemCreate) SetNillableShortDesc(s *string) *ItemCreate {
	if s != nil {
		ic.SetShortDesc(*s)
	}
	return ic
}

// SetDescr sets the "descr" field.
func (ic *ItemCreate) SetDescr(s string) *ItemCreate {
	ic.mutation.SetDescr(s)
	return ic
}

// SetNillableDescr sets the "descr" field if the given value is not nil.
func (ic *ItemCreate) SetNillableDescr(s *string) *ItemCreate {
	if s != nil {
		ic.SetDescr(*s)
	}
	return ic
}

// SetImageUrl1 sets the "image_url1" field.
func (ic *ItemCreate) SetImageUrl1(s string) *ItemCreate {
	ic.mutation.SetImageUrl1(s)
	return ic
}

// SetNillableImageUrl1 sets the "image_url1" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImageUrl1(s *string) *ItemCreate {
	if s != nil {
		ic.SetImageUrl1(*s)
	}
	return ic
}

// SetImageUrl2 sets the "image_url2" field.
func (ic *ItemCreate) SetImageUrl2(s string) *ItemCreate {
	ic.mutation.SetImageUrl2(s)
	return ic
}

// SetNillableImageUrl2 sets the "image_url2" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImageUrl2(s *string) *ItemCreate {
	if s != nil {
		ic.SetImageUrl2(*s)
	}
	return ic
}

// SetImageUrl3 sets the "image_url3" field.
func (ic *ItemCreate) SetImageUrl3(s string) *ItemCreate {
	ic.mutation.SetImageUrl3(s)
	return ic
}

// SetNillableImageUrl3 sets the "image_url3" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImageUrl3(s *string) *ItemCreate {
	if s != nil {
		ic.SetImageUrl3(*s)
	}
	return ic
}

// SetImageUrl4 sets the "image_url4" field.
func (ic *ItemCreate) SetImageUrl4(s string) *ItemCreate {
	ic.mutation.SetImageUrl4(s)
	return ic
}

// SetNillableImageUrl4 sets the "image_url4" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImageUrl4(s *string) *ItemCreate {
	if s != nil {
		ic.SetImageUrl4(*s)
	}
	return ic
}

// SetImageUrl5 sets the "image_url5" field.
func (ic *ItemCreate) SetImageUrl5(s string) *ItemCreate {
	ic.mutation.SetImageUrl5(s)
	return ic
}

// SetNillableImageUrl5 sets the "image_url5" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImageUrl5(s *string) *ItemCreate {
	if s != nil {
		ic.SetImageUrl5(*s)
	}
	return ic
}

// SetCondition sets the "condition" field.
func (ic *ItemCreate) SetCondition(s string) *ItemCreate {
	ic.mutation.SetCondition(s)
	return ic
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCondition(s *string) *ItemCreate {
	if s != nil {
		ic.SetCondition(*s)
	}
	return ic
}

// SetBrand sets the "brand" field.
func (ic *ItemCreate) SetBrand(s string) *ItemCreate {
	ic.mutation.SetBrand(s)
	return ic
}

// SetNillableBrand sets the "brand" field if the given value is not nil.
func (ic *ItemCreate) SetNillableBrand(s *string) *ItemCreate {
	if s != nil {
		ic.SetBrand(*s)
	}
	return ic
}

// SetGender sets the "gender" field.
func (ic *ItemCreate) SetGender(s string) *ItemCreate {
	ic.mutation.SetGender(s)
	return ic
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (ic *ItemCreate) SetNillableGender(s *string) *ItemCreate {
	if s != nil {
		ic.SetGender(*s)
	}
	return ic
}

// SetAgeGroup sets the "age_group" field.
func (ic *ItemCreate) SetAgeGroup(s string) *ItemCreate {
	ic.mutation.SetAgeGroup(s)
	return ic
}

// SetNillableAgeGroup sets the "age_group" field if the given value is not nil.
func (ic *ItemCreate) SetNillableAgeGroup(s *string) *ItemCreate {
	if s != nil {
		ic.SetAgeGroup(*s)
	}
	return ic
}

// SetColor sets the "color" field.
func (ic *ItemCreate) SetColor(s string) *ItemCreate {
	ic.mutation.SetColor(s)
	return ic
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (ic *ItemCreate) SetNillableColor(s *string) *ItemCreate {
	if s != nil {
		ic.SetColor(*s)
	}
	return ic
}

// SetSize sets the "size" field.
func (ic *ItemCreate) SetSize(s string) *ItemCreate {
	ic.mutation.SetSize(s)
	return ic
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSize(s *string) *ItemCreate {
	if s != nil {
		ic.SetSize(*s)
	}
	return ic
}

// SetUxmItemIDNum sets the "uxm_item_id_num" field.
func (ic *ItemCreate) SetUxmItemIDNum(i int64) *ItemCreate {
	ic.mutation.SetUxmItemIDNum(i)
	return ic
}

// SetNillableUxmItemIDNum sets the "uxm_item_id_num" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUxmItemIDNum(i *int64) *ItemCreate {
	if i != nil {
		ic.SetUxmItemIDNum(*i)
	}
	return ic
}

// SetSent sets the "sent" field.
func (ic *ItemCreate) SetSent(b bool) *ItemCreate {
	ic.mutation.SetSent(b)
	return ic
}

// SetNillableSent sets the "sent" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSent(b *bool) *ItemCreate {
	if b != nil {
		ic.SetSent(*b)
	}
	return ic
}

// SetSentAt sets the "sent_at" field.
func (ic *ItemCreate) SetSentAt(t time.Time) *ItemCreate {
	ic.mutation.SetSentAt(t)
	return ic
}

// SetNillableSentAt sets the "sent_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSentAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetSentAt(*t)
	}
	return ic
}

// SetHasError sets the "has_error" field.
func (ic *ItemCreate) SetHasError(b bool) *ItemCreate {
	ic.mutation.SetHasError(b)
	return ic
}

// SetNillableHasError sets the "has_error" field if the given value is not nil.
func (ic *ItemCreate) SetNillableHasError(b *bool) *ItemCreate {
	if b != nil {
		ic.SetHasError(*b)
	}
	return ic
}

// SetLastLogID sets the "last_log_id" field.
func (ic *ItemCreate) SetLastLogID(i int64) *ItemCreate {
	ic.mutation.SetLastLogID(i)
	return ic
}

// SetNillableLastLogID sets the "last_log_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableLastLogID(i *int64) *ItemCreate {
	if i != nil {
		ic.SetLastLogID(*i)
	}
	return ic
}

// SetFeedResponse sets the "feed_response" field.
func (ic *ItemCreate) SetFeedResponse(s string) *ItemCreate {
	ic.mutation.SetFeedResponse(s)
	return ic
}

// SetNillableFeedResponse sets the "feed_response" field if the given value is not nil.
func (ic *ItemCreate) SetNillableFeedResponse(s *string) *ItemCreate {
	if s != nil {
		ic.SetFeedResponse(*s)
	}
	return ic
}

// SetKeywords sets the "keywords" field.
func (ic *ItemCreate) SetKeywords(s string) *ItemCreate {
	ic.mutation.SetKeywords(s)
	return ic
}

// SetNillableKeywords sets the "keywords" field if the given value is not nil.
func (ic *ItemCreate) SetNillableKeywords(s *string) *ItemCreate {
	if s != nil {
		ic.SetKeywords(*s)
	}
	return ic
}

// SetSync sets the "sync" field.
func (ic *ItemCreate) SetSync(b bool) *ItemCreate {
	ic.mutation.SetSync(b)
	return ic
}

// SetNillableSync sets the "sync" field if the given value is not nil.
func (ic *ItemCreate) SetNillableSync(b *bool) *ItemCreate {
	if b != nil {
		ic.SetSync(*b)
	}
	return ic
}

// SetUserID sets the "user_id" field.
func (ic *ItemCreate) SetUserID(s string) *ItemCreate {
	ic.mutation.SetUserID(s)
	return ic
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUserID(s *string) *ItemCreate {
	if s != nil {
		ic.SetUserID(*s)
	}
	return ic
}

// SetShippingProfileID sets the "shipping_profile_id" field.
func (ic *ItemCreate) SetShippingProfileID(s string) *ItemCreate {
	ic.mutation.SetShippingProfileID(s)
	return ic
}

// SetNillableShippingProfileID sets the "shipping_profile_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableShippingProfileID(s *string) *ItemCreate {
	if s != nil {
		ic.SetShippingProfileID(*s)
	}
	return ic
}

// SetOptimizationID sets the "optimization_id" field.
func (ic *ItemCreate) SetOptimizationID(s string) *ItemCreate {
	ic.mutation.SetOptimizationID(s)
	return ic
}

// SetNillableOptimizationID sets the "optimization_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableOptimizationID(s *string) *ItemCreate {
	if s != nil {
		ic.SetOptimizationID(*s)
	}
	return ic
}

// SetUxmItemID sets the "uxm_item_id" field.
func (ic *ItemCreate) SetUxmItemID(s string) *ItemCreate {
	ic.mutation.SetUxmItemID(s)
	return ic
}

// SetNillableUxmItemID sets the "uxm_item_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUxmItemID(s *string) *ItemCreate {
	if s != nil {
		ic.SetUxmItemID(*s)
	}
	return ic
}

// SetParentID sets the "parent_id" field.
func (ic *ItemCreate) SetParentID(s string) *ItemCreate {
	ic.mutation.SetParentID(s)
	return ic
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableParentID(s *string) *ItemCreate {
	if s != nil {
		ic.SetParentID(*s)
	}
	return ic
}

// SetItemTypeID sets the "item_type_id" field.
func (ic *ItemCreate) SetItemTypeID(s string) *ItemCreate {
	ic.mutation.SetItemTypeID(s)
	return ic
}

// SetNillableItemTypeID sets the "item_type_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableItemTypeID(s *string) *ItemCreate {
	if s != nil {
		ic.SetItemTypeID(*s)
	}
	return ic
}

// SetUID sets the "uid" field.
func (ic *ItemCreate) SetUID(s string) *ItemCreate {
	ic.mutation.SetUID(s)
	return ic
}

// SetIsParent sets the "is_parent" field.
func (ic *ItemCreate) SetIsParent(b bool) *ItemCreate {
	ic.mutation.SetIsParent(b)
	return ic
}

// SetNillableIsParent sets the "is_parent" field if the given value is not nil.
func (ic *ItemCreate) SetNillableIsParent(b *bool) *ItemCreate {
	if b != nil {
		ic.SetIsParent(*b)
	}
	return ic
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "Item.uid"`)}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(item.Table, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.TxTypeID(); ok {
		_spec.SetField(item.FieldTxTypeID, field.TypeInt64, value)
		_node.TxTypeID = value
	}
	if value, ok := ic.mutation.Sku(); ok {
		_spec.SetField(item.FieldSku, field.TypeString, value)
		_node.Sku = value
	}
	if value, ok := ic.mutation.Title(); ok {
		_spec.SetField(item.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ic.mutation.SubTitle(); ok {
		_spec.SetField(item.FieldSubTitle, field.TypeString, value)
		_node.SubTitle = value
	}
	if value, ok := ic.mutation.StatusID(); ok {
		_spec.SetField(item.FieldStatusID, field.TypeInt64, value)
		_node.StatusID = value
	}
	if value, ok := ic.mutation.NoIndex(); ok {
		_spec.SetField(item.FieldNoIndex, field.TypeBool, value)
		_node.NoIndex = value
	}
	if value, ok := ic.mutation.ItemCategoryID(); ok {
		_spec.SetField(item.FieldItemCategoryID, field.TypeInt64, value)
		_node.ItemCategoryID = value
	}
	if value, ok := ic.mutation.BaseOnSalePrice(); ok {
		_spec.SetField(item.FieldBaseOnSalePrice, field.TypeBool, value)
		_node.BaseOnSalePrice = value
	}
	if value, ok := ic.mutation.RetailPrice(); ok {
		_spec.SetField(item.FieldRetailPrice, field.TypeFloat64, value)
		_node.RetailPrice = value
	}
	if value, ok := ic.mutation.SalePrice(); ok {
		_spec.SetField(item.FieldSalePrice, field.TypeFloat64, value)
		_node.SalePrice = value
	}
	if value, ok := ic.mutation.ShippingProfileIDNum(); ok {
		_spec.SetField(item.FieldShippingProfileIDNum, field.TypeInt64, value)
		_node.ShippingProfileIDNum = value
	}
	if value, ok := ic.mutation.OptimizationIDNum(); ok {
		_spec.SetField(item.FieldOptimizationIDNum, field.TypeInt64, value)
		_node.OptimizationIDNum = value
	}
	if value, ok := ic.mutation.Qty(); ok {
		_spec.SetField(item.FieldQty, field.TypeInt64, value)
		_node.Qty = value
	}
	if value, ok := ic.mutation.MaximumOrderQty(); ok {
		_spec.SetField(item.FieldMaximumOrderQty, field.TypeInt64, value)
		_node.MaximumOrderQty = value
	}
	if value, ok := ic.mutation.PkgWidth(); ok {
		_spec.SetField(item.FieldPkgWidth, field.TypeFloat64, value)
		_node.PkgWidth = value
	}
	if value, ok := ic.mutation.PkgHeight(); ok {
		_spec.SetField(item.FieldPkgHeight, field.TypeFloat64, value)
		_node.PkgHeight = value
	}
	if value, ok := ic.mutation.PkgLength(); ok {
		_spec.SetField(item.FieldPkgLength, field.TypeFloat64, value)
		_node.PkgLength = value
	}
	if value, ok := ic.mutation.PkgWeight(); ok {
		_spec.SetField(item.FieldPkgWeight, field.TypeFloat64, value)
		_node.PkgWeight = value
	}
	if value, ok := ic.mutation.ShortDesc(); ok {
		_spec.SetField(item.FieldShortDesc, field.TypeString, value)
		_node.ShortDesc = value
	}
	if value, ok := ic.mutation.Descr(); ok {
		_spec.SetField(item.FieldDescr, field.TypeString, value)
		_node.Descr = value
	}
	if value, ok := ic.mutation.ImageUrl1(); ok {
		_spec.SetField(item.FieldImageUrl1, field.TypeString, value)
		_node.ImageUrl1 = value
	}
	if value, ok := ic.mutation.ImageUrl2(); ok {
		_spec.SetField(item.FieldImageUrl2, field.TypeString, value)
		_node.ImageUrl2 = value
	}
	if value, ok := ic.mutation.ImageUrl3(); ok {
		_spec.SetField(item.FieldImageUrl3, field.TypeString, value)
		_node.ImageUrl3 = value
	}
	if value, ok := ic.mutation.ImageUrl4(); ok {
		_spec.SetField(item.FieldImageUrl4, field.TypeString, value)
		_node.ImageUrl4 = value
	}
	if value, ok := ic.mutation.ImageUrl5(); ok {
		_spec.SetField(item.FieldImageUrl5, field.TypeString, value)
		_node.ImageUrl5 = value
	}
	if value, ok := ic.mutation.Condition(); ok {
		_spec.SetField(item.FieldCondition, field.TypeString, value)
		_node.Condition = value
	}
	if value, ok := ic.mutation.Brand(); ok {
		_spec.SetField(item.FieldBrand, field.TypeString, value)
		_node.Brand = value
	}
	if value, ok := ic.mutation.Gender(); ok {
		_spec.SetField(item.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	if value, ok := ic.mutation.AgeGroup(); ok {
		_spec.SetField(item.FieldAgeGroup, field.TypeString, value)
		_node.AgeGroup = value
	}
	if value, ok := ic.mutation.Color(); ok {
		_spec.SetField(item.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if value, ok := ic.mutation.Size(); ok {
		_spec.SetField(item.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := ic.mutation.UxmItemIDNum(); ok {
		_spec.SetField(item.FieldUxmItemIDNum, field.TypeInt64, value)
		_node.UxmItemIDNum = value
	}
	if value, ok := ic.mutation.Sent(); ok {
		_spec.SetField(item.FieldSent, field.TypeBool, value)
		_node.Sent = value
	}
	if value, ok := ic.mutation.SentAt(); ok {
		_spec.SetField(item.FieldSentAt, field.TypeTime, value)
		_node.SentAt = value
	}
	if value, ok := ic.mutation.HasError(); ok {
		_spec.SetField(item.FieldHasError, field.TypeBool, value)
		_node.HasError = value
	}
	if value, ok := ic.mutation.LastLogID(); ok {
		_spec.SetField(item.FieldLastLogID, field.TypeInt64, value)
		_node.LastLogID = value
	}
	if value, ok := ic.mutation.FeedResponse(); ok {
		_spec.SetField(item.FieldFeedResponse, field.TypeString, value)
		_node.FeedResponse = value
	}
	if value, ok := ic.mutation.Keywords(); ok {
		_spec.SetField(item.FieldKeywords, field.TypeString, value)
		_node.Keywords = value
	}
	if value, ok := ic.mutation.Sync(); ok {
		_spec.SetField(item.FieldSync, field.TypeBool, value)
		_node.Sync = value
	}
	if value, ok := ic.mutation.UserID(); ok {
		_spec.SetField(item.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := ic.mutation.ShippingProfileID(); ok {
		_spec.SetField(item.FieldShippingProfileID, field.TypeString, value)
		_node.ShippingProfileID = value
	}
	if value, ok := ic.mutation.OptimizationID(); ok {
		_spec.SetField(item.FieldOptimizationID, field.TypeString, value)
		_node.OptimizationID = value
	}
	if value, ok := ic.mutation.UxmItemID(); ok {
		_spec.SetField(item.FieldUxmItemID, field.TypeString, value)
		_node.UxmItemID = value
	}
	if value, ok := ic.mutation.ParentID(); ok {
		_spec.SetField(item.FieldParentID, field.TypeString, value)
		_node.ParentID = value
	}
	if value, ok := ic.mutation.ItemTypeID(); ok {
		_spec.SetField(item.FieldItemTypeID, field.TypeString, value)
		_node.ItemTypeID = value
	}
	if value, ok := ic.mutation.UID(); ok {
		_spec.SetField(item.FieldUID, field.TypeString, value)
		_node.UID = value
	}
	if value, ok := ic.mutation.IsParent(); ok {
		_spec.SetField(item.FieldIsParent, field.TypeBool, value)
		_node.IsParent = value
	}
	return _node, _spec
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	builders []*ItemCreate
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
