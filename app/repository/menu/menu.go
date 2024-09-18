package menu

import (
	"chi-onion-arch/app/model"
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type MenuRepository interface {
	List(ctx context.Context) ([]*model.Menu, error)
	Detail(ctx context.Context, id uuid.UUID) (*model.Menu, error)
	Insert(ctx context.Context, menu model.Menu) (*model.Menu, error)
	Update(ctx context.Context, datas map[string]interface{}, id uuid.UUID) error
	CheckDuplicateNameOrPath(ctx context.Context, name string, path string, excludeID uuid.UUID) (*model.Menu, error)
}

type menuRepository struct {
	Conn *sql.DB
}

func NewMenuRepository(Conn *sql.DB) MenuRepository {
	return &menuRepository{Conn}
}

func (r *menuRepository) List(ctx context.Context) ([]*model.Menu, error) {
	query := `SELECT id, name, path, created_at, created_by, updated_at, updated_by FROM menus`
	rows, err := r.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*model.Menu
	for rows.Next() {
		var menu model.Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Path,
			&menu.CreatedAt,
			&menu.CreatedBy,
			&menu.UpdatedAt,
			&menu.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}

	return menus, nil
}

func (r *menuRepository) Detail(ctx context.Context, id uuid.UUID) (*model.Menu, error) {
	var menu model.Menu
	query := `SELECT id, name, path, created_at, created_by, updated_at, updated_by FROM menus WHERE id = $1`
	err := r.Conn.QueryRowContext(ctx, query, id).Scan(
		&menu.ID,
		&menu.Name,
		&menu.Path,
		&menu.CreatedAt,
		&menu.CreatedBy,
		&menu.UpdatedAt,
		&menu.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Insert(ctx context.Context, menu model.Menu) (*model.Menu, error) {
	query := `INSERT INTO menus (id, name, path, created_at, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.Conn.QueryRowContext(ctx, query, menu.ID, menu.Name, menu.Path, menu.CreatedAt, menu.CreatedBy).Scan(&menu.ID)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Update(ctx context.Context, datas map[string]interface{}, id uuid.UUID) error {
	query := `UPDATE menus SET`
	var data []interface{}
	i := 1
	for key, value := range datas {
		query = query + " " + key + ` = $` + strconv.Itoa(i) + `,`
		data = append(data, value)
		i++
	}
	query = strings.TrimRight(query, ",") + " WHERE id = $" + strconv.Itoa(i)
	data = append(data, id)

	_, err := r.Conn.ExecContext(ctx, query, data...)
	return err
}

func (r *menuRepository) CheckDuplicateNameOrPath(ctx context.Context, name string, path string, excludeID uuid.UUID) (*model.Menu, error) {
	var menu model.Menu
	query := `SELECT id, name, path FROM menus WHERE (name = $1 OR path = $2) AND id != $3`
	err := r.Conn.QueryRowContext(ctx, query, name, path, excludeID).Scan(&menu.ID, &menu.Name, &menu.Path)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &menu, nil
}
