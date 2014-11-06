/*
유저 데이터베이스를 관리한다.
입력받은 유저 정보를 JSON 파일로 저장하고
요청한 유저 정보를 JSON 파일에서 읽어서 반환한다.
*/
package database

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Database struct {
	File   string
	Record Vcard
}

type Vcard struct {
	Name    string
	Address string
	Blog    string
	Email   string
}

func (db *Database) Open(file string) *Database {
	db.File = file
	return db
}

func (db *Database) FindVcard(name string) (*Vcard, error) {
	info, err := ioutil.ReadFile(db.File + "/" + name)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(info))
	return &db.Record, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal Error : %s", err.Error())
		os.Exit(1)
	}
}
