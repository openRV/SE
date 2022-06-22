package database

import (
	"SE/src/interface/admin/index"
	"fmt"
)

type DailyIncreaseRet struct {
	Success bool
	D_Data  []index.D_increaseData
	M_Data  []index.M_increaseData
	Msg     string
}

func GetIncrease() DailyIncreaseRet {
	rows, err := DB.Query("SELECT createDate from Doc")
	if err != nil {
		fmt.Println(err)
		return DailyIncreaseRet{Success: false, Msg: "database err"}
	}
	defer rows.Close()

	var D_data []index.D_increaseData
	var M_data []index.M_increaseData
	for rows.Next() {
		var D_date, M_date string
		rows.Scan(&D_date)
		D_date = D_date[0:10]
		M_date = D_date[0:7]

		//对 每日增长结构体数组 进行增长
		if D_data == nil {
			var date index.D_increaseData
			date.Date = D_date
			date.Num = 1
			// D_data[0].Date = D_date
			// D_data[0].Num = 1
			D_data = append(D_data, date)
		} else {
			D_nonExist := false
			for i, D_element := range D_data {
				if D_element.Date == D_date {
					// D_element.Num += 1
					D_data[i].Num += 1
					break
				}
				if i == len(D_data)-1 {
					D_nonExist = true
				}
			}
			if D_nonExist {
				tmp1 := index.D_increaseData{
					Date: D_date,
					Num:  1,
				}
				D_data = append(D_data, tmp1)
			}
		}

		//对 每月增长结构体数组进行增长
		if M_data == nil {
			var date index.M_increaseData
			date.Month = M_date
			date.Num = 1
			M_data = append(M_data, date)
		} else {
			M_nonExist := false
			for i, M_element := range M_data {
				if M_element.Month == M_date {
					// M_element.Num += 1
					M_data[i].Num += 1
					break
				}
				if i == len(M_data)-1 {
					M_nonExist = true
				}
			}
			if M_nonExist {
				tmp2 := index.M_increaseData{
					Month: M_date,
					Num:   1,
				}
				M_data = append(M_data, tmp2)
			}
		}
	}

	return DailyIncreaseRet{
		Success: true,
		D_Data:  D_data,
		M_Data:  M_data,
	}
}
