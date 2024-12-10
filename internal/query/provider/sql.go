package provider

import (
	"database/sql"
	"errors"

	"github.com/cod1ng-space/web-10/pkg/vars"
)

func (p *Provider) SelectQuery() (string, error) {
	var msg string

	// Получаем одно сообщение из таблицы hello, отсортированной в случайном порядке
	err := p.conn.QueryRow("SELECT name FROM query").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", vars.ErrDBNotInitialized
		}
		return "", err
	}
	return msg, nil
}

func (p *Provider) UpdateQuery(msg string) error {
	_, err := p.conn.Exec("UPDATE query SET name = $1", msg)
	if err != nil {
		return err
	}

	return nil
}
