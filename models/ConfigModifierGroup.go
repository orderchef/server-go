
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigModifierGroup struct {
	Id int `db:"id" json:"id"`

	Name string `db:"name" json:"name"`
	Required bool `db:"choice_required" json:"choice_required"`
	Deleted bool `db:"deleted" json:"-"`
}

func init() {
	database.Mysql().AddTableWithName(ConfigModifierGroup{}, "config__modifier_group").SetKeys(true, "id")
}

func GetAllModifierGroups() ([]ConfigModifierGroup, error) {
	db := database.Mysql()

	var objs []ConfigModifierGroup
	if _, err := db.Select(&objs, "select * from config__modifier_group where deleted=0 order by name"); err != nil {
		return nil, err
	}

	return objs, nil
}

func (modifierGroup *ConfigModifierGroup) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&modifierGroup, "select * from config__modifier_group where id=?", modifierGroup.Id); err != nil {
		return err
	}

	return nil
}

func (modifierGroup *ConfigModifierGroup) GetModifiers() ([]ConfigModifier, error) {
	db := database.Mysql()

	var objs []ConfigModifier
	if _, err := db.Select(&objs, "select * from config__modifier where deleted=0 and group_id=? order by name", modifierGroup.Id); err != nil {
		return nil, err
	}

	return objs, nil
}

func (modifierGroup *ConfigModifierGroup) Save() error {
	db := database.Mysql()

	var err error
	if modifierGroup.Id <= 0 {
		err = db.Insert(modifierGroup)
	} else {
		_, err = db.Update(modifierGroup)
	}

	if err != nil {
		return err
	}

	return nil
}

func (modifierGroup *ConfigModifierGroup) Remove() error {
	db := database.Mysql()

	if _, err := db.Exec("update config__modifier set deleted=1 where group_id=?", modifierGroup.Id); err != nil {
		return err
	}

	modifierGroup.Deleted = true
	if _, err := db.Update(modifierGroup); err != nil {
		return err
	}

	return nil
}
