package task

import (
	"greatestworks/internal/gameplay/task"
	event2 "greatestworks/internal/note/event"
)

type Base struct {
	Id       uint64
	ConfigId uint32
}

func (b *Base) SetStatus(status task.Status) {
}

func (b *Base) OnEvent(event event2.IEvent) {

}

func (b *Base) GetTaskData() task.ITaskData {
	return nil
}
