package inmemory

import (
	"fmt"
	"posts/dbmodel"
	"sync"
	"time"
)

type InMemoryDB struct {
	posts         map[int32]dbmodel.Post
	comments      map[int32][]dbmodel.Comment
	nextPostID    int32
	nextCommentID int32
	mu            sync.Mutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		posts:         make(map[int32]dbmodel.Post),
		comments:      make(map[int32][]dbmodel.Comment),
		nextPostID:    1,
		nextCommentID: 1,
	}
}

func (db *InMemoryDB) CreatePost(post *dbmodel.Post) {
	db.mu.Lock()
	defer db.mu.Unlock()
	post.ID = db.nextPostID
	db.nextPostID++
	post.PublishedAt = time.Now()
	db.posts[post.ID] = *post
}

func (db *InMemoryDB) GetAllPosts() []dbmodel.Post {
	fmt.Println(1111111111111111)
	db.mu.Lock()
	defer db.mu.Unlock()
	posts := []dbmodel.Post{}

	if len(db.posts) == 0 {
		fmt.Println(1111111111111111)
		return posts
	}
	for _, post := range db.posts {
		posts = append(posts, post)
	}
	return posts
}

func (db *InMemoryDB) GetOnePost(id int32) (dbmodel.Post, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	post, exists := db.posts[id]
	return post, exists
}

func (db *InMemoryDB) CreateComment(comment *dbmodel.Comment) {
	db.mu.Lock()
	defer db.mu.Unlock()
	comment.ID = db.nextCommentID
	db.nextCommentID++
	comment.CreatedAt = time.Now()
	db.comments[comment.PostID] = append(db.comments[comment.PostID], *comment)
}

func (db *InMemoryDB) GetCommentsByPost(postID int32) []dbmodel.Comment {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.comments[postID]
}
