package chat

import (
	"github.com/nsqio/go-nsq"
	"github.com/phuhao00/greatestworks-proto/messageId"
	"google.golang.org/protobuf/proto"
)

type MangerOwner interface {
	BroadcastSystemMsg(message proto.Message)
	BroadcastOnlineChatMsg(message proto.Message)
	BroadcastCrossZoneChatMsg(message proto.Message)
	BroadcastZoneChatMsg(message proto.Message)
	BroadcastCrossSrvChatMsg(message proto.Message)
	SyncOfflineOnlineChatMsg() []proto.Message
}

type Transfer interface {
	ForwardCrossZoneChatMsg(proto.Message)
}

type ServerTransfer interface {
	ForwardCrossSrvChatMsg(proto.Message)
}

type PrivateTransfer interface {
	ForwardPlayer(proto.Message)
}

type ZoneTransfer interface {
	ForwardZoneChatMsg(proto.Message)
}

type SystemTransfer interface {
	ForwardSysMsg(proto.Message)
}

type Handler interface {
	InitNsqHandler(channel string)
	HandleMessage(message nsq.Message) error
	PublishChatMsg(chatMsg interface{}) error
	Stop()
}

type Owner interface {
	SendMsg(ID messageId.MessageId, message proto.Message)
}
