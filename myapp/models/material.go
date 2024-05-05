package models

import ( "github.com/astaxie/beego/orm"
)

type Material struct {
	MaterialID   int    `orm:"column(material_id);pk"`
	MaterialName string `orm:"column(material_name)"`
	MaterialSubtype string `orm:"column(material_subtype)"`
}

func init() {
	orm.RegisterModel(new(Material))
}

func GetAllMaterials() ([]*Material, error) {
	o := orm.NewOrm()
	materials := make([]*Material, 0)

	_, err := o.QueryTable("material").All(&materials)
	return materials, err
}
