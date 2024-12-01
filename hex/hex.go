package hex

import "errors"

func ToHexString(bytes []byte, toLower bool) string {
	var hexChars string
	if toLower {
		hexChars = "0123456789abcdef"
	} else {
		hexChars = "0123456789ABCDEF"
	}
	// Helper function to convert a byte to a hex string
	byteToHex := func(b byte) string {
		return string([]byte{hexChars[b>>4], hexChars[b&0x0F]})
	}
	hexString := ""
	for i := 0; i < len(bytes); i++ {
		hexString += byteToHex(bytes[i])
	}
	return hexString
}

func FromHexString(hexString string) ([]byte, error) {
	// hexString = strings.ReplaceAll(hexString, " ", "")
	if len(hexString)%2 != 0 {
		return nil, errors.New("invalid hexadecimal string length")
	}
	bytes := make([]byte, len(hexString)/2)
	for i := 0; i < len(hexString); i += 2 {
		high, err := hexCharToByte(hexString[i])
		if err != nil {
			return nil, err
		}
		low, err := hexCharToByte(hexString[i+1])
		if err != nil {
			return nil, err
		}
		bytes[i/2] = (high << 4) | low
	}
	return bytes, nil
}

func hexCharToByte(c byte) (byte, error) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', nil
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, nil
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, nil
	default:
		return 0, errors.New("invalid hexadecimal character")
	}
}