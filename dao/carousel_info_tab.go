package dao

import (
	"fmt"
	"remote-part-job-back/common/mysql"
	"sort"
	"strconv"
)

func CreateCarouselInfo(user *CarouselInfo) (uint, error) {
	db := mysql.Orm.Table(CarouselInfo{}.TableName()).Create(user)
	if db.Error != nil {
		err := fmt.Errorf("create carousel info db error:" + db.Error.Error())
		//metrics.MysqlErrorInc(MysqlAppBucketInfo{}.TableName(), user.AppId, "update")
		return 0, err
	}
	//mysqlAppBucketInfoData.cache.Store(user.AppId, *user)
	return user.Id, nil
}

func UpdateCarouselInfo(user *CarouselInfo) error {
	db := mysql.Orm.Table(CarouselInfo{}.TableName()).Save(user)
	if db.Error != nil {
		err := fmt.Errorf("Carousel_id=" + strconv.Itoa(int(user.Id)) + " update Carousel info db error:" + db.Error.Error())
		//metrics.MysqlErrorInc(MysqlAppBucketInfo{}.TableName(), user.AppId, "update")
		return err
	}
	//mysqlAppBucketInfoData.cache.Store(user.AppId, *user)
	return nil
}

func DeleteCarouselInfo(id string) error {
	db := mysql.Orm.Table(CarouselInfo{}.TableName()).Where("id = ?", id).Delete(&CarouselInfo{})
	if db.Error != nil {
		err := fmt.Errorf("Carousel_id = " + id + " delete mysql Carousel info db error:" + db.Error.Error())
		return err
	}
	//mysqlAppInfoData.cache.Delete(appid)
	return nil
}

func GetCarouselInfoList() (infoList CarouselInfoList, err error) {
	result := mysql.Orm.Table(CarouselInfo{}.TableName())
	result.Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	sort.Sort(infoList)
	return infoList, nil
}

type CarouselInfoList []CarouselInfo

func (list CarouselInfoList) Len() int {
	return len(list)
}

func (list CarouselInfoList) Less(i, j int) bool {
	if list[i].IsTop != list[j].IsTop {
		return list[i].IsTop
	}
	return list[i].CreatedAt < list[j].CreatedAt
}

func (list CarouselInfoList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
