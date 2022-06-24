// @Title dailyIncrease.go
// @description 统计日/月活跃新增文件数
package database

import (
	"SE/src/interface/admin/index"
	"fmt"
)

// DailyIncreaseRet 统计日月新增文件数的返回结果
type DailyIncreaseRet struct {
	Success bool                   //统计是否成功
	D_Data  []index.D_increaseData //日新增文件数数组。按照{Date,Num}组织
	M_Data  []index.M_increaseData //月新增文件数数组，按照{Month,Num}组织
	Msg     string                 //如果统计失败，填写封装后的错误信息
}

// @title GetIncrease
// @description 用于统计日/月新增用户
// @author 矫晓佳 ${DATE} ${TIME}
// @return _ DailyIncreaseRet “包含日/月新增文件数及日期数据的结构体”
func GetIncrease() DailyIncreaseRet {
	//从 Doc 表中获得日期数据
	rows, err := DB.Query("SELECT createDate from Doc order by createDate")
	if err != nil {
		fmt.Println(err)
		return DailyIncreaseRet{Success: false, Msg: "database err"}
	}
	defer rows.Close()

	//构造存储结果的中间变量
	var D_data []index.D_increaseData
	var M_data []index.M_increaseData
	//对每个数据遍历操作
	for rows.Next() {
		var D_date, M_date string
		rows.Scan(&D_date)
		//获取日期、月份
		D_date = D_date[0:10]
		M_date = D_date[0:7]

		//对 每日增长结构体数组 进行增长
		if D_data == nil {
			//如果为空，则创建第一个结构体
			//构造数组增长的辅助变量
			var date index.D_increaseData
			date.Date = D_date
			date.Num = 1
			D_data = append(D_data, date)
		} else {
			//初始化不存在标志。
			D_nonExist := false
			//遍历结果数组
			for i, D_element := range D_data {
				//如果日期已经存在于数组中，对对应的结构体的数量进行增长
				if D_element.Date == D_date {
					// D_element.Num += 1
					D_data[i].Num += 1
					break
				}
				if i == len(D_data)-1 {
					D_nonExist = true
				}
			}
			//如果不存在数组中，则添加
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
