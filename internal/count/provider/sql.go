package provider

import (
	"database/sql"
	"errors"

	"github.com/cod1ng-space/web-10/pkg/vars"
)

func (p *Provider) SelectCount() (int, error) {
	var msg int

	// Получаем одно сообщение из таблицы hello, отсортированной в случайном порядке
	err := p.conn.QueryRow("SELECT num FROM counter").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, vars.ErrDBNotInitialized
		}
		return 0, err
	}
	return msg, nil
}

func (p *Provider) UpdateCount(msg int) error {
	_, err := p.conn.Exec("UPDATE counter SET num = num + $1", msg)
	if err != nil {
		return err
	}

	return nil
}
