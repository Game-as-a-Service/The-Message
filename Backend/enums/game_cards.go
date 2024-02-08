package enums

const (
	LockOn      = "鎖定"
	LureAway    = "調虎離山"
	Probe       = "試探"
	Intercept   = "截獲"
	Decipher    = "破譯"
	Diversion   = "退回"
	Burn        = "燒毀"
	BlurOfTruth = "真偽莫辯"
	SeeThrough  = "識破"
)

func ToIntelligenceType(actionType string) int {
	switch actionType {
	case LockOn, LureAway, Probe, Decipher:
		return SecretTelegram
	case Intercept, Burn, SeeThrough:
		return Direct
	case Diversion, BlurOfTruth:
		return Document
	default:
		return 0
	}
}
