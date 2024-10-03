package wooricard_test

import (
	"testing"

	"github.com/kangseokgyu/auto-gagebu/internal/pkg/wooricard"
)

func TestFetch(t *testing.T) {
	normal_msg := "우리(0000)승인\n강*규님\n999,999원 일시불\n12/14 20:38\n누적999,000원\n스타벅스"
	_, e := wooricard.Fetch(normal_msg)
	if e != nil {
		t.Errorf("failed to parse message")
	}

	abnormal_msg := "스팸 메시지 입니다."
	_, e = wooricard.Fetch(abnormal_msg)
	if e == nil {
		t.Errorf("abnormal")
	}
}
