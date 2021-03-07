package db

import (
	"github.com/go-pg/pg"
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	"github.com/sepuka/focalism/internal/config"
	"net"
	"strconv"
)

const DataBaseDef = "db.def"

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: DataBaseDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					db *pg.DB
				)

				db = pg.Connect(&pg.Options{
					User:     cfg.Db.User,
					Password: cfg.Db.Password,
					Addr:     net.JoinHostPort(cfg.Db.Host, strconv.Itoa(cfg.Db.Port)),
					Database: cfg.Db.Name,
				})

				_, err := db.Exec(`SET timezone TO 'UTC'`)

				return db, err
			},
		})
	})
}
