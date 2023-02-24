package database

import (
	"github.com/Unkn0wnCat/calapi/internal/db_model"
	"github.com/objectbox/objectbox-go/objectbox"
)

var (
	ObjectBox *objectbox.ObjectBox
)

func Initialize() error {
	builder := objectbox.NewBuilder()
	builder = builder.Model(db_model.ObjectBoxModel())
	builder = builder.Directory("./data")

	objectBox, err := builder.Build()
	if err != nil {
		return err
	}

	ObjectBox = objectBox

	return nil
}

func Shutdown() {
	if ObjectBox == nil {
		return
	}

	ObjectBox.Close()

	ObjectBox = nil
}
