package ws

// TargetFunc メッセージ送信対象関数
type TargetFunc func(s Client) bool

// TargetAll 全セッションを対象に送信します
func TargetAll() TargetFunc {
	return func(_ Client) bool {
		return true
	}
}
