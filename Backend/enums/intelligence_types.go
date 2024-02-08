package enums

const (
	SecretTelegram = 1
	Direct         = 2
	Document       = 3
)

func ToString(secretTelegramType int) string {
	switch secretTelegramType {
	case SecretTelegram:
		return "密電"
	case Direct:
		return "直達"
	case Document:
		return "文件"
	default:
		return ""
	}
}
