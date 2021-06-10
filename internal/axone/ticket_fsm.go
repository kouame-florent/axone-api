package axone

import "fmt"

type TicketStatus string

const (
	TICKET_STATUS_NEW     TicketStatus = "New"     //after creation
	TICKET_STATUS_OPEN    TicketStatus = "Open"    //has been evaluated and is assigned
	TICKET_STATUS_PENDING TicketStatus = "Pending" //need more informations from end user
	TICKET_STATUS_SOLVED  TicketStatus = "Solved"  //issue no longer exists, or the work has been completed
	TICKET_STATUS_CLOSED  TicketStatus = "Closed"  //resolved and a sufficient amount of time has passed (1 week)
)

type Event uint

const (
	EVENT_CREATION = iota
	EVENT_ASSIGMENT
	EVENT_INFO_REQUIRED
	EVENT_COMMENT_ADDED
	EVENT_WORK_COMPLETED
	EVENT_TIME_HAS_PASSED
)

type EventStatus struct {
	Evt    Event
	Status TicketStatus
}

type TransitionFunc func(r *Ticket) *Ticket

var StateTransitionTable = map[EventStatus]TransitionFunc{
	{EVENT_ASSIGMENT, TICKET_STATUS_NEW}: func(t *Ticket) *Ticket {
		t.Status = TICKET_STATUS_OPEN
		return t
	},
	{EVENT_INFO_REQUIRED, TICKET_STATUS_OPEN}: func(t *Ticket) *Ticket {
		t.Status = TICKET_STATUS_PENDING
		return t
	},
	{EVENT_COMMENT_ADDED, TICKET_STATUS_PENDING}: func(t *Ticket) *Ticket {
		t.Status = TICKET_STATUS_OPEN
		return t
	},
	{EVENT_WORK_COMPLETED, TICKET_STATUS_OPEN}: func(t *Ticket) *Ticket {
		t.Status = TICKET_STATUS_SOLVED
		return t
	},
	{EVENT_TIME_HAS_PASSED, TICKET_STATUS_SOLVED}: func(t *Ticket) *Ticket {
		t.Status = TICKET_STATUS_CLOSED
		return t
	},
}

func (t *Ticket) ExecuteEvt(evt Event) (*Ticket, error) {
	evtStat := EventStatus{evt, t.Status}

	if f := StateTransitionTable[evtStat]; f == nil {
		return t, fmt.Errorf("%s", "unknown event")
	} else {

		return f(t), nil
	}

}
