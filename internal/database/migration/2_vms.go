package migration

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Info("Ceating table vms...")
		_, err := db.Exec(`CREATE TABLE vms 
							(id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
							 name VARCHAR(50) UNIQUE, 
							 status INT4,
							 host_id UUID REFERENCES hosts(id),
							 vcpu INTEGER,
							 memory INTEGER,
							 kernel VARCHAR(255),
							 root_file_system VARCHAR(255),
							 address VARCHAR(16) UNIQUE)`)
		return err
	}, func(db migrations.DB) error {
		log.Info("Dropping table vms...")
		_, err := db.Exec(`DROP TABLE vms`)
		return err
	})
}
