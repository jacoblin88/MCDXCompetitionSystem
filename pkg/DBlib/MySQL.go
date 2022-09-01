package DBlib

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

//interface

type SqlDatabase interface {
	CheckStatus() error
	//Check if We can connect the database

	Open() error
	Close() error
}

type DataBaseOperator interface {
	InsertToFlagTable(TeamID []int, MachineID []int, flag []string, Round []int) error
	//Insert New flag collection to FlagTable
	//This will be used when the competition start only

	UpdateFlagTable(TeamID []int, MachineID []int, flag []string, Round []int) error
	//Update FlagTable with a new flag collection of all team when starting a new Round

	InsertToFlagRecord(TeamID []int, MachineID []int, flag []string, Round []int) error
	//Record the new flag collection in FlagRecord when starting a new Round

	InsertScoreRecord(TeamID []int, AtkScore []int, DefScore []int, Round []int) error
	//Record AtkScore and DefScore of every Team in ScoreRecord when every Round ends.

	AddToScoreTotal(TeamID []int, AtkScore []int, DefScore []int) error
	//When a Round end,the system will add AtkScore and DefScore of Each Team to ScoreTotal.

	InsertToLogTable(TeamID []int, MachineID []int, Event []string, Round []int, Turn []int) error
	//When a Round end,the system will record all events occured during the Round.

	InitScoreTotal(TeamNum int) error
	//Initialize ScoreTotal Table according to participated Team number.
	InitLogTable(TeamNum int) error
	//Initialize Log Table according to participated Team number.

}

type MySQL struct {
	Username string `yaml:"uname"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	//Dbtype   string `yaml:"dbtype"`
	db sql.DB
}

func NewMySQL(config_name string) *MySQL {
	data, _ := ioutil.ReadFile(string(config_name))

	p := new(MySQL)

	err := yaml.Unmarshal([]byte(data), &p)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//Set Up connection string

	conn_str := "%s:%s@tcp(%s:%s)/%s"
	sql_str := fmt.Sprintf(conn_str, p.Username, p.Password, p.Host, p.Port, string(p.Dbname))
	db_obj, err := sql.Open("mysql", sql_str)
	p.db = *db_obj

	if err != nil {
		panic(err.Error())
	}

	//p.SetUpDB()
	//fmt.Println(p)
	return p
}

func (mysql *MySQL) Open() error {
	conn_str := "%s:%s@tcp(%s:%s)/%s"
	sql_str := fmt.Sprintf(conn_str, mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Dbname)
	db_obj, err := sql.Open("mysql", sql_str)
	mysql.db = *db_obj
	return err
}

func (mysql *MySQL) Close() error {
	err := mysql.db.Close()
	return err
}

func (mysql *MySQL) CheckStatus() error {
	err := mysql.db.Ping()
	return err
}

//====================================flag=====================

func (mysql *MySQL) InsertToFlagRecord(TeamID []int, MachineID []int, flag []string, Round []int) error {
	return mysql.InsertFlag(TeamID, MachineID, flag, Round, "FlagRecord")
}

func (mysql *MySQL) InsertToFlagTable(TeamID []int, MachineID []int, flag []string, Round []int) error {
	return mysql.InsertFlag(TeamID, MachineID, flag, Round, "FlagTable")
}

func (mysql *MySQL) InsertFlag(TeamID []int, MachineID []int, flag []string, Round []int, table string) error {

	tx, _ := mysql.db.Begin()

	for i := 0; i < len(TeamID); i++ {

		_, err := tx.Exec("INSERT INTO "+table+"(TeamID,MachineID,flag,Round) VALUES(?,?,?,?)", (TeamID[i]), MachineID[i], flag[i], Round[i])

		if err != nil {
			tx.Rollback()
			log.Print(err)
			return err
		}
	}
	return tx.Commit()
}

func (mysql *MySQL) UpdateFlagTable(TeamID []int, MachineID []int, flag []string, Round []int) error {

	tx, _ := mysql.db.Begin()

	for i := 0; i < len(TeamID); i++ {
		_, err := tx.Exec("Update FlagTable set flag=?,Round=? where TeamID=? and MachineID=?", flag[i], Round[i], TeamID[i], MachineID[i])

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

//====================================flag=====================

//=========================start Score=======================
func (mysql *MySQL) InsertScoreRecord(TeamID []int, AtkScore []int, DefScore []int, Round []int) error {

	tx, _ := mysql.db.Begin()

	for i := 0; i < len(TeamID); i++ {

		_, err := tx.Exec("INSERT INTO ScoreRecord(TeamID,AtkScore,DefScore,Round) VALUES(?,?,?,?)", (TeamID[i]), AtkScore[i], DefScore[i], Round[i])

		if err != nil {
			tx.Rollback()
			log.Panic(err)
			return err
		}
	}
	return tx.Commit()
}

func (mysql *MySQL) InitScoreTotal(TeamNum int) error {
	tx, _ := mysql.db.Begin()

	for i := 0; i < TeamNum; i++ {
		_, err := tx.Exec("INSERT INTO ScoreTotal(TeamID,AtkScore,DefScore) VALUES(?,?,?)", i+1, 0, 0)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (mysql *MySQL) AddToScoreTotal(TeamID []int, AtkScore []int, DefScore []int) error {
	tx, _ := mysql.db.Begin()

	for i := 0; i < len(TeamID); i++ {
		_, err := tx.Exec("Update ScoreTotal set AtkScore=AtkScore+?,DefScore=DefScore+? where TeamID=?", AtkScore[i], DefScore[i], TeamID[i])

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

//=========================end of Score=======================

func (mysql *MySQL) InitLogTable(TeamNum int) error {
	init_stmt := "DROP TABLE IF EXISTS %s;"
	create_stmt := "CREATE TABLE %s( no int not null AUTO_INCREMENT,MachineID int not null, Event varchar(32) not null,Round int not null,Turn int not null,PRIMARY KEY(no));"
	tx, _ := mysql.db.Begin()

	for i := 0; i < TeamNum; i++ {
		tb_name := "Team" + strconv.Itoa(i+1) + "_Log"
		stmt := fmt.Sprintf(init_stmt, tb_name)
		stmt1 := fmt.Sprintf(create_stmt, tb_name)
		fmt.Println(init_stmt)

		_, err := tx.Exec(stmt)

		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec(stmt1)

		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (mysql *MySQL) InsertToLogTable(TeamID []int, MachineID []int, Event []string, Round []int, Turn []int) error {
	tx, _ := mysql.db.Begin()

	for i := 0; i < len(TeamID); i++ {
		tb_name := "Team" + strconv.Itoa(TeamID[i]) + "_Log"
		stmt := "INSERT INTO " + tb_name + "(MachineID,Event,Round,Turn) VALUES(?,?,?,?)"
		//fmt.Sprintf(stmt, tb_name)
		_, err := tx.Exec(stmt, MachineID[i], Event[i], Round[i], Turn[i])

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()

}

///######################### end MySQL Class
