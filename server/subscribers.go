package server

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Subscriber struct {
	ID        int64     `json:"id"`
	Topic     string    `json:"topic"`
	Nickname  string    `json:"nickname"`
	IP        string    `json:"ip"`
	Device    string    `json:"device"`
	CreatedAt time.Time `json:"created_at"`
}

type SubscriberStore struct {
	db *sql.DB
	mu sync.Mutex
}

func NewSubscriberStore(filename string) (*SubscriberStore, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open subscriber db: %w", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS subscribers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		topic TEXT NOT NULL,
		nickname TEXT,
		ip TEXT NOT NULL,
		device TEXT,
		created_at DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_subscribers_topic ON subscribers(topic);
	`

	if _, err := db.Exec(createTableSQL); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create subscribers table: %w", err)
	}

	return &SubscriberStore{db: db}, nil
}

func (s *SubscriberStore) Add(topic, nickname, ip, device string) (*Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	res, err := s.db.Exec(
		"INSERT INTO subscribers (topic, nickname, ip, device, created_at) VALUES (?, ?, ?, ?, ?)",
		topic, nickname, ip, device, now,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		ID:        id,
		Topic:     topic,
		Nickname:  nickname,
		IP:        ip,
		Device:    device,
		CreatedAt: now,
	}, nil
}

func (s *SubscriberStore) ListAll() ([]*Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	rows, err := s.db.Query("SELECT id, topic, COALESCE(nickname, ''), ip, COALESCE(device, ''), created_at FROM subscribers ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Subscriber
	for rows.Next() {
		var sub Subscriber
		if err := rows.Scan(&sub.ID, &sub.Topic, &sub.Nickname, &sub.IP, &sub.Device, &sub.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &sub)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SubscriberStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
