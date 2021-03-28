package repository

import (
	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	"github.com/sepuka/focalism/def/db"
	"github.com/sepuka/focalism/internal/config"
	"github.com/sepuka/focalism/internal/repository"
)

const (
	TaskRepoDef     = `repo.task.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: TaskRepoDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					dbConn = ctx.Get(db.DataBaseDef).(*pg.DB)
				)

				return repository.NewTaskRepository(dbConn), nil
			},
		})
	})
}
