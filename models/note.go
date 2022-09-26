package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	UserID    uint64 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll(user *User) *[]Note {
	var notes []Note
	DB.Where("deleted_at is NULL and user_id = ?", user.ID).Order("updated_at desc").Find(&notes)
	return &notes
}

func NoteCreate(user *User, name string, content string) *Note {
	entry := Note{Name: name, Content: content, UserID: user.ID}
	DB.Create(&entry)
	return &entry
}

func NotesFind(user *User, id uint64) *Note {
	var note Note
	DB.Where("id = ? and user_id = ?", id, user.ID).First(&note)
	return &note
}

func (note *Note) Update(name string, content string) {
	note.Name = name
	note.Content = content
	DB.Save(note)
}

func NotesMarkDelete(user *User, id uint64) {
	// UPDATE notes SET deleted_at=<Current Time> WHERE id = <id> and user_id = <user_id>
	DB.Where("id = ? and user_id = ?", id, user.ID).Delete(&Note{})
}
