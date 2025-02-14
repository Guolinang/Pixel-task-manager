package tasks

import (
	"database/sql"
	"server/types"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTask(t *types.Task) error {

	_, err := s.db.Exec("insert into tasks (userid,id,isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", t.UserID, t.TaskID, t.IsImportant, t.TaskName, t.Difficulty, t.SDescription, t.Type, t.Stats, t.Deadline, pq.Array(t.Repeat), pq.Array(t.Subtask), t.FDescription, t.Done)
	if err != nil {
		return err
	}
	return nil

}

func (s *Store) GetUserTasks(userid int) ([]types.Task, error) {

	rows, err := s.db.Query("select userid,id,isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done from tasks where userid=$1", userid)
	if err != nil {
		return nil, err
	}
	var t *types.Task
	tarr := make([]types.Task, 0)
	for rows.Next() {
		t, err = scanRowsIntoTask(rows)
		if err != nil {
			return nil, err
		}
		tarr = append(tarr, *t)
	}
	return tarr, nil

}

func scanRowsIntoTask(rows *sql.Rows) (*types.Task, error) {

	t := new(types.Task)
	err := rows.Scan(&t.UserID, &t.TaskID, &t.IsImportant, &t.TaskName, &t.Difficulty, &t.SDescription, &t.Type, &t.Stats, &t.Deadline, &t.Repeat, &t.Subtask, &t.FDescription, &t.Done)
	if err != nil {
		return nil, err
	}
	return t, nil
}
