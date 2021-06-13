package axone

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
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
	ROLE_VALUE_REQUESTER     RoleValue = "requester"
	ROLE_VALUE_AGENT         RoleValue = "agent"
	ROLE_VALUE_ADMINISTRATOR RoleValue = "administrator"
)

type Requester struct {
	UserID  uuid.UUID `gorm:"type:varchar(36)"`
	Tickets []Ticket  `gorm:"foreignKey:RequesterID;references:UserID"`
}

type Agent struct {
	UserID      uuid.UUID `gorm:"type:varchar(36)"`
	Bio         string
	Level       AgentLevel
	Assignments []Assignment `gorm:"foreignKey:AssigneeID;references:UserID"`
}

type AgentLevel string

const (
	AGENT_LEVEL_ONE AgentLevel = "one" //can manage ticket with any status and route to any agent
	AGENT_LEVEL_TWO AgentLevel = "two" //cannot manage ticket with new status, can route to agent in the same organization only
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
	Status      TicketStatus
	Type        TicketType
	Priority    TicketPriority
	Rate        uint
	Tags        []*Tag `gorm:"many2many:ticket_tags;"`
	Assignments []Assignment
	Comments    []Comment
	Attachments []Attachment
}

type TicketType string

const (
	TICKET_TYPE_QUESTION TicketType = "question"
	TICKET_TYPE_PROBLEM  TicketType = "problem"
	TICKET_TYPE_TASK     TicketType = "task"
)

//ticket priority are only seen by agents
type TicketPriority string

const (
	TICKET_PRIORITY_LOW    TicketPriority = "low"
	TICKET_PRIORITY_MEDIUM TicketPriority = "medium"
	TICKET_PRIORITY_HIGH   TicketPriority = "high"
	TICKET_PRIORITY_URGENT TicketPriority = "urgent"
)

type Assignment struct {
	//	Model
	TicketID   uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	AssigneeID uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	//Assignment_date time.Time
	Status         AssignmentStatus
	DisabledReason string `gorm:"type:varchar(150)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type AssignmentStatus string

const (
	ASSIGNMENT_STATUS_ENABLED  AssignmentStatus = "enabled"
	ASSIGNMENT_STATUS_DISABLED AssignmentStatus = "disabled"
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
	ATTACHMENT_KIND_QUESTION AttachmentKind = "question"
	ATTACHMENT_KIND_ANSWER   AttachmentKind = "answer"
)

type CommentKind string

const (
	COMMENT_KIND_QUESTION CommentKind = "question"
	COMMENT_KIND_ANSWER   CommentKind = "answer"
)

type CommentCategory string

const (
	COMMENT_CATEGORY_PUBLIC  CommentCategory = "public"
	COMMENT_CATEGORY_PRIVATE CommentCategory = "private"
)

type Knowledge struct {
	Model
	Problem  string `gorm:"type:varchar(1000)"`
	Solution string `gorm:"type:varchar(1000)"`
}
