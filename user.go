package monketype

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"

	"encoding/json"
	"time"
)

type User struct {
	ID                string `json:"id"                    db:"id"`
	Email             string `json:"email"                 db:"email"`
	Nick              string `json:"nick"                  db:"nick"`
	Bio               string `json:"bio"                   db:"bio"`
	SubscriberCount   int    `json:"subscriber_count"      db:"subscriber_count"`
	SubscriptionCount int    `json:"subscription_count"    db:"subscription_count"`
	PostCount         int    `json:"post_count"            db:"post_count"`
	Created           int64  `json:"created"               db:"created"`
}

func (_ User) FromMap(data map[string]interface{}) (it MonkeType, err error) {
	var config mapstructure.DecoderConfig = mapstructure.DecoderConfig{
		Metadata: nil,
		TagName:  "json",
		Result:   &it,
	}

	var decoder *mapstructure.Decoder
	if decoder, err = mapstructure.NewDecoder(&config); err == nil {
		err = decoder.Decode(data)
	}

	return
}

func (user User) Map() (data map[string]interface{}, err error) {
	var bytes []byte
	if bytes, err = json.Marshal(user); err == nil {
		err = json.Unmarshal(bytes, &data)
	}

	return
}

func (user User) JSON() (data []byte, err error) {
	data, err = json.Marshal(user)
	return
}

func NewUser(nick, bio, email string) (made User) {
	made = User{
		Nick:  nick,
		Email: email,
		Bio:   bio,

		ID:      uuid.New().String(),
		Created: time.Now().Unix(),
	}

	return
}
