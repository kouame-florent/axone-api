package axo

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Model
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	EndUsers       []EndUser
	Administrators []Administrator
	Agents         []Agent
}

type EndUser struct {
	UserID uuid.UUID `gorm:"type:varchar(36)"`
}

type Agent struct {
	UserID uuid.UUID `gorm:"type:varchar(36)"`
}

type Administrator struct {
	UserID uuid.UUID `gorm:"type:varchar(36)"`
}

type Organization struct {
	Model
	Name      string
	Users     []User
	Questions []Question
}

type Ticket struct {
	Model
	Question
	Answer
	EndUserID  uuid.UUID `gorm:"type:varchar(36)" json:"end_user_id"`
	AssigneeID uuid.UUID `gorm:"type:varchar(36)" json:"assignee_id"`
	Status     TicketStatus
	Tags       []*Tag `gorm:"many2many:question_question_tags;"`
}

type Question struct {
	Title               string
	Body                string
	QuestionAttachments []QuestionAttachment
	QuestionComments    []QuestionComment
}

type Answer struct {
	Body              string
	AnswerAttachments []AnswerAttachment
	AnswerComments    []AnswerComment
}

type QuestionComment struct {
	Body     string
	TicketID uuid.UUID `gorm:"type:varchar(36)"`
}

type AnswerComment struct {
	Body     string
	TicketID uuid.UUID `gorm:"type:varchar(36)"`
}

type CommentStatus string

const (
	PUBLIC  CommentStatus = "Public"
	PRIVATE CommentStatus = "Private"
)

type priority string

const (
	PR_Low    priority = "Low"
	PR_Medium priority = "Medium"
	PR_High   priority = "High"
	PR_Urgent priority = "Urgent"
)

type role string

const (
	ROLE_ENDUSER role = "END_USER"
	PR_AGENT     role = "AGENT"
	PR_ADMIN     role = "ADMIN"
)

type Assignment struct {
	Model
	TickerID        uuid.UUID `gorm:"type:varchar(36)" json:"question_id"`
	AgentID         uuid.UUID `gorm:"type:varchar(36)" json:"assignee_id"`
	Assignment_date time.Time
}

type QuestionAttachment struct {
	Model
	UploadedName string    `json:"uploaded_name"`
	Size         int64     `json:"size"`
	MimeType     string    `json:"mime_type"`
	StorageName  string    `json:"storage_name"`
	TickerID     uuid.UUID `gorm:"type:varchar(36)"`
}

type AnswerAttachment struct {
	Model
	UploadedName string
	Size         int64
	MimeType     string
	StorageName  string
	TicketID     uuid.UUID `gorm:"type:varchar(36)"`
}

type Tag struct {
	Model
	Name    string
	Tickets []*Ticket `gorm:"many2many:ticket_tags;"`
}

type Knowledge struct {
	Model
	Problem  string
	Solution string
}
