package dao

import (
	"fmt"
	"remote-part-job-back/common/mysql"
	"sort"
	"strconv"
)

func CreateJobInfo(user *JobInfo) (uint, error) {
	db := mysql.Orm.Table(JobInfo{}.TableName()).Create(user)
	if db.Error != nil {
		err := fmt.Errorf("job_id=" + strconv.Itoa(int(user.Id)) + " create job info db error:" + db.Error.Error())
		//metrics.MysqlErrorInc(MysqlAppBucketInfo{}.TableName(), user.AppId, "update")
		return 0, err
	}
	//mysqlAppBucketInfoData.cache.Store(user.AppId, *user)
	return user.Id, nil
}

func UpdateJobInfo(user *JobInfo) error {
	db := mysql.Orm.Table(JobInfo{}.TableName()).Save(user)
	if db.Error != nil {
		err := fmt.Errorf("job_id=" + strconv.Itoa(int(user.Id)) + " update job info db error:" + db.Error.Error())
		//metrics.MysqlErrorInc(MysqlAppBucketInfo{}.TableName(), user.AppId, "update")
		return err
	}
	//mysqlAppBucketInfoData.cache.Store(user.AppId, *user)
	return nil
}

func GetJobInfoByJobId(jobId string) (*JobInfo, error) {
	var info JobInfo
	result := mysql.Orm.Table(JobInfo{}.TableName()).Where("id=?", jobId).First(&info)

	if result.Error != nil {
		return nil, result.Error
	}
	return &info, nil
}

func DeleteJobInfo(id string) error {
	db := mysql.Orm.Table(JobInfo{}.TableName()).Where("id = ?", id).Delete(&JobInfo{})
	if db.Error != nil {
		err := fmt.Errorf("job_id = " + id + " delete mysql job info db error:" + db.Error.Error())
		return err
	}
	//mysqlAppInfoData.cache.Delete(appid)
	return nil
}

func GetJobInfoList() (infoList JobInfoList, err error) {
	result := mysql.Orm.Table(JobInfo{}.TableName())
	result.Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	sort.Sort(infoList)
	return infoList, nil
}

type JobInfoList []JobInfo

func (list JobInfoList) Len() int {
	return len(list)
}

func (list JobInfoList) Less(i, j int) bool {
	if list[i].IsTop != list[j].IsTop {
		return list[i].IsTop
	}
	return list[i].CreatedAt > list[j].CreatedAt
}

func (list JobInfoList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
