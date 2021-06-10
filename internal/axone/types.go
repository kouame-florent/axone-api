package axone

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
	Login          string
	Password       string
	OrganizationID uuid.UUID `gorm:"type:varchar(36)"`
	EndUsers       []EndUser
	Administrators []Administrator
	Agents         []Agent
	Roles          []*Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	Model
	Value RoleValue
	Users []*User `gorm:"many2many:user_roles;"`
}

type RoleValue string

const (
	ROLE_END_USER      RoleValue = "END_USER"
	ROLE_AGENT         RoleValue = "AGENT"
	ROLE_ADMINISTRATOR RoleValue = "ADMINISTRATOR"
)

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
	Name  string
	Users []User
}

type Ticket struct {
	Model
	Title       string
	Question    string
	Answer      string
	EndUserID   uuid.UUID `gorm:"type:varchar(36)" json:"end_user_id"`
	AssigneeID  uuid.UUID `gorm:"type:varchar(36)" json:"assignee_id"`
	Status      TicketStatus
	Tags        []*Tag `gorm:"many2many:ticket_tags;"`
	Comments    []Comment
	Attachments []Attachment
}

type Tag struct {
	Model
	Name    string
	Tickets []*Ticket `gorm:"many2many:ticket_tags;"`
}

type Comment struct {
	Model
	Body     string
	TicketID uuid.UUID       `gorm:"type:varchar(36)"`
	Kind     CommentKind     `gorm:"not null"`
	Category CommentCategory `gorm:"not null"`
}

type Attachment struct {
	Model
	UploadedName string
	Size         int64
	MimeType     string
	StorageName  string
	Kind         AttachmentKind
	TicketID     uuid.UUID `gorm:"type:varchar(36)"`
}

type AttachmentKind string

const (
	QUESTION_ATTACHMENT AttachmentKind = "QUESTION"
	ANSWER_ATTACHMENT   AttachmentKind = "ANSWER"
)

type CommentKind string

const (
	QUESTION CommentKind = "QUESTION"
	ANSWER   CommentKind = "ANSWER"
)

type CommentCategory string

const (
	PUBLIC  CommentCategory = "Public"
	PRIVATE CommentCategory = "Private"
)

type priority string

const (
	PR_Low    priority = "Low"
	PR_Medium priority = "Medium"
	PR_High   priority = "High"
	PR_Urgent priority = "Urgent"
)

type Assignment struct {
	Model
	TickerID        uuid.UUID `gorm:"type:varchar(36)" json:"question_id"`
	AgentID         uuid.UUID `gorm:"type:varchar(36)" json:"assignee_id"`
	Assignment_date time.Time
}

type Knowledge struct {
	Model
	Problem  string
	Solution string
}
