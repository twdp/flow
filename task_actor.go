package goflow

//任务参与者
type TaskActor struct {
	Id      string `xorm:"varchar(36) pk notnull"`                        //主键ID
	TaskId  string `xorm:"varchar(36) notnull index(IDX_TASKACTOR_TASK)"` //任务ID
	ActorId string `xorm:"varchar(36) notnull"`                           //参与者ID
}
