package sqlClient

import "server/msg"

// 插入数据
func RegisterInsetUser(user *msg.UserSt) (int, error) {

	query := `insert into user.sql()`
	result, err := Db.Exec(query)
	id, _ := result.LastInsertId()

	return int(id), err
}

// 数据
func QueryUser(user *msg.UserSt) ([]*msg.UserSt, error) {

	query := `select id from user.sql where email like ? or phone like ?`
	result, err := Db.Query(query, user.Email, user.Phone)

	var users []*msg.UserSt
	for result.Next() {
		var u msg.UserSt
		err := result.Scan(&u.Id)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)

	}
	return users, err
}
