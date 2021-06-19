package svc

import (
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/api/grpc/gen"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type attachmentSvc struct {
	Repo *repo.AttachmentRepo
}

func NewAttachmentSvc(r *repo.AttachmentRepo) *attachmentSvc {
	return &attachmentSvc{
		Repo: r,
	}
}

func (s *attachmentSvc) RegisterMetas(req *gen.AttachmentRequest, storgeName string) (uuid.UUID, error) {
	meta := req.GetInfo()
	storageName := storgeName
	tickerID, err := uuid.Parse(meta.GetTicketID())
	if err != nil {
		return uuid.UUID{}, err
	}

	att := &axone.Attachment{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UploadedName: meta.GetUploadedName(),
		MimeType:     meta.GetMimeType(),
		Size:         meta.GetSize(),
		StorageName:  storageName,
		Kind:         axone.ATTACHMENT_KIND_REQUEST,
		TicketID:     tickerID,
	}

	//	rep := repo.NewAttachmentRepo(db)
	return s.Repo.Create(att)

}
