package inmemory

import (
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
	post.CommentsAllowed = true
	db.posts[post.ID] = *post
}

func (db *InMemoryDB) GetAllPosts() []dbmodel.Post {
	db.mu.Lock()
	defer db.mu.Unlock()
	posts := make([]dbmodel.Post, 0, len(db.posts))
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

func (db *InMemoryDB) UpdatePost(id int32, updatedPost dbmodel.Post) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	post, exists := db.posts[id]
	if !exists {
		return false
	}
	if updatedPost.Title != "" {
		post.Title = updatedPost.Title
	}
	if updatedPost.Content != "" {
		post.Content = updatedPost.Content
	}
	if updatedPost.Author != "" {
		post.Author = updatedPost.Author
	}
	if updatedPost.UpdatedAt != nil {
		post.UpdatedAt = updatedPost.UpdatedAt
	}
	if updatedPost.CommentsAllowed != post.CommentsAllowed {
		post.CommentsAllowed = updatedPost.CommentsAllowed
	}
	db.posts[id] = post
	return true
}

func (db *InMemoryDB) CreateComment(comment *dbmodel.Comment) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	post, exists := db.posts[comment.PostID]
	if !exists || !post.CommentsAllowed {
		return false
	}
	comment.ID = db.nextCommentID
	db.nextCommentID++
	comment.CreatedAt = time.Now()
	db.comments[comment.PostID] = append(db.comments[comment.PostID], *comment)
	return true
}

func (db *InMemoryDB) GetCommentsByPost(postID int32) []dbmodel.Comment {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.comments[postID]
}
