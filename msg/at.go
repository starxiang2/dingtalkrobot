package msg

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

func (a *At) SetAtMobiles(mobiles []string) {
	a.AtMobiles = mobiles
}

func (a *At) SetAtUserIds(userIds []string) {
	a.AtUserIds = userIds
}

func (a *At) AtAll() {
	a.IsAtAll = true
}
