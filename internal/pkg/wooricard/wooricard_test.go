package wooricard

import (
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	normal_msg := message{
		context: "우리(0000)승인\n강*규님\n999,999원 일시불\n12/14 20:38\n누적999,000원\n스타벅스",
		sender:  "+8215889955",
		date:    time.Now().UnixNano(),
	}

	_, e := Fetch(normal_msg)
	if e != nil {
		t.Errorf("failed to parse message")
	}

	abnormal_msg := message{
		context: "스팸 메시지 입니다.",
		sender:  "+8215889955",
		date:    time.Now().UnixNano(),
	}
	_, e = Fetch(abnormal_msg)
	if e == nil {
		t.Errorf("abnormal")
	}
}
