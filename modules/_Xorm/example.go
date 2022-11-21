package sxorm

import "fmt"

// DbExample ...
func DbExample(db, db1 *DbInfo) {

	multiData := []User{}

	for i := 0; i < 5; i++ {

		// var updateInt int64 = 0

		switch i {

		case 0:
			// +-------------------------------------
			// | 	單筆新增測試
			// +-------------------------------------
			var insertCount int64 = 0
			for i := 0; i < 100; i++ {

				tempName := fmt.Sprintf("test%d", i)
				// random, _ := rand.Int(rand.Reader, big.NewInt(2)) // 亂數產生0~100
				insertInt, err := db.Engine.Insert(&User{Id: 111, Name: tempName, Sex: (i % 2) + 1})

				insertCount += insertInt

				if err != nil {
					panic(err)
				}
			}
			fmt.Println(fmt.Sprintf("單筆新增100次：%d", insertCount))

		case 1:
			// +-------------------------------------
			// | 	多筆新增測試
			// +-------------------------------------
			insertInt, err := db.Engine.Insert(&User{Name: "multi_id1", Sex: 1}, &User{Name: "multi_id2", Sex: 2}, &User{Name: "multi_id3", Sex: 3})
			if err != nil {
				panic(err)
			} else {
				fmt.Println(fmt.Sprintf("一次多筆(3)新增：%d", insertInt))
			}
		case 2:
			// +-------------------------------------
			// | 	Get 單筆查詢測試
			// +-------------------------------------

			// SELECT * FROM user LIMIT 1
			tempData := &User{}
			searchInt, err := db.Engine.Get(tempData)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf(fmt.Sprintf("單筆查詢1筆                 ：%t -> ", searchInt))
				fmt.Printf("%+v \n", tempData) //%#v is with type
			}

			// SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1
			// ASC 小 ~ 大	|	DESC 大 ~ 小
			tempData1 := &User{}
			searchInt, err = db.Engine.Where("name = ?", "multi_id1").Desc("id").Get(tempData1)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf(fmt.Sprintf("單筆查詢指定條件+排序       ：%t -> ", searchInt))
				fmt.Printf("%+v \n", tempData1) //%#v is with type
			}

			// SELECT id FROM user WHERE name = ? (multi_id1)
			var tempData2 int16
			searchInt, err = db.Engine.Table(&User{}).Where("name = ?", "multi_id1").Cols("id").Get(&tempData2)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf(fmt.Sprintf("單筆查詢透過cols指定查出欄位：%t -> ", searchInt))
				fmt.Printf("%d \n", tempData2) //%#v is with type
			}
			// 同 SELECT id FROM user WHERE name = ? (multi_id1) ，只是改成直接下 SQL
			tempData2 = 0
			searchInt, err = db.Engine.SQL("select `id` FROM `user` WHERE `name`='multi_id1'").Get(&tempData2)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf(fmt.Sprintf("同上單筆查詢，直接用SQL     ：%t -> ", searchInt))
				fmt.Printf("%d \n", tempData2) //%#v is with type
			}

		case 3:
			// +-------------------------------------
			// | 	Find 多筆讀取測試
			// +-------------------------------------
			// SELECT * FROM user WHERE name = ? AND age > 10 limit 10 offset 0
			tempData := []User{}
			err := db.Engine.Where("sex = ?", 1).Limit(100, 0).Find(&tempData)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("全部資料為:")
				fmt.Println(tempData) //%#v is with type
				multiData = tempData
			}
		case 4:
			insertInt, err := db1.Engine.Insert(&multiData)
			if err != nil {
				panic(err)
			} else {
				fmt.Println(fmt.Sprintf("SQL Server array新增：%d", insertInt))
			}

		case 5:
			// updateInt, _ = db.Engine.ID(2).Update(&_Xorm.User{Name: "test99", Sex: 2}) //指定ID更新

		case 6:
			// updateInt, _ = db.Engine.ID(2).Update(&_Xorm.User{Name: "test99", Sex: 2}) //指定ID更新

		case 7:
			// updateInt, _ = db.Engine.ID(2).Update(&_Xorm.User{Name: "test99", Sex: 2}) //指定ID更新
		}

		// fmt.Println(fmt.Sprintf("[%d]update:%d", i, updateInt))

	}

}
