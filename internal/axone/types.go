package axo

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Customers   []Customer
	Agents      []Agent
}

type Customer struct {
	UserID string `gorm:"type:varchar(36)"`
}

type Agent struct {
	UserID string  `gorm:"type:varchar(36)"`
	Pool   []*Pool `gorm:"many2many:agent_pools;"`
}

type Pool struct {
	Model
	Name      string   `json:"name"`
	Agents    []*Agent `gorm:"many2many:agent_pools;"`
	Questions []Question
}

func NewPool(name string) *Pool {
	return &Pool{
		Model: Model{
			ID: uuid.New(),
		},
		Name: name,
	}
}

type Message struct {
	Model
	Body string `json:"body"`
}

type Question struct {
	Message
	Title               string         `json:"title"`
	Footer              string         `json:"footer"`
	QuestionTags        []*QuestionTag `gorm:"many2many:question_question_tags;"`
	ChannelID           uuid.UUID      `gorm:"type:varchar(36)" json:"channel_id"`
	PoolID              uuid.UUID      `gorm:"type:varchar(36)" json:"pool_id"`
	Answers             []Answer
	CustomerID          uuid.UUID `gorm:"type:varchar(36)" json:"customer_id"`
	Comments            []QuestionComment
	QuestionAttachments []QuestionAttachment
	Assignments         []Assignment
	Status              questionStatus `json:"status"`
}

func NewQuestion(title, body string, channelID uuid.UUID, poolID uuid.UUID, customerID uuid.UUID) *Question {
	return &Question{
		Title:      title,
		ChannelID:  channelID,
		PoolID:     poolID,
		CustomerID: customerID,
		Message: Message{
			Model: Model{
				ID: uuid.New(),
			},
			Body: body,
		},
	}
}

type questionStatus string

const (
	QS_New     questionStatus = "New"
	QS_Open    questionStatus = "Open"
	QS_Pending questionStatus = "Pending"
	QS_Solved  questionStatus = "Solved"
	QS_Closed  questionStatus = "Closed"
)

type priority string

const (
	PR_Low    priority = "Low"
	PR_Medium priority = "Medium"
	PR_High   priority = "High"
	PR_Urgent priority = "Urgent"
)

type AnswerStatus string

const (
	AS_Open      AnswerStatus = "Open"    //open for validation
	AS_Pending   AnswerStatus = "Pending" //need additional information
	AS_Validated AnswerStatus = "Validated"
)

type Answer struct {
	Message
	AnswerTags        []*AnswerTag `gorm:"many2many:answer_answer_tags;"`
	QuestionID        uuid.UUID    `gorm:"type:varchar(36)" json:"question_id"`
	AssigneeID        uuid.UUID    `gorm:"type:varchar(36)" json:"assignee_id"`
	Comments          []AnswerComment
	AnswerAttachments []AnswerAttachment
	Status            AnswerStatus `json:"status"`
}

/*
type ChannelType string

const (
	CT_AXONE ChannelType = "AXONE"
	CT_EMAIL ChannelType = "EMAIL"
)
*/
/*
type Channel struct {
	Model
	Name      string      `json:"name"`
	Type      ChannelType `json:"type"`
	Questions []Question
}
*/

/*
func NewChannel(name string, channelType ChannelType) *Channel {
	return &Channel{
		Model: Model{
			ID: uuid.New(),
		},
		Name: name,
		Type: channelType,
	}
}
*/

type Assignment struct {
	Model
	QuestionID      uuid.UUID `gorm:"type:varchar(36)" json:"question_id"`
	AssigneeID      uuid.UUID `gorm:"type:varchar(36)" json:"assignee_id"`
	Assignment_date time.Time `json:"assignment_date"`
}

func NewAssignment(questionID uuid.UUID, assigneeID uuid.UUID) *Assignment {
	return &Assignment{
		Model: Model{
			ID: uuid.New(),
		},
		QuestionID: questionID,
		AssigneeID: assigneeID,
	}
}

type QuestionComment struct {
	Message
	QuestionID uuid.UUID `json:"question_id"`
}

type AnswerComment struct {
	Message
	AnswerID uuid.UUID `json:"answer_id"`
}

type Attachment struct {
	Model
	UploadedName string `json:"uploaded_name"`
	Size         int64  `json:"size"`
	MimeType     string `json:"mime_type"`
	StorageName  string `json:"storage_name"`
}

type QuestionAttachment struct {
	Attachment
	QuestionID uuid.UUID `json:"question_id"`
}

type AnswerAttachment struct {
	Attachment
	AnswerID uuid.UUID `json:"answer_id"`
}

type QuestionTag struct {
	Model
	Name      string      `json:"name"`
	Questions []*Question `gorm:"many2many:question_question_tags;"`
}

func NewQuestionTag(name string) *QuestionTag {
	return &QuestionTag{
		Model: Model{
			ID: uuid.New(),
		},
		Name: name,
	}
}

type AnswerTag struct {
	Model
	Name    string    `gorm:"unique" json:"name"`
	Answers []*Answer `gorm:"many2many:answer_answer_tags;"`
}
