package dbdata

import "testing"

func TestMysqlDeamonCode(t *testing.T) {
	db, err := MysqlDeamonCode()
	if err != nil {
		t.Errorf(err.Error())
	}
	query := `
	CREATE TABLE users (
	       id INT AUTO_INCREMENT,
	       username TEXT NOT NULL,
	       password TEXT NOT NULL,
	       created_at DATETIME,
	       PRIMARY KEY (id)
	   );`
	_, err = db.Exec(query)
}
