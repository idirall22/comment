package provider

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/idirall22/comment/models"
)

type postgresProvider struct {
	db        *sql.DB
	tableName string
}

// New create a comment
func (p *postgresProvider) New(ctx context.Context, content string, userID int64, groupID int64) (*models.Comment, error) {

	query := fmt.Sprintf(`
        INSERT INTO %s
        (content, user_id, group_id)
        VALUES (%s, %d, %d) RETURNING id, created_at`,
		p.tableName, content, userID, groupID)

	comment := &models.Comment{}
	err := p.db.QueryRowContext(ctx, query).Scan(comment.ID, comment.CreatedAt)

	if err != nil {
		errOut := parseError(err)
		return nil, errOut
	}

	return comment, nil
}

// Get get a comment
func (p *postgresProvider) List(ctx context.Context, postID int64, limit, offset uint) ([]*models.Comment, error) {

	query := fmt.Sprintf(`
		 SELECT c.id, c.content, c.user_id, c.post_id, c.created_at, u.avatar, u.username
		 FROM
		 (
			 SELECT id, content, user_id, post_id, created_at
			 FROM %s WHERE post_id= %d LIMIT %d OFFSET %d
		 ) c
		 LEFT JOIN users u ON u.id = c.user_id
		`, p.tableName, postID, limit, offset*limit)

	rows, err := p.db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	comments := []*models.Comment{}

	if rows.Next() {
		comment := &models.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.UserID,
			&comment.CreatedAt,
			&comment.PostID,
			&comment.AvatarURL,
			&comment.Username,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

// Update update a comment
func (p *postgresProvider) Update(ctx context.Context, id int64, content string) error {

	query := fmt.Sprintf(`UPDATE %s SET(content) VALUES(%s) WHERE id=%d`, p.tableName, content, id)

	_, err := p.db.ExecContext(ctx, query)

	if err != nil {
		return parseError(err)
	}
	return nil
}

// Delete delete a comment
func (p *postgresProvider) Delete(ctx context.Context, commentID int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=%d`, p.tableName, commentID)
	_, err := p.db.ExecContext(ctx, query)

	if err != nil {
		return parseError(err)
	}
	return nil
}
