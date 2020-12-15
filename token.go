package ngxhmactoken

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

// GenerateToken ....
func GenerateToken(url string, secret string, expires int) string {
	now := time.Now().Unix()
	st := fmt.Sprintf("%s|%d|%d", url, now, expires)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(st))
	b64 := base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))

	r := strings.NewReplacer("/", "_", "=", "", "+", "-")
	token := r.Replace(b64)

	return token
}
