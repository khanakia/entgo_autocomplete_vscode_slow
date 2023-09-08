package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("tx_type_id").Optional(),
		field.String("sku").Optional(),
		field.String("title").Optional(),
		field.String("sub_title").Optional(),
		field.Int64("status_id").Optional(),
		field.Bool("no_index").Optional(),
		field.Int64("item_category_id").Optional(),
		field.Bool("base_on_sale_price").Optional(),
		field.Float("retail_price").Optional(),
		field.Float("sale_price").Optional(),
		field.Int64("shipping_profile_id_num").Optional(),
		field.Int64("optimization_id_num").Optional(),
		field.Int64("qty").Optional(),
		field.Int64("maximum_order_qty").Optional(),
		field.Float("pkg_width").Optional(),
		field.Float("pkg_height").Optional(),
		field.Float("pkg_length").Optional(),
		field.Float("pkg_weight").Optional(),
		field.String("short_desc").Optional(),
		field.String("descr").Optional(),
		field.String("image_url1").Optional(),
		field.String("image_url2").Optional(),
		field.String("image_url3").Optional(),
		field.String("image_url4").Optional(),
		field.String("image_url5").Optional(),
		field.String("condition").Optional(),
		field.String("brand").Optional(),
		field.String("gender").Optional(),
		field.String("age_group").Optional(),
		field.String("color").Optional(),
		field.String("size").Optional(),
		field.Int64("uxm_item_id_num").Optional(),
		field.Bool("sent").Optional(),
		field.Time("sent_at").Optional(),
		field.Bool("has_error").Optional(),
		field.Int64("last_log_id").Optional(),
		field.Text("feed_response").Optional(),
		field.String("keywords").Optional(),
		field.Bool("sync").Optional(),
		field.String("user_id").Optional(),
		field.String("shipping_profile_id").Optional(),
		field.String("optimization_id").Optional(),
		field.String("uxm_item_id").Optional(),
		field.String("parent_id").Optional(),
		field.Text("item_type_id").Optional(),
		field.String("uid").Unique(),       // unique id auto generated
		field.Bool("is_parent").Optional(), // has childrens if true otherwise item has no childrens
	}
}
