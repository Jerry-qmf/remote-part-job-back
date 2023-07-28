package dao

import (
	"fmt"
	"remote-part-job-back/common/mysql"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
	jobMap  = &sync.Map{}
	jobList JobInfoList
)

func initJobInfo() {
	rev, err := getJobInfoList()
	if err != nil {
		panic("mysql init error: " + err.Error())
	}
	jobList = rev

	for _, job := range jobList {
		jobMap.Store(job.Id, job)
	}

	fmt.Println("mysql app info data init success!")
}

func GetJobInfoList() JobInfoList {
	return jobList
}

func CreateJobInfo(user *JobInfo) (uint, error) {
	db := mysql.Orm.Table(JobInfo{}.TableName()).Create(user)
	if db.Error != nil {
		err := fmt.Errorf("job_id=" + strconv.Itoa(int(user.Id)) + " create job info db error:" + db.Error.Error())
		//metrics.MysqlErrorInc(MysqlAppBucketInfo{}.TableName(), user.AppId, "update")
		return 0, err
	}
	jobMap.Store(user.Id, *user)
	jobList = append(jobList, *user)
	sort.Sort(jobList)
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
	jobMap.Store(user.Id, *user)
	for k, v := range jobList {
		if v.Id == user.Id {
			jobList[k] = *user
		}
	}
	sort.Sort(jobList)
	return nil
}

func GetJobInfoByJobId(jobId string) (*JobInfo, error) {
	id, _ := strconv.Atoi(jobId)
	info, ok := jobMap.Load(uint(id))
	if !ok {
		return nil, fmt.Errorf("没这个")
	}
	rev := info.(JobInfo)
	return &rev, nil
}

func DeleteJobInfo(id string) error {
	db := mysql.Orm.Table(JobInfo{}.TableName()).Where("id = ?", id).Delete(&JobInfo{})
	if db.Error != nil {
		err := fmt.Errorf("job_id = " + id + " delete mysql job info db error:" + db.Error.Error())
		return err
	}
	uerId, _ := strconv.Atoi(id)
	//mysqlAppInfoData.cache.Delete(appid)
	jobMap.Delete(uint(uerId))
	for k, v := range jobList {
		if v.Id == uint(uerId) {
			jobList = append(jobList[:k], jobList[k+1:]...)
			return nil
		}
	}
	return nil
}

func getJobInfoList() (infoList JobInfoList, err error) {
	result := mysql.Orm.Table(JobInfo{}.TableName())
	result.Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	sort.Sort(infoList)
	return infoList, nil
}

type JobInfoList []JobInfo

func (list JobInfoList) FilterKey(key string) JobInfoList {
	var filteredList JobInfoList

	for _, job := range list {
		if !strings.Contains(job.JobTitle, key) && !strings.Contains(job.JobLabel, key) {
			// 如果 jobTitle 和 jobLabel 均不包含 key，则不将其加入 filteredList
			continue
		}
		filteredList = append(filteredList, job)
	}
	return filteredList
}

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
