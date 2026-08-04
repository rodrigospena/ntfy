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

type TopicItem struct {
	Name         string    `json:"name"`
	DisplayName  string    `json:"display_name"`
	DefaultIcon  string    `json:"default_icon"`
	DefaultImage string    `json:"default_image"`
	CreatedAt    time.Time `json:"created_at"`
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

	CREATE TABLE IF NOT EXISTS custom_topics (
		name TEXT PRIMARY KEY,
		display_name TEXT,
		default_icon TEXT,
		default_image TEXT,
		created_at DATETIME NOT NULL
	);
	`

	if _, err := db.Exec(createTableSQL); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create subscribers/topics tables: %w", err)
	}

	store := &SubscriberStore{db: db}
	_ = store.seedDefaultTopic()

	return store, nil
}

func (s *SubscriberStore) seedDefaultTopic() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM custom_topics").Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		now := time.Now()
		_, err = s.db.Exec(
			"INSERT INTO custom_topics (name, display_name, default_icon, default_image, created_at) VALUES (?, ?, ?, ?, ?)",
			"cafecomfamilia", "Café com Família", "", "", now,
		)
	}
	return err
}

func (s *SubscriberStore) AddSubscriber(topic, nickname, ip, device string) (*Subscriber, error) {
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

func (s *SubscriberStore) ListAllSubscribers() ([]*Subscriber, error) {
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

func (s *SubscriberStore) AddTopic(name, displayName, defaultIcon, defaultImage string) (*TopicItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	if displayName == "" {
		displayName = name
	}
	_, err := s.db.Exec(
		"INSERT INTO custom_topics (name, display_name, default_icon, default_image, created_at) VALUES (?, ?, ?, ?, ?)",
		name, displayName, defaultIcon, defaultImage, now,
	)
	if err != nil {
		return nil, err
	}

	return &TopicItem{
		Name:         name,
		DisplayName:  displayName,
		DefaultIcon:  defaultIcon,
		DefaultImage: defaultImage,
		CreatedAt:    now,
	}, nil
}

func (s *SubscriberStore) UpdateTopic(name, displayName, defaultIcon, defaultImage string) (*TopicItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if displayName == "" {
		displayName = name
	}
	_, err := s.db.Exec(
		"UPDATE custom_topics SET display_name = ?, default_icon = ?, default_image = ? WHERE name = ?",
		displayName, defaultIcon, defaultImage, name,
	)
	if err != nil {
		return nil, err
	}

	return &TopicItem{
		Name:         name,
		DisplayName:  displayName,
		DefaultIcon:  defaultIcon,
		DefaultImage: defaultImage,
	}, nil
}

func (s *SubscriberStore) ListTopics() ([]*TopicItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	rows, err := s.db.Query("SELECT name, COALESCE(display_name, ''), COALESCE(default_icon, ''), COALESCE(default_image, ''), created_at FROM custom_topics ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*TopicItem
	for rows.Next() {
		var item TopicItem
		if err := rows.Scan(&item.Name, &item.DisplayName, &item.DefaultIcon, &item.DefaultImage, &item.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SubscriberStore) DeleteTopic(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec("DELETE FROM custom_topics WHERE name = ?", name)
	return err
}

func (s *SubscriberStore) UpdateSubscriber(id int64, nickname, device, topic string) (*Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(
		"UPDATE subscribers SET nickname = ?, device = ?, topic = ? WHERE id = ?",
		nickname, device, topic, id,
	)
	if err != nil {
		return nil, err
	}

	var sub Subscriber
	err = s.db.QueryRow("SELECT id, topic, COALESCE(nickname, ''), ip, COALESCE(device, ''), created_at FROM subscribers WHERE id = ?", id).Scan(
		&sub.ID, &sub.Topic, &sub.Nickname, &sub.IP, &sub.Device, &sub.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (s *SubscriberStore) DeleteSubscriber(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec("DELETE FROM subscribers WHERE id = ?", id)
	return err
}

func (s *SubscriberStore) DeleteSubscribersByTopic(topic string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec("DELETE FROM subscribers WHERE topic = ?", topic)
	return err
}

func (s *SubscriberStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
