package amqp

type Queue string

const QueueEnter Queue = "enter"
const QueueStep1 Queue = "step1"
const QueueStep2 Queue = "step2"
const QueueStep3 Queue = "step3"

func (t *Queue) String() string {
	return string(*t)
}
