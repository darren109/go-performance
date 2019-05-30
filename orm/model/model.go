package model

import (
	"database/sql"
	"fmt"

	"github.com/sunmi-OS/gocore/gorm"
)

var MachineHarder = Machine{}

type Machine struct {
	Id        int64  `gorm:"column:id"`
	Level     string `gorm:"column:level"`
	Name      string `gorm:"column:name"`
	ParentId  string `gorm:"column:parent_id"`
	TreePath  string `gorm:"column:tree_path"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

func (m *Machine) TableName() string {
	return "params_level"
}

func (m *Machine) GetList() []Machine {

	MachineList := []Machine{}

	err := gorm.GetORM().Find(&MachineList).Error

	if err != nil {
		fmt.Println(err.Error())
	}
	return MachineList
}

func (m *Machine) GetList2() []Machine {

	rows, err := gorm.GetORM().Model(m).Select("id,level,name,parent_id,tree_path,created_at,updated_at").Rows()
	if err != nil {
		fmt.Println(err.Error())
	}

	return m.rowsToMachine(rows)
}

func (m *Machine) rowsToMachine(rows *sql.Rows) []Machine {
	MachineList := []Machine{}

	defer rows.Close()
	for rows.Next() {

		MC := Machine{}
		rows.Scan(&MC.Id, &MC.Level, &MC.Name, &MC.ParentId, &MC.TreePath, &MC.CreatedAt, &MC.UpdatedAt)
		MachineList = append(MachineList, MC)
	}

	return MachineList
}
