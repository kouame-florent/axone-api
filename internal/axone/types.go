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
	FirstName      string    `gorm:"type:varchar(100)"`
	LastName       string    `gorm:"type:varchar(100)"`
	Email          string    `gorm:"type:varchar(100)"`
	PhoneNumber    string    `gorm:"type:varchar(100)"`
	Login          string    `gorm:"type:varchar(100)"`
	Password       string    `gorm:"type:varchar(100)"`
	OrganizationID uuid.UUID `gorm:"type:varchar(36)"`
	Requesters     []Requester
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
	ROLE_REQUESTER     RoleValue = "REQUESTER"
	ROLE_AGENT         RoleValue = "AGENT"
	ROLE_ADMINISTRATOR RoleValue = "ADMINISTRATOR"
)

type Requester struct {
	UserID  uuid.UUID `gorm:"type:varchar(36)"`
	Tickets []Ticket  `gorm:"foreignKey:RequesterID;references:UserID"`
}

type Agent struct {
	UserID  uuid.UUID `gorm:"type:varchar(36)"`
	Bio     string
	Level   AgentLevel
	Tickets []Ticket `gorm:"foreignKey:AssigneeID;references:UserID"`
}

type AgentLevel string

const (
	LEVEL_ONE AgentLevel = "ONE" //can manage ticket with any status
	LEVEL_TWO AgentLevel = "TWO" //cannot manage ticket with new status
)

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
	Subject     string    `gorm:"type:varchar(150)"`
	Request     string    `gorm:"type:varchar(1000)"`
	Answer      string    `gorm:"type:varchar(1000)"`
	RequesterID uuid.UUID `gorm:"type:varchar(36)"`
	AssigneeID  uuid.UUID `gorm:"type:varchar(36)"`
	Status      TicketStatus
	Type        TicketType
	Priority    TicketPriority
	Rate        uint
	Tags        []*Tag `gorm:"many2many:ticket_tags;"`
	Comments    []Comment
	Attachments []Attachment
}

type TicketType string

const (
	TICKET_TYPE_QUESTION TicketType = "Question"
	TICKET_TYPE_PROBLEM  TicketType = "Problem"
	TICKET_TYPE_TASK     TicketType = "Task"
)

//ticket priority are only seen by agents
type TicketPriority string

const (
	TICKET_PRIORITY_LOW    TicketPriority = "Low"
	TICKET_PRIORITY_MEDIUM TicketPriority = "Medium"
	TICKET_PRIORITY_HIGH   TicketPriority = "High"
	TICKET_PRIORITY_URGENT TicketPriority = "Urgent"
)

type Tag struct {
	Model
	Name    string    `gorm:"type:varchar(50)"`
	Tickets []*Ticket `gorm:"many2many:ticket_tags;"`
}

type Comment struct {
	Model
	Text     string          `gorm:"type:varchar(300)"`
	TicketID uuid.UUID       `gorm:"type:varchar(36)"`
	Kind     CommentKind     `gorm:"not null"`
	Category CommentCategory `gorm:"not null"`
}

type Attachment struct {
	Model
	UploadedName string `gorm:"type:varchar(100)"`
	Size         int64
	MimeType     string `gorm:"type:varchar(100)"`
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

type Assignment struct {
	Model
	TickerID        uuid.UUID `gorm:"type:varchar(36)"`
	AgentID         uuid.UUID `gorm:"type:varchar(36)"`
	Assignment_date time.Time
}

type Knowledge struct {
	Model
	Problem  string `gorm:"type:varchar(1000)"`
	Solution string `gorm:"type:varchar(1000)"`
}
