package repositories

import (
	"postly/configs"
	"postly/models"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (pr PostRepository) Store(newPost models.Post) error {
	_, err := configs.DB.Exec("insert into posts(user_id, content) values($1,$2)", newPost.UserID, newPost.Content)
	return err
}

func (pr PostRepository) GetPost(postID int) (models.Post, error) {
	var post models.Post
	err := configs.DB.QueryRow("select * from posts where id = $1", postID).Scan(
		&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
	return post, err
}

func (pr PostRepository) Update(post models.Post) (bool, error) {
	result, err := configs.DB.Exec("update posts set content = $1 where id = $2", post.Content, post.ID)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	return rows > 0, err
}

func (pr PostRepository) Delete(postID int) error {
	_, err := configs.DB.Exec("delete from posts where id = $1", postID)
	return err
}

func (pr PostRepository) GetPosts(userID int, postsType string) ([]models.PostResponse, error) {
	var posts []models.PostResponse
	query := `
		select
			p.id,
			p.content,
			p.created_at,
			u.id,
			u.username,
			u.firstname,
			u.lastname,
			count(distinct l.id) as likes,
			count(distinct c.id) as comments
		from posts p
		join users u on p.user_id = u.id
		left join likes l on l.post_id = p.id
		left join comments c on c.post_id = p.id`

	switch postsType {
	case "feed":
		query += ` where p.user_id in (
				select follow_id from follows where user_id = $1
			)`
	case "explore":
		query += ` where p.user_id != $1
			    and p.user_id not in (
				select follow_id from follows where user_id = $1
			)`
	default:
		query += ` where p.user_id = $1`
	}

	query += ` group by p.id, u.id
		order by p.created_at desc
		`

	rows, err := configs.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.PostResponse

		if err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.CreatedAt,
			&post.User.ID,
			&post.User.Username,
			&post.User.Firstname,
			&post.User.Lastname,
			&post.Likes,
			&post.Comments,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
