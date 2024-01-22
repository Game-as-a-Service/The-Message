package enums

const (
	SecretTelegram = 1
	DIRECT         = 2
	DOCUMENT       = 3
)

func ToString(secretTelegramType int) string {
	switch secretTelegramType {
	case SecretTelegram:
		return "密電"
	case DIRECT:
		return "直達"
	case DOCUMENT:
		return "文件"
	default:
		return ""
	}
}
