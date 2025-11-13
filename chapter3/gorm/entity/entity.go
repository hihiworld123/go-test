package entity

import "gorm.io/gorm"

type User struct {
	Id      int    `gorm:"primary_key"`
	Name    string `gorm:"size:255"`
	Posts   []Post `gorm:"foreignkey:UserId"`
	PostNum int
}

type Post struct {
	Id            int       `gorm:"primary_key"`
	Title         string    `gorm:"size:255"`
	Body          string    `gorm:"size:255"`
	UserId        int       `gorm:"index"`
	Comments      []Comment `gorm:"foreignkey:PostId"`
	CommentStatus string
}

type Comment struct {
	Id      int    `gorm:"primary_key"`
	Content string `gorm:"size:255"`
	PostId  int    `gorm:"index"`
	Post    *Post  `gorm:"foreignkey:PostId"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	err := tx.Exec("UPDATE users SET post_num = post_num +1 WHERE id = ?", p.UserId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {

	cnt := int64(0)
	err := tx.Model(c).Where("post_id = ?", c.PostId).Count(&cnt).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if cnt == 0 {
		err = tx.Exec("UPDATE posts SET comment_status ='无评论' WHERE id =?", c.PostId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	return nil
}
