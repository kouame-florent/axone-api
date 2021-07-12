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
)

//var dsn string
//var db *gorm.DB

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

	dsn := config.DataSourceName()
	db, err := store.OpenDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateSchema(db); err != nil {
		log.Fatal(err)
	}

	//defer store.CloseDB(db)

	fs := NewFakeSvc(db)

	orgID, err := fs.CreatefakeOrganization()
	if err != nil {
		log.Fatal(err)
	}

	_, err = fs.createFakeTags()
	if err != nil {
		log.Fatal(err)
	}

	uID, err := fs.CreateFakeRequesterUser(orgID)
	if err != nil {
		log.Fatal(err)
	}

	rep := repo.NewRequesterRepo(db)
	eus := NewRequesterSvc(rep)

	requesterID, err = eus.CreateRequester(uID)
	if err != nil {
		log.Fatal(err)
	}

	auid, err := fs.CreateFakeLevelOneAgentUser(orgID)
	if err != nil {
		log.Fatal(err)
	}

	agtRep := repo.NewAgentRepo(db)
	as := NewAgentSvc(agtRep)

	agentID, err = as.CreateAgent(auid, axone.AGENT_LEVEL_ONE, "maitrise en histoire")
	if err != nil {
		log.Fatal(err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)

}

var requesterID uuid.UUID
var agentID uuid.UUID

func TestSendNewTicket(t *testing.T) {

	dsn := config.DataSourceName()
	db, err := store.OpenDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	rep := repo.NewTicketRepo(db)
	svc := NewTicketSvc(rep)

	tID := uuid.MustParse("3d13de55-6b84-43f0-8ae7-fb5a438dfaf7")

	log.Printf("TicketID: %s", tID)
	log.Printf("service: %+v", svc)
	log.Printf("RequesterID: %s", requesterID)

	id, err := svc.SendNewTicket(tID, "Réacteur en surchauffe", "&{Repo:0xc0000b6b60}&{Repo:0xc0000b6b60}mment refroidir le réacteur en surchauffe",
		axone.TICKET_TYPE_PROBLEM, requesterID)
	if err != nil {
		t.Errorf("expected no errors got: %v", err)

	}

	if id.String() != "3d13de55-6b84-43f0-8ae7-fb5a438dfaf7" {
		t.Errorf("expected id '3d13de55-6b84-43f0-8ae7-fb5a438dfaf7' got: %s", id.String())
	}

}
