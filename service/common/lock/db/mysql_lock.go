package lock

import (
	"github.com/robfig/cron"
	"time"
	"xhyl-micro/service/common/db"
)

/*
   mysql 方式实现分布式锁
*/

func init() {
	clearTimeoutLockTask()
}

type MethodLock struct {
	Id         int64     `gorm:"primary_key"`
	MethodName string    `gorm:"type:varchar(20);not null;column:method_name;"` //添加唯一索引
	Desc       string    `gorm:"type:varchar(256);"`
	UpdateTime time.Time `gorm:"type:varchar(256);column:update_time;"`
}

//获取分布式锁（向mysql插入一条数据，如果插入成功则表示获取锁成功）
func SetMysqlLock(method string, desc ...string) {
	mysqlLock := MethodLock{
		MethodName: method,
	}
	for {
		mysqlLock.UpdateTime = time.Now().UTC()
		lockFlag := insertLock(mysqlLock)
		if lockFlag {
			break
		}
	}
}

//释放分布式锁
func ReleaseMysqlLock(method string) {
	deleteLock(method)
}

func insertLock(lock MethodLock) bool {
	mysqlDb := db.NewMysqlDb()
	defer mysqlDb.Close()

	//if err := mysqlDb.Create(&lock).Error; err != nil {
	//	return false
	//}
	//
	//return true

	return mysqlDb.NewRecord(&lock)
}

func deleteLock(method string) {
	mysqlDb := db.NewMysqlDb()
	defer mysqlDb.Close()

	mysqlDb.Delete(MethodLock{}, "method_name = ?", method)
}

//定时清除超时的锁
func clearTimeoutLockTask() {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		mysqlDb := db.NewMysqlDb()
		defer mysqlDb.Close()

		mysqlDb.Delete(MethodLock{}, "TIMESTAMPDIFF(SECOND,update_time,NOW()) > 5")
		//fmt.Println("cron err: ",err)
	})
	c.Start()
}
