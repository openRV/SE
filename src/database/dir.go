package database

import "fmt"

type Dir struct {
	Id         string
	Name       string
	Owner      string
	CreateDate string
	LastView   string
	Subdir     []Dir
}

type UserDirRet struct {
	Success bool
	Name    string
	Data    []Dir
	Msg     string
}

func UserDir(id string, root bool) UserDirRet {

	var result UserDirRet
	result.Success = true

	// fill in dir name

	stmt := `
				select dirName
				from dir
				where dirId = ?
			`
	err := DB.QueryRow(stmt, id).Scan(&result.Name)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}

	// find sub dir Id from table Tree
	stmt = `
				select subId
				from Tree
				where dirId = ? AND root = ? AND subType = ?
			`
	rows, err := DB.Query(stmt, id, root, "dir")
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	defer rows.Close()

	for rows.Next() {
		var dir Dir
		err = rows.Scan(&dir.Id)
		if err != nil {
			fmt.Println(err)
			return UserDirRet{Success: false, Msg: "database error"}
		}

		// fill in blank name, owner, createDate, lastView of result
		fillinDirInfo(dir)
		// fill in subdir
		fillinSubDir(dir)

		result.Data = append(result.Data, dir)
	}

	return result

}

func fillinSubDir(dir Dir) UserDirRet {
	stmt := `
				select subId
				from Tree
				where dirId = ? AND subType = ?
			`
	row, err := DB.Query(stmt, dir.Id)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}

	for row.Next() {
		var subDir Dir
		row.Scan(&subDir.Id)
		res := fillinSubDir(subDir)
		if !res.Success {
			return res
		}
		res = fillinDirInfo(subDir)
		if !res.Success {
			return res
		}

		dir.Subdir = append(dir.Subdir, subDir)
	}
	return UserDirRet{Success: true}

}

func fillinDirInfo(dir Dir) UserDirRet {
	stmt := `
				select dirName , owner , createDate , lastView
				from Dir
				where dirId = ?
			`
	err := DB.QueryRow(stmt, dir.Id).Scan(dir.Name, dir.Owner, dir.CreateDate, dir.LastView)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	return UserDirRet{Success: true}
}
