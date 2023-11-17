package pojo

import (
	"database/sql/driver"
	"fmt"
	"goStudy/go_x_project_online_test/common"
	"time"
)

type MyTime time.Time

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(common.DateLayout))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format(common.DateLayout), nil
}
