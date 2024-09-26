package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	msg    []byte
	sender string
}

func main() {
	fmt.Println("== Read Message ==")
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
	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.msg, &msg.sender)
		if err != nil {
			panic(err)
		}
		msgs := strings.SplitAfter(string(msg.msg), "NSNumber")[0]
		msgs = strings.SplitAfter(msgs, "NSString")[1]
		msgs = strings.Split(msgs, "NSDictionary")[0]
		fmt.Println(msgs[6 : len(msgs)-12])
		fmt.Println()
	}
}
