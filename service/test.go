package service

import (
	"errors"
	"fmt"

	. "go-battle/model"

	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var testService *TestService

type TestService struct {
	db *gorm.DB
}

func TestServiceInit(db *gorm.DB) {
	testService = &TestService{db}
}

func GetTestService() *TestService {
	return testService
}

func (t *TestService) Ping() string {
	return "Pong"
}

func (t *TestService) TestName(name string) string {
	configName := viper.GetString("testName")
	return fmt.Sprintf("request name is %s , config name is %s", name, configName)
}

func (t *TestService) Create() *Test {
	test := &Test{Name: "", Age: 0}
	res := t.db.Create(test)
	//返回error,插入记录条数
	fmt.Printf("插入数据 error = %v  ; 插入条数 = %v", res.Error, res.RowsAffected)
	return test
}
func (t *TestService) UpInsert(name string, age int) (test *Test) {
	test = &Test{Name: name, Age: age}
	t.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, //id冲突时就更新列为新值
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(test)
	return
}

func (t *TestService) Last() *Test {
	var test *Test
	result := t.db.Last(test)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("last not found !!!")
		return nil
	}
	return test
}

func (t *TestService) Find(idList []int, name string) []Test {
	var testList []Test
	t.db.Select("id", "name", "age", "crated_at", "updated_at").
		Order("age desc, name").Limit(5).
		Where("id IN ? and name like ?", idList, name).Find(testList)
	return testList
}

func (t *TestService) FindByTable() *Test {
	var test *Test
	t.db.Table("test").Limit(1).Find(test)
	return test
}

func (t *TestService) QueryByRaw(name string) *Test {
	var test *Test
	t.db.Raw("select * from test where name like ?", name).Scan(test)
	return test
}
func (t *TestService) Save() bool {
	t.db.Save(&Test{})
	return true
}

func (t *TestService) Update(id int) int64 {
	res := t.db.Model(Test{}).Where("id=?", id).Updates(Test{Name: "update"})
	return res.RowsAffected
}

func (t *TestService) Delete(id int) int64 {
	res := t.db.Delete(&Test{}, id)
	return res.RowsAffected
}

func (t *TestService) Transaction() string {
	return "Pong"
}

func (t *TestService) Hook() string {
	return "Pong"
}
