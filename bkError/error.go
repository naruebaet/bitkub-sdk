package bkError

const (
	NoError                            = 0
	InvalidJSONPayload                 = 1
	MissingXBTKAPIKEY                  = 2
	InvalidAPIKey                      = 3
	APIPendingForActivation            = 4
	IPNotAllowed                       = 5
	MissingInvalidSignature            = 6
	MissingTimestamp                   = 7
	InvalidTimestamp                   = 8
	InvalidUser                        = 9
	InvalidParameter                   = 10
	InvalidSymbol                      = 11
	InvalidAmount                      = 12
	InvalidRate                        = 13
	ImproperRate                       = 14
	AmountTooLow                       = 15
	FailedToGetBalance                 = 16
	WalletIsEmpty                      = 17
	InsufficientBalance                = 18
	FailedToInsertOrderIntoDb          = 19
	FailedToDeductBalance              = 20
	InvalidOrderForCancellation        = 21
	InvalidSide                        = 22
	FailedToUpdateOrderStatus          = 23
	InvalidOrderForLookup              = 24
	KYCLevel1IsRequiredToProceed       = 25
	LimitExceeds                       = 30
	PendingWithdrawalExists            = 40
	InvalidCurrencyForWithdrawal       = 41
	AddressIsNotInWhitelist            = 42
	FailedToDeductCrypto               = 43
	FailedToCreateWithdrawalRecord     = 44
	NonceHasToBeNumeric                = 45
	InvalidNonce                       = 46
	WithdrawalLimitExceeds             = 47
	InvalidBankAccount                 = 48
	BankLimitExceeds                   = 49
	PendingWithdrawalExists2           = 50
	WithdrawalIsUnderMaintenance       = 51
	InvalidPermission                  = 52
	InvalidInternalAddress             = 53
	AddressHasBeenDeprecated           = 54
	CancelOnlyMode                     = 55
	UserHasBeenSuspendedFromPurchasing = 56
	UserHasBeenSuspendedFromSelling    = 57
	ServerError                        = 90
)

// ErrorText is a function for get error text or description from error code
func ErrorText(errorCode int) string {
	switch errorCode {
	case NoError:
		return "No error"
	case InvalidJSONPayload:
		return "Invalid JSON payload"
	case MissingXBTKAPIKEY:
		return "Missing X-BTK-APIKEY"
	case InvalidAPIKey:
		return "Invalid API key"
	case APIPendingForActivation:
		return "API pending for activation"
	case IPNotAllowed:
		return "IP not allowed"
	case MissingInvalidSignature:
		return "Missing / invalid signature"
	case MissingTimestamp:
		return "Missing timestamp"
	case InvalidTimestamp:
		return "Invalid timestamp"
	case InvalidUser:
		return "Invalid user"
	case InvalidParameter:
		return "Invalid parameter"
	case InvalidSymbol:
		return "Invalid symbol"
	case InvalidAmount:
		return "Invalid amount"
	case InvalidRate:
		return "Invalid rate"
	case ImproperRate:
		return "Improper rate"
	case AmountTooLow:
		return "Amount too low"
	case FailedToGetBalance:
		return "Failed to get balance"
	case WalletIsEmpty:
		return "Wallet is empty"
	case InsufficientBalance:
		return "Insufficient balance"
	case FailedToInsertOrderIntoDb:
		return "Failed to insert order into db"
	case FailedToDeductBalance:
		return "Failed to deduct balance"
	case InvalidOrderForCancellation:
		return "Invalid order for cancellation"
	case InvalidSide:
		return "Invalid side"
	case FailedToUpdateOrderStatus:
		return "Failed to update order status"
	case InvalidOrderForLookup:
		return "Invalid order for lookup"
	case KYCLevel1IsRequiredToProceed:
		return "KYC level 1 is required to proceed"
	case LimitExceeds:
		return "Limit exceeds"
	case PendingWithdrawalExists:
		return "Pending withdrawal exists"
	case InvalidCurrencyForWithdrawal:
		return "Invalid currency for withdrawal"
	case AddressIsNotInWhitelist:
		return "Address is not in whitelist"
	case FailedToDeductCrypto:
		return "Failed to deduct crypto"
	case FailedToCreateWithdrawalRecord:
		return "Failed to create withdrawal record"
	case NonceHasToBeNumeric:
		return "Nonce has to be numeric"
	case InvalidNonce:
		return "Invalid nonce"
	case WithdrawalLimitExceeds:
		return "Withdrawal limit exceeds"
	case InvalidBankAccount:
		return "Invalid bank account"
	case BankLimitExceeds:
		return "Bank limit exceeds"
	case PendingWithdrawalExists2:
		return "Pending withdrawal exists"
	case WithdrawalIsUnderMaintenance:
		return "Withdrawal is under maintenance"
	case InvalidPermission:
		return "Invalid permission"
	case InvalidInternalAddress:
		return "Invalid internal address"
	case AddressHasBeenDeprecated:
		return "Address has been deprecated"
	case CancelOnlyMode:
		return "Cancel only mode"
	case UserHasBeenSuspendedFromPurchasing:
		return "User has been suspended from purchasing"
	case UserHasBeenSuspendedFromSelling:
		return "User has been suspended from selling"
	case ServerError:
		return "Server error (please contact support)"
	default:
		return "error code not found!"
	}
}
