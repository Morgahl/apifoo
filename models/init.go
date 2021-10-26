package models

import "github.com/gobuffalo/pop"

var DB *pop.Connection

func INIT(confName string) (err error) {
	if DB, err = pop.Connect(confName); err != nil {
		return err
	}
	return nil
}
