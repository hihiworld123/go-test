package main

import (
	"go-test/chapter3/gorm/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//user, err := GetUserPostAndComment(1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(user)

	//data, err := GetMaxCommentNumPost()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)

	//post := entity.Post{
	//	Title:  "test",
	//	Body:   "test",
	//	UserId: 1,
	//}
	//err := CreatePost(post)
	//if err != nil {
	//	fmt.Println(err)
	//}

	DeleteComment(2)

}

var db *gorm.DB

func init() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动创建和迁移表
	err = db.AutoMigrate(&entity.User{})
	// 这里的db就是上述gorm.open的db哈
	if err != nil {
		panic("User 创建/迁移表格失败, error = " + err.Error())
	}

	err = db.AutoMigrate(&entity.Post{})
	// 这里的db就是上述gorm.open的db哈
	if err != nil {
		panic("Post 创建/迁移表格失败, error = " + err.Error())
	}

	err = db.AutoMigrate(&entity.Comment{})
	// 这里的db就是上述gorm.open的db哈
	if err != nil {
		panic("Comment 创建/迁移表格失败, error = " + err.Error())
	}
}

func GetUserPostAndComment(userId int) ([]entity.Post, error) {
	var posts []entity.Post
	err := db.Where("user_id =?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	for i := range posts {
		err = db.Model(&posts[i]).Association("Comments").Find(&posts[i].Comments)
		if err != nil {
			return posts, err
		}
	}
	//var posts []entity.Post
	//err = db.Model(&user).Association("Posts").Find(&posts)

	return posts, err
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息
func GetMaxCommentNumPost() (entity.Post, error) {

	data := []map[string]any{}
	scan := db.Raw("SELECT COUNT(1) cnt,t.post_id FROM comments t GROUP BY t.post_id").Scan(&data)
	if scan.Error != nil {
		return entity.Post{}, scan.Error
	}

	maxCnt := int64(0)
	maxId := int64(0)
	for _, datum := range data {
		cnt, ok := datum["cnt"]
		if ok && cnt != nil {
			cnt1, ok1 := cnt.(int64)
			if ok1 && maxCnt < cnt1 {
				maxCnt = cnt1
				maxId = datum["post_id"].(int64)
			}
		}
	}

	post := entity.Post{}
	tx := db.Where("id = ?", maxId).First(&post)
	if tx.Error != nil {
		return post, tx.Error
	}

	err := db.Model(&post).Association("Comments").Find(&post.Comments)
	if err != nil {
		return post, err
	}

	return post, nil
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func CreatePost(post entity.Post) error {
	err := db.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
// 如果评论数量为 0，则更新文章的评论状态为 "无评论"
func DeleteComment(postId int) error {

	tx := db.Delete(&entity.Comment{PostId: postId}, "post_id = ?", postId)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
