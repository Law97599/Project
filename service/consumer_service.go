package service

import (
	"Project/cache"
	"Project/common/util"
	"Project/model"
	"Project/vo"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"
)

func Listen() {
	for {
		val, err := cache.RedisClient.XRead(&redis.XReadArgs{
			Streams: []string{"BookCourseStream", "0"},
			Count:   1,
			Block:   0,
		}).Result()
		if err != nil {
			util.Log().Error("Listen XRead Error : %v\n", err)
			continue
		}
		if len(val) == 0 {
			continue
		}
		if len(val[0].Messages) == 0 {
			continue
		}
		bookCourseJson := val[0].Messages[0].Values["StudentCourseObj"]
		var bookCourseVo vo.BookCourseRequest
		_ = json.Unmarshal([]byte(bookCourseJson.(string)), &bookCourseVo)
		sid, _ := strconv.ParseInt(bookCourseVo.StudentID, 10, 64)
		cid, _ := strconv.ParseInt(bookCourseVo.CourseID, 10, 64)
		sc := model.StudentCourse{
			StudentID: sid,
			CourseID:  cid,
		}
		//先查如果有，就不插入
		count := int64(0)
		model.DB.Model(&model.StudentCourse{}).Where("STUDENT_ID = ? AND COURSE_ID = ?", sid, cid).Count(&count)
		if count > 0 {
			//删除消息
			cache.RedisClient.XDel("BookCourseStream", val[0].Messages[0].ID)
			continue
		}
		//双主键约束保证幂等性
		if err = model.DB.Create(&sc).Error; err != nil {
			util.Log().Error("Create StudentCourse Error : %v\n ", err)
		} else {
			//课程容量减1
			model.DB.Exec("UPDATE t_course SET COURSE_STOCK = COURSE_STOCK - 1 where COURSE_ID = ?", cid)
			//删除课表缓存
			cache.RedisClient.HDel("GetStudentCourse", bookCourseVo.StudentID)
			//删除消息
			cache.RedisClient.XDel("BookCourseStream", val[0].Messages[0].ID)
		}

	}
}
