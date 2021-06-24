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
	Email          string    `gorm:"type:varchar(100);unique"`
	PhoneNumber    string    `gorm:"type:varchar(100)"`
	Login          string    `gorm:"type:varchar(100);unique"`
	Password       string    `gorm:"type:varchar(100)"`
	OrganizationID uuid.UUID `gorm:"type:varchar(36)"`
	Requesters     []Requester
	Administrators []Administrator
	Agents         []Agent
	Roles          []*Role `gorm:"many2many:user_roles;"`
	Status         UserStatus
}

type UserStatus string

const (
	USER_STATUS_ENABLED  UserStatus = "ENABLED"
	USER_STATUS_DISABLED UserStatus = "DISABLED"
)

type Role struct {
	Model
	Value RoleValue
	Users []*User `gorm:"many2many:user_roles;"`
}

type RoleValue string

const (
	ROLE_VALUE_REQUESTER     RoleValue = "REQUESTER"
	ROLE_VALUE_AGENT         RoleValue = "AGENT"
	ROLE_VALUE_ADMINISTRATOR RoleValue = "ADMINISTRATOR"
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
	AGENT_LEVEL_ONE AgentLevel = "ONE" //can manage ticket with any status and route to any agent
	AGENT_LEVEL_TWO AgentLevel = "TWO" //cannot manage ticket with new status, can route to agent in the same organization only
)

type Administrator struct {
	UserID uuid.UUID `gorm:"type:varchar(36)"`
}

//Notif *fyne.Notification
//Notif *fyne.Notification

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
	TicketType  TicketType
	Priority    TicketPriority
	Rate        uint
	Tags        []*Tag `gorm:"many2many:ticket_tags;"`
	Assignments []Assignment
	Comments    []Comment
	Attachments []Attachment
}

type TicketStatus string

const (
	TICKET_STATUS_NEW     TicketStatus = "NEW"     //after creation
	TICKET_STATUS_OPEN    TicketStatus = "OPEN"    //has been evaluated and is assigned
	TICKET_STATUS_PENDING TicketStatus = "PENDING" //need more informations from end user
	TICKET_STATUS_SOLVED  TicketStatus = "SOLVED"  //issue no longer exists, or the work has been completed
	TICKET_STATUS_CLOSED  TicketStatus = "CLOSED"  //resolved and a sufficient amount of time has passed (1 week)
)

type TicketType string

const (
	TICKET_TYPE_QUESTION TicketType = "QUESTION"
	TICKET_TYPE_PROBLEM  TicketType = "PROBLEM"
	TICKET_TYPE_TASK     TicketType = "TASK"
)

//ticket priority are only seen by agents
type TicketPriority string

const (
	TICKET_PRIORITY_LOW    TicketPriority = "LOW"
	TICKET_PRIORITY_MEDIUM TicketPriority = "MEDIUM"
	TICKET_PRIORITY_HIGH   TicketPriority = "HIGH"
	TICKET_PRIORITY_URGENT TicketPriority = "URGENT"
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
	ASSIGNMENT_STATUS_ENABLED  AssignmentStatus = "ENABLED"
	ASSIGNMENT_STATUS_DISABLED AssignmentStatus = "DISABLED"
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
	Size         uint32
	MimeType     string `gorm:"type:varchar(100)"`
	StorageName  string
	Kind         AttachmentKind
	TicketID     uuid.UUID `gorm:"type:varchar(36)"`
}

type AttachmentKind string

const (
	ATTACHMENT_KIND_REQUEST AttachmentKind = "REQUEST"
	ATTACHMENT_KIND_ANSWER  AttachmentKind = "ANSWER"
)

type CommentKind string

const (
	COMMENT_KIND_QUESTION CommentKind = "QUESTION"
	COMMENT_KIND_ANSWER   CommentKind = "ANSWER"
)

type CommentCategory string

const (
	COMMENT_CATEGORY_PUBLIC  CommentCategory = "PUBLIC"
	COMMENT_CATEGORY_PRIVATE CommentCategory = "PRIVATE"
)

type Knowledge struct {
	Model
	Problem  string `gorm:"type:varchar(1000)"`
	Solution string `gorm:"type:varchar(1000)"`
}

type Credential struct {
	Login    string
	Password string
}

type UserProfile struct {
	UserID    string
	Login     string
	Password  string
	Email     string
	FirstName string
	LastName  string
}
