package event

import (
	"testing"
)

func TestClient_PostEvent(t *testing.T) {
	c, err := NewClient(Config{
		Url:       "https://app.wizardcloud.cn/a/app-033f5501e4000e4b/v1/events",
		AppId:     "033f5501e4000e4b",
		AppSecret: "ZBWRGW2n97nn/5n7CeLxhdSV3/POSDlpRQ/cMVIGxhQ=",
	})
	if nil != err {
		t.Fatal(err)
	}

	t.Log(c.PostEvent("whoru", nil, "application/json"))
}
