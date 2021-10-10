package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

//本系統中的实体用到一对一，一对多，多对多关系
//多条件关联，单条件管理，自关联

type Module struct {
	Id        uint           `gorm:"column:id;primary_key;AUTO_INCREMENT" db:"id" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

//浏览历史
type Browsed struct {
	Module
	UserId    uint `gorm:"column:user_id;" db:"user_id" json:"user_id"`
	MediaId   uint `gorm:"column:media_id;" db:"media_id" json:"media_id"`
	MediaType uint `gorm:"column:media_type;type:int(6)" db:"media_type" json:"media_type"`
}
type Collection struct {
	Module
	UserId    uint `gorm:"column:user_id;" db:"user_id" json:"user_id"`
	MediaId   uint `gorm:"column:media_id;" db:"media_id" json:"media_id"`
	MediaType uint `gorm:"column:media_type;type:int(6)" db:"media_type" json:"media_type"`
}
type Comment struct {
	Module
	MediaId    uint   `gorm:"column:media_id;" db:"media_id" json:"media_id"`
	MediaType  uint   `gorm:"column:media_type;type:int(6)" db:"media_type" json:"media_type"`
	ReleaserId uint   `gorm:"column:releaser_id;" db:"releaser_id" json:"releaser_id"`
	Content    string `gorm:"column:content" db:"content" json:"content"`
	LikeNum    uint   `gorm:"column:like_num;type:int(11)" db:"like_num" json:"like_num"`

	Releaser  *User   `gorm:"foreignKey:ReleaserId" json:"releaser"`
	LikeUsers []*User `gorm:"many2many:comment_user" json:"like_users"`
}

//type Follow struct {
//	Module
//	UserId   int `gorm:"column:user_id" db:"user_id" json:"user_id"`
//	FollowId int `gorm:"column:follow_id" db:"follow_id" json:"follow_id"`
//}
type Media interface {
	Like()
	//Collection()
	//Browse()
}
type CommMedia struct {
}

func (this *CommMedia) Like() {
	fmt.Println("点赞")
}

type Label struct {
	Module
	LabelName   string `gorm:"column:label_name;type:varchar(255)" db:"label_name" json:"label_name"`
	LabelType   string `gorm:"column:label_type;type:varchar(255)" db:"label_type" json:"label_type"` //标签类型，分区标签，活动标签，普通标签等
	LabelRouter uint   `gorm:"column:label_router;" db:"label_router" json:"label_router"`            //标签绑定分区或活动ID
	//使用了该标签的媒体
	//Medias []*Media `gorm:"many2many:label_media" json:"medias"`
}

//使用了某标签的媒体
type LabelMedia struct {
	Module
	LabelId   uint `gorm:"column:label_id;" db:"label_id" json:"label_id"`                  //标签ID
	MediaId   uint `gorm:"column:media_id;" db:"media_id" json:"media_id"`                  //媒体ID
	MediaType uint `gorm:"column:media_type;type:int(6)" db:"media_type" json:"media_type"` //媒体类型
}
type Like struct {
	Module
	UserId    uint `gorm:"column:user_id;" db:"user_id" json:"user_id"`
	MediaId   uint `gorm:"column:media_id;" db:"media_id" json:"media_id"`
	MediaType uint `gorm:"column:media_type;type:int(6)" db:"media_type" json:"media_type"`
}

type Message struct {
	Module
	CommMedia
	PartId      uint   `gorm:"column:part_id;" db:"part_id" json:"part_id"`
	Title       string `gorm:"column:title;type:varchar(255)" db:"title" json:"title"`
	CoverImgUrl string `gorm:"column:cover_img_url;type:varchar(255)" db:"cover_img_url" json:"cover_img_url"`
	Content     string `gorm:"column:content" db:"content" json:"content"`
	ReleaserId  uint   `gorm:"column:releaser_id;" db:"releaser_id" json:"releaser_id"`

	Part     *Partition `gorm:"foreignKey:PartId" json:"part"`
	Releaser *User      `gorm:"foreignKey:ReleaserId" json:"releaser"`
}

type Music struct {
	Module
	CommMedia
	PartId     uint   `gorm:"column:part_id;" db:"part_id" json:"part_id"`
	Name       string `gorm:"column:name;type:varchar(255)" db:"name" json:"name"`
	ReleaserId uint   `gorm:"column:releaser_id;" db:"releaser_id" json:"releaser_id"`
	CoverUrl   string `gorm:"column:cover_url;type:varchar(255)" db:"cover_url" json:"cover_url"`
	ContentUrl string `gorm:"column:content_url;type:varchar(255)" db:"content_url" json:"content_url"`

	Part     *Partition `gorm:"foreignKey:PartId" json:"part"`
	Releaser *User      `gorm:"foreignKey:ReleaserId" json:"releaser"`
}

// 用户关注媒体分区中间表
//type PartUser struct {
//	Module
//	PartId int `gorm:"column:part_id" db:"part_id" json:"part_id"`
//	UserId int `gorm:"column:user_id" db:"user_id" json:"user_id"`
//}
type Partition struct {
	Module
	PartTypeId uint   `gorm:"column:part_type_id;" db:"part_type_id" json:"part_type_id"`
	PartTitle  string `gorm:"column:part_title;type:varchar(255)" db:"part_title" json:"part_title"`
	PartDesc   string `gorm:"column:part_desc;type:varchar(255)" db:"part_desc" json:"part_desc"`
	CoverUrl   string `gorm:"column:cover_url;type:varchar(255)" db:"cover_url" json:"cover_url"`

	PartitionType *PartitionType `gorm:"foreignKey:PartTypeId" json:"partition_type"`
	Musics        []*Music       `gorm:"foreignKey:PartId" json:"musics"`
	Videos        []*Video       `gorm:"foreignKey:PartId" json:"videos"`
	Posters       []*Poster      `gorm:"foreignKey:PartId" json:"posters"`
	Messages      []*Message     `gorm:"foreignKey:PartId" json:"messages"`
	//关注该分区的用户
	Users []*User `gorm:"many2many:part_user" json:"users"`
}
type PartitionType struct {
	Module
	PartType string `gorm:"column:part_type;type:varchar(255)" db:"part_type" json:"part_type"`
}
type Poster struct {
	Module
	CommMedia
	PartId     uint   `gorm:"column:part_id;" db:"part_id" json:"part_id"`
	Content    string `gorm:"column:content" db:"content" json:"content"`
	ImgsUrl    string `gorm:"column:imgs_url;type:varchar(255)" db:"imgs_url" json:"imgs_url"`
	ReleaserId uint   `gorm:"column:releaser_id;" db:"releaser_id" json:"releaser_id"`

	Part     *Partition `gorm:"foreignKey:PartId" json:"part"`
	Releaser *User      `gorm:"foreignKey:ReleaserId" json:"releaser"`
}

type Video struct {
	Module
	CommMedia
	PartId     uint   `gorm:"column:part_id;" db:"part_id" json:"part_id"`
	Title      string `gorm:"column:title;type:varchar(255)" db:"title" json:"title"`
	CoverUrl   string `gorm:"column:cover_url;type:varchar(255)" db:"cover_url" json:"cover_url"`
	ContentUrl string `gorm:"column:content_url;type:varchar(255)" db:"content_url" json:"content_url"`
	ReleaserId uint   `gorm:"column:releaser_id;" db:"releaser_id" json:"releaser_id"`
	Resume     string `gorm:"column:resume;type:varchar(300)" db:"resume" json:"resume"`

	Part     *Partition `gorm:"foreignKey:PartId" json:"part"`
	Releaser *User      `gorm:"foreignKey:ReleaserId" json:"releaser"`
}
type Title struct {
	Module
	Title     string `gorm:"column:title;type:varchar(255)" db:"title" json:"title"`
	SoulMoney int    `gorm:"column:soul_money;type:int(11)" db:"soul_money" json:"soul_money"`
	Desc      string `gorm:"column:desc;type:varchar(255)" db:"desc" json:"desc"`
	ImgUrl    string `gorm:"column:img_url;type:varchar(255)" db:"img_url" json:"img_url"`

	Users []*User `gorm:"many2many:title_user;" json:"users"`
}

// 用户头衔库
//type TitleUser struct {
//	Module
//	Users  []User  `gorm:"column:user_id;many2many:t_title_user;" db:"user_id" json:"user_id"`
//	Titles []Title `gorm:"column:title_id" db:"title_id" json:"title_id"`
//}
type User struct {
	Module
	Username     string    `gorm:"column:username;type:varchar(255)" db:"username" json:"username"`
	Password     string    `gorm:"column:password;type:varchar(255)" db:"password" json:"password"`
	HeadImgUrl   string    `gorm:"column:head_img_url;type:varchar(255)" db:"head_img_url" json:"head_img_url"`
	TopImgUrl    string    `gorm:"column:top_img_url;type:varchar(255)" db:"top_img_url" json:"top_img_url"`
	Signature    string    `gorm:"column:signature;type:varchar(255)" db:"signature" json:"signature"`
	Email        string    `gorm:"column:email;type:varchar(255)" db:"email" json:"email"`
	Birthday     time.Time `gorm:"column:birthday" db:"birthday" json:"birthday"`
	State        int       `gorm:"column:state;type:int(4)" db:"state" json:"state"` //状态
	Root         int       `gorm:"column:root;type:int(4)" db:"root" json:"root"`    //管理员标识
	SoulCurrency int       `gorm:"column:soul_currency;int(11)" db:"soul_currency" json:"soul_currency"`
	SoulTitleId  uint      `gorm:"column:soul_title_id;" db:"soul_title_id" json:"soul_title_id"` //当前头衔ID

	SoulTitle    *Title       `gorm:"foreignKey:SoulTitleId" json:"soul_title"`
	Titles       []*Title     `gorm:"many2many:title_user" json:"titles"`
	Partitions   []*Partition `gorm:"many2many:part_user" json:"partitions"`
	LikeComments []*Comment   `gorm:"many2many:comment_user" json:"like_comments"`
	//自关联
	Friends []*User `gorm:"many2many:friends;joinForeignKey:friend_id"  json:"friends"`
}
