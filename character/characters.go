package character

import (
	"database/sql"
	"fmt"
	"server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) UpdateCharacter(c *types.Character) error {
	_, err := s.db.Exec("update character set (levelC,expC,maxexpC,hpC,maxhpC,strC,intC,charC,wisC,cnstC,head,face,body,dress,other) = ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) where userid=$16", c.Level, c.Exp, c.MaxExp, c.Hp, c.MaxHp, c.Str, c.Int, c.Char, c.Wis, c.Cnst, c.Head, c.Face, c.Body, c.Dress, c.Other, c.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateCharacter(c *types.Character) error {

	_, err := s.db.Exec("insert into character (userid,levelC,expC,maxexpC,hpC,maxhpC,strC,intC,charC,wisC,cnstC, head, face, body, dress,other) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)", c.UserID, c.Level, c.Exp, c.MaxExp, c.Hp, c.MaxHp, c.Str, c.Int, c.Char, c.Wis, c.Cnst, c.Head, c.Face, c.Body, c.Dress, c.Other)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCharacter(userid int) (*types.Character, error) {

	rows, err := s.db.Query("select userid,levelC,expC,maxexpC,hpC,maxhpC,strC,intC,charC,wisC,cnstC,head,face,body,dress,other from character where userid=$1", userid)
	if err != nil {
		return nil, err
	}
	var c *types.Character
	for rows.Next() {
		c, err = scanRowsIntoCharacter(rows)
		if err != nil {
			return nil, err
		}
	}
	if c == nil {
		return nil, fmt.Errorf("character not found")
	}
	return c, nil
}

func scanRowsIntoCharacter(rows *sql.Rows) (*types.Character, error) {
	c := new(types.Character)
	err := rows.Scan(&c.UserID, &c.Level, &c.Exp, &c.MaxExp, &c.Hp, &c.MaxHp, &c.Str, &c.Int, &c.Char, &c.Wis, &c.Cnst, &c.Head, &c.Face, &c.Body, &c.Dress, &c.Other)
	if err != nil {
		return nil, err
	}
	return c, nil

}
