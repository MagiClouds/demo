package main

type ContentMsg struct {
    EffectId int         `json:"effect_id"`
    Text     string      `json:"text"`
    Data     interface{} `json:"data"`
}

//复制代码通过new一个对象，或者利用Golang本身的&方式来生成一个对象并返回一个对象指针：
func NewContentMsg(data, effectId int) *ContentMsg {
    instance := new(ContentMsg)
    instance.Data = data
    instance.EffectId = effectId
    return instance
}

func NewContentMsgV2(data, effectId int) *ContentMsg {
    return &ContentMsg{
        Data:     data,
        EffectId: effectId,
    }
}
