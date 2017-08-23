package encode

import (
    "bytes"
    "encoding/hex"
    "encoding/base64"
)

func HexToBase64(s string) (string, error) {
    decoded, err := hex.DecodeString(s)
    if err != nil {
        return "", err
    }

    buf := new(bytes.Buffer)
    encoder := base64.NewEncoder(base64.StdEncoding, buf)
    _, writeErr := encoder.Write(decoded)
    if writeErr != nil {
        return "", writeErr
    }
    encoder.Close()
    return buf.String(), nil
}
