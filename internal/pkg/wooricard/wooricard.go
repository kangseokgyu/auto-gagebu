package wooricard

import (
	"database/sql"
	"errors"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Receipt은 사용 내역을 담는 구조체입니다.
// 구조는 notion 테이블과 동일하게 사용합니다.
type Receipt struct {
	// 내역
	title string
	// 사용자
	user string
	// 금액
	amount uint64
	// 날짜
	date string
	// 항목
	reason []string
	// 수입/지출
	transaction_type string
	// 결제수단
	payment_method string
}

type message struct {
	context []byte
	sender  string
}

func GetMessages() []string {
	path := "/Users/kangseokgyu/Library/Messages/chat.db"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT m.attributedBody, h.id AS id FROM message m INNER JOIN handle h ON m.handle_id = h.ROWID WHERE h.id = \"+8215889955\" ORDER BY m.date DESC")
	if err != nil {
		panic(err)
	}

	var messages []string
	for rows.Next() {
		var msg message
		err = rows.Scan(&msg.context, &msg.sender)
		if err != nil {
			panic(err)
		}
		msgs := strings.SplitAfter(string(msg.context), "NSNumber")[0]
		msgs = strings.SplitAfter(msgs, "NSString")[1]
		msgs = strings.Split(msgs, "NSDictionary")[0]
		// fmt.Println(msgs[6 : len(msgs)-12])
		// fmt.Println()
		messages = append(messages, msgs[6:len(msgs)-12])
	}

	return messages
}

func Fetch(msg string) (*Receipt, error) {
	r, e := regexp.Compile(`.*\n.*\n([0-9,]*)원 .*\n([0-9]{1,2}\/[0-9]{1,2}).*\n.*\n(.*)`)
	if e != nil {
		return nil, errors.New("failed to compile regex")
	}

	s := r.FindStringSubmatch(msg)
	if len(s) == 0 {
		return nil, errors.New("failed to match regex")
	}
	amount, err := strconv.ParseUint(strings.ReplaceAll(s[1], ",", ""), 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse amount")
	}

	return &Receipt{
			title:            s[3],
			user:             "kangseokgyu",
			amount:           amount,
			date:             s[2],
			transaction_type: "지출",
			payment_method:   "신용카드"},
		nil
}

func GetReceipts() []Receipt {
	var receipts []Receipt

	messages := GetMessages()
	for _, msg := range messages {
		if r, err := Fetch(msg); err == nil {
			receipts = append(receipts, *r)
		}
	}

	return receipts
}
