package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	commonPkg "go-park-mail-ru/2022_2_BugOverload/pkg"
)

var secret []byte

type TokenData struct {
	SessionID string
	User      models.User
	Exp       int64
}

// init prepare secret
func init() {
	var err error
	secret, err = commonPkg.GenerateRandomBytes(pkg.CsrfSecretLength)
	if err != nil {
		secret = []byte(pkg.CsrfSecretDefault)
	}
}

// CreateCsrfToken create CSRF token using aes crypt
func CreateCsrfToken(session *models.Session) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", errors.ErrCsrfTokenCreate
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.ErrCsrfTokenCreate
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.ErrCsrfTokenCreate
	}

	td := &TokenData{SessionID: session.ID, User: *session.User, Exp: time.Now().Add(time.Hour).Unix()}
	data, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)
	return token, nil
}

// CheckCsrfToken return true if CSRF token correct, false otherwise
func CheckCsrfToken(session *models.Session, inputToken string) (bool, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return false, stdErrors.WithMessagef(errors.ErrCsrfTokenCheckInternal,
			"csrf aes.NewCipher error, input token: [%s]",
			inputToken)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, stdErrors.WithMessagef(errors.ErrCsrfTokenCheckInternal,
			"csrf cipher.NewGCM error, input token: [%s]",
			inputToken)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(inputToken)
	if err != nil {
		return false, stdErrors.WithMessagef(errors.ErrCsrfTokenCheckInternal,
			"csrf decode base64 input token error, input token: [%s]",
			inputToken)
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, stdErrors.WithMessagef(errors.ErrCsrfTokenCheckInternal,
			"csrf noncesize error, cipher text len = [%d], noncesize = [%d], input token: [%s]",
			len(ciphertext), nonceSize, inputToken)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, stdErrors.WithMessagef(errors.ErrCsrfTokenCheckInternal,
			"csrf error with decode token(aesgcm.Open), inputToken: [%s]",
			inputToken)
	}

	td := TokenData{}
	err = json.Unmarshal(plaintext, &td)
	if err != nil {
		return false, errors.ErrCsrfTokenCheck
	}

	if td.Exp < time.Now().Unix() {
		return false, errors.ErrCsrfTokenExpired
	}

	return session.ID == td.SessionID && *session.User == td.User, nil
}
