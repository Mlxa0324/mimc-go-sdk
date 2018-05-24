package mimc

import (
	"container/list"
	"github.com/Xiaomi-mimc/mimc-go-sdk/message"
)

type Token interface {
	FetchToken() *string
}

type StatusDelegate interface {
	/**
	 * @param[isOnline bool] true: 在线，false：离线
	 * @param[errType *string] 登录失败类型
	 * @param[errReason *string] 登录失败原因
	 * @param[errDescription *string] 登录失败原因描述
	 */
	HandleChange(isOnline bool, errType, errReason, errDescription *string)
}

type MessageHandlerDelegate interface {
	HandleMessage(packets *list.List)
	HandleGroupMessage(packets *list.List)
	HandleServerAck(packetId *string, sequence, timestamp *int64)
	HandleSendMessageTimeout(message *msg.P2PMessage)
	HandleSendGroupMessageTimeout(message *msg.P2TMessage)
}
