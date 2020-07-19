package event

const (
	// WSConnected ユーザーがWSストリームに接続した
	// 	Fields:
	// 		user_id: uuid.UUID
	//		req: *http.Request
	WSConnected = "ws.connected"
	// WSDisconnected ユーザーがWSストリームから切断した
	// 	Fields:
	// 		user_id: uuid.UUID
	//		req: *http.Request
	WSDisconnected = "ws.disconnected"
	// StateUpdated Stateが変更された
	// 	Fields:
	// 		state: string
	StateUpdated = "state.updated"
)