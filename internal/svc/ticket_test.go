package svc

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/config"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/store"
	"gorm.io/gorm"
)

var dsn string
var db *gorm.DB

func TestMain(m *testing.M) {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	config.InitEnv(home)
	err = config.CreateAttachmentFolder(home)
	if err != nil {
		log.Fatal(err)
	}

	dsn = config.DataSourceName()
	db = store.OpenDB(dsn)

	if err := store.CreateSchema(db); err != nil {
		log.Fatal(err)
	}

	defer store.CloseDB(db)

	fs := NewFakeSvc(db)

	orgID, err := fs.CreatefakeOrganization()
	if err != nil {
		log.Fatal(err)
	}

	uID, err := fs.CreateFakeUser(orgID)
	if err != nil {
		log.Fatal(err)
	}

	rep := repo.NewRequesterRepo(db)
	eus := NewEndUserSvc(rep)

	requesterID, err = eus.CreateRequester(uID)
	if err != nil {
		log.Fatal(err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)

}

var requesterID uuid.UUID

func TestSendNewTicket(t *testing.T) {

	rep := repo.NewTicketRepo(db)
	s := NewTicketSvc(rep)

	tID, err := uuid.Parse("3d13de55-6b84-43f0-8ae7-fb5a438dfaf7")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("RequesterID: %s", requesterID)
	id, err := s.SendNewTicket(tID, "Réacteur en surchauffe", "comment refroidir le réacteur en surchauffe",
		axone.TICKET_TYPE_PROBLEM, requesterID)
	if err != nil {
		t.Errorf("expected no errors got: %v", err)
	}
	if id.String() != "3d13de55-6b84-43f0-8ae7-fb5a438dfaf7" {
		t.Errorf("expected id '3d13de55-6b84-43f0-8ae7-fb5a438dfaf7' got: %s", id.String())
	}
}
