package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Cview struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(255)" description:"视图名称"`
	Ctype      *Ctype    `orm:"column(ctype);rel(fk)" description:"对应的设备类型"`
	Keys       string    `orm:"column(keys);size(1024)" description:"视图keys 有排序"`
	IsDefault  int8      `orm:"column(is_default)" description:"是否是默认试图"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);auto_now_add"`
	UdpateTime time.Time `orm:"column(udpate_time);type(datetime);null;auto_now"`
}

func (t *Cview) TableName() string {
	return "cview"
}

func init() {
	orm.RegisterModel(new(Cview))
}

// AddCview insert a new Cview into database and returns
// last inserted Id on success.
func AddCview(m *Cview) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCviewById retrieves Cview by Id. Returns error if
// Id doesn't exist
func GetCviewById(id int) (v *Cview, err error) {
	o := orm.NewOrm()
	v = &Cview{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCview retrieves all Cview matches certain condition. Returns empty list if
// no records exist
func GetAllCview(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cview))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Cview
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCview updates Cview by Id and returns error if
// the record to be updated doesn't exist
func UpdateCviewById(m *Cview) (err error) {
	o := orm.NewOrm()
	v := Cview{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCview deletes Cview by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCview(id int) (err error) {
	o := orm.NewOrm()
	v := Cview{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Cview{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
