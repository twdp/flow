package entity

import (
	"time"
)

//流程工作单实体类（一般称为流程实例）
type Order struct {
	Id             string    `xorm:"varchar(32) pk notnull"`                         //主键ID
	Version        int       `xorm:"tinyint"`                                        //版本
	ProcessId      string    `xorm:"varchar(32) notnull index(IDX_ORDER_PROCESSID)"` //流程定义ID
	Creator        string    `xorm:"varchar(100)"`                                   //流程实例创建者ID
	CreateTime     time.Time `xorm:"datetime notnull"`                               //流程实例创建时间
	ParentId       string    `xorm:"varchar(32) index(FK_ORDER_PARENTID)"`           //流程实例为子流程时，该字段标识父流程实例ID
	ParentNodeName string    `xorm:"varchar(100)"`                                   //流程实例为子流程时，该字段标识父流程哪个节点模型启动的子流程
	ExpireTime     time.Time `xorm:"datetime"`                                       //流程实例期望完成时间
	LastUpdateTime time.Time `xorm:"datetime"`                                       //流程实例上一次更新时间
	LastUpdator    string    `xorm:"varchar(100)"`                                   //流程实例上一次更新人员ID
	Priority       int       `xorm:"tinyint"`                                        //流程实例优先级
	OrderNo        string    `xorm:"varchar(100) index(IDX_ORDER_NO)"`               //流程实例编号
	Variable       string    `xorm:"varchar(2000)"`                                  //流程实例附属变量
}
