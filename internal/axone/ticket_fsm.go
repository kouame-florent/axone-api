package axone

import "fmt"

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
