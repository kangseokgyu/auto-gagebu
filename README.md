# auto-gagebu
결제 내역을 자동으로 파싱하여 가계부에 작성하는 프로젝트

## 1. woori-parser
우리카드 결제 메시지 가져오기
iMessage는 `/Users/[Username]/Library/Messages/chat.db`에 저장된다.
메시지는 `message` 테이블이 저장되고, 발신자 정보는 `handle` 테이블에 저장된다.

```
sudo go run cmd/woori-parser/woori-parser.go
```

### 1.1 message table 구조
메시지 내용은 text 컬럼이나 attributedBody 컬럼에 작성되어 있다.
발신자 정보는 handle_id를 통해 알 수 있다.

우리카드 메시지는 attributedBody에 존재한다.
```
`sqlite3 chat.db ".schema message"
```
```sql
CREATE TABLE message (ROWID INTEGER PRIMARY KEY AUTOINCREMENT,
                      guid TEXT UNIQUE NOT NULL,
                      text TEXT,
                      replace INTEGER DEFAULT 0,
                      service_center TEXT,
                      handle_id INTEGER DEFAULT 0,
                      subject TEXT,
                      country TEXT,
                      attributedBody BLOB,
                      version INTEGER DEFAULT 0,
                      type INTEGER DEFAULT 0,
                      service TEXT,
                      account TEXT,
                      account_guid TEXT,
                      error INTEGER DEFAULT 0,
                      date INTEGER,
                      date_read INTEGER,
                      date_delivered INTEGER,
                      is_delivered INTEGER DEFAULT 0,
                      is_finished INTEGER DEFAULT 0,
                      is_emote INTEGER DEFAULT 0,
                      is_from_me INTEGER DEFAULT 0,
                      is_empty INTEGER DEFAULT 0,
                      is_delayed INTEGER DEFAULT 0,
                      is_auto_reply INTEGER DEFAULT 0,
                      is_prepared INTEGER DEFAULT 0,
                      is_read INTEGER DEFAULT 0,
                      is_system_message INTEGER DEFAULT 0,
                      is_sent INTEGER DEFAULT 0,
                      has_dd_results INTEGER DEFAULT 0,
                      is_service_message INTEGER DEFAULT 0,
                      is_forward INTEGER DEFAULT 0,
                      was_downgraded INTEGER DEFAULT 0,
                      is_archive INTEGER DEFAULT 0,
                      cache_has_attachments INTEGER DEFAULT 0,
                      cache_roomnames TEXT,
                      was_data_detected INTEGER DEFAULT 0,
                      was_deduplicated INTEGER DEFAULT 0,
                      is_audio_message INTEGER DEFAULT 0,
                      is_played INTEGER DEFAULT 0,
                      date_played INTEGER,
                      item_type INTEGER DEFAULT 0,
                      other_handle INTEGER DEFAULT 0,
                      group_title TEXT,
                      group_action_type INTEGER DEFAULT 0,
                      share_status INTEGER DEFAULT 0,
                      share_direction INTEGER DEFAULT 0,
                      is_expirable INTEGER DEFAULT 0,
                      expire_state INTEGER DEFAULT 0,
                      message_action_type INTEGER DEFAULT 0,
                      message_source INTEGER DEFAULT 0,
                      associated_message_guid TEXT,
                      associated_message_type INTEGER DEFAULT 0,
                      balloon_bundle_id TEXT,
                      payload_data BLOB,
                      expressive_send_style_id TEXT,
                      associated_message_range_location INTEGER DEFAULT 0,
                      associated_message_range_length INTEGER DEFAULT 0,
                      time_expressive_send_played INTEGER,
                      message_summary_info BLOB,
                      ck_sync_state INTEGER DEFAULT 0,
                      ck_record_id TEXT,
                      ck_record_change_tag TEXT,
                      destination_caller_id TEXT,
                      sr_ck_sync_state INTEGER DEFAULT 0,
                      sr_ck_record_id TEXT,
                      sr_ck_record_change_tag TEXT,
                      is_corrupt INTEGER DEFAULT 0,
                      reply_to_guid TEXT,
                      sort_id INTEGER,
                      is_spam INTEGER DEFAULT 0,
                      has_unseen_mention INTEGER DEFAULT 0,
                      thread_originator_guid TEXT,
                      thread_originator_part TEXT,
                      syndication_ranges TEXT DEFAULT NULL,
                      was_delivered_quietly INTEGER DEFAULT 0,
                      did_notify_recipient INTEGER DEFAULT 0,
                      synced_syndication_ranges TEXT DEFAULT NULL,
                      date_retracted INTEGER DEFAULT 0,
                      date_edited INTEGER DEFAULT 0,
                      was_detonated INTEGER DEFAULT 0,
                      part_count INTEGER,
                      is_stewie INTEGER DEFAULT 0,
                      is_kt_verified INTEGER DEFAULT 0,
                      is_sos INTEGER DEFAULT 0,
                      is_critical INTEGER DEFAULT 0,
                      bia_reference_id TEXT DEFAULT NULL,
                      fallback_hash TEXT DEFAULT NULL,
                      associated_message_emoji TEXT DEFAULT NULL,
                      is_pending_satellite_send INTEGER DEFAULT 0,
                      needs_relay INTEGER DEFAULT 0,
                      schedule_type INTEGER DEFAULT 0,
                      schedule_state INTEGER DEFAULT 0,
                      sent_or_received_off_grid INTEGER DEFAULT 0);
```
### 1.2 handle table 구조
id 컬럼에 전화번호, email로 정보가 저장된다.
```bash
sqlite3 chat.db ".schema handle"
```
```sql
CREATE TABLE handle (ROWID INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
                     id TEXT NOT NULL,
                     country TEXT,
                     service TEXT NOT NULL,
                     uncanonicalized_id TEXT,
                     person_centric_id TEXT,
                     UNIQUE (id, service) );
```
