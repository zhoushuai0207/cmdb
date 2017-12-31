package models

type Coperation struct {
	Id_RENAME int    `orm:"column(id)"`
	Name      string `orm:"column(name);size(64)"`
	Tag       string `orm:"column(tag);size(255)"`
	Uri       string `orm:"column(uri);size(255);null"`
	Desc      string `orm:"column(desc);size(255);null"`
}
