package provider

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/idirall22/comment/models"
)

// PostgresProvider structure
type PostgresProvider struct {
	DB        *sql.DB
	TableName string
}

// New create a comment
func (p *PostgresProvider) New(ctx context.Context, content string, userID int64, postID int64) (*models.Comment, error) {

	query := fmt.Sprintf(`
        INSERT INTO %s
        (content, user_id, post_id)
        VALUES ('%s', %d, %d) RETURNING id, created_at`,
		p.TableName, content, userID, postID)

	stmt, err := p.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	comment := &models.Comment{}
	err = stmt.QueryRowContext(ctx).Scan(&comment.ID, &comment.CreatedAt)

	if err != nil {
		errOut := parseError(err)
		return nil, errOut
	}

	comment.Content = content
	comment.UserID = userID
	comment.UserID = postID
	return comment, nil
}

// List get a comment
func (p *PostgresProvider) List(ctx context.Context, postID int64, limit, offset uint) ([]*models.Comment, error) {

	query := fmt.Sprintf(`
		 SELECT
		 c.id, c.content, c.user_id, c.post_id, c.created_at, u.avatar, u.username
		 FROM
		 (
			 SELECT id, content, user_id, post_id, created_at
			 FROM %s WHERE post_id= %d LIMIT %d OFFSET %d
		 ) c
		 LEFT JOIN users u ON u.id = c.user_id
		`, p.TableName, postID, limit, offset*limit)

	stmt, err := p.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	comments := []*models.Comment{}

	for rows.Next() {
		comment := &models.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.UserID,
			&comment.PostID,
			&comment.CreatedAt,
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
func (p *PostgresProvider) Update(ctx context.Context, userID, id int64, content string) (*models.Comment, error) {

	query := fmt.Sprintf(`
		UPDATE %s SET content='%s' WHERE user_id=%d AND id=%d
		RETURNING id, content, post_id
		`, p.TableName, content, userID, id)

	stmt, err := p.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	comment := &models.Comment{}
	err = stmt.QueryRowContext(ctx).Scan(
		&comment.ID,
		&comment.Content,
		&comment.PostID,
	)

	if err != nil {
		return nil, parseError(err)
	}
	return comment, nil
}

// Delete delete a comment
func (p *PostgresProvider) Delete(ctx context.Context, userID, commentID int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=%d AND user_id=%d`, p.TableName, commentID, userID)

	stmt, err := p.DB.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx)

	if err != nil {
		return parseError(err)
	}
	return nil
}
