package service

import (
	"time"

	"github.com/baowk/dilu-core/core"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

// SetStore 设置store
func SetStore(s base64Captcha.Store) {
	base64Captcha.DefaultMemStore = s
}

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return cap.Generate()
}

// Verify 校验验证码
func Verify(id, code string, clear bool) bool {
	return base64Captcha.DefaultMemStore.Verify(id, code, clear)
}

func NewCacheStore(expiration time.Duration) *CacheStore {
	return &CacheStore{
		expiration: expiration,
	}
}

type CacheStore struct {
	//sync.RWMutex
	expiration time.Duration
}

func (s *CacheStore) GetExpiration() time.Duration {
	if s.expiration == 0 {
		return time.Minute * 10
	}
	return s.expiration
}

func (s *CacheStore) Set(id string, value string) error {
	return core.Cache.Set("captcha:"+id, value, s.GetExpiration())
}

func (s *CacheStore) Get(id string, clear bool) (value string) {
	str, err := core.Cache.Get("captcha:" + id)
	if clear {
		defer core.Cache.Del("captcha:" + id)
	}
	if err != nil {
		return ""
	}
	return str
}

func (s *CacheStore) Verify(id, answer string, clear bool) bool {
	if answer == "" {
		return false
	}
	v := s.Get(id, clear)
	return v == answer
}
