package tasks

import (
	"database/sql"
	"fmt"
	"server/types"
	"server/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) DeleteTask(t *types.Task) error {

	_, err := s.db.Exec("delete from tasks where id=$1", t.TaskID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateTask(t *types.Task) error {

	_, err := s.db.Exec("update tasks set (isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done) = ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) where id=$12", t.IsImportant, t.TaskName, t.Difficulty, t.SDescription, t.Type, t.Stats, t.Deadline, t.Repeat, t.Subtask, t.FDescription, t.Done, t.TaskID)

	if err != nil {
		return err
	}
	return nil

}

func (s *Store) CreateTask(t *types.Task) error {

	_, err := s.db.Exec("insert into tasks (userid,isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)", t.UserID, t.IsImportant, t.TaskName, t.Difficulty, t.SDescription, t.Type, t.Stats, t.Deadline, t.Repeat, t.Subtask, t.FDescription, t.Done)
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

func (s *Store) GetSortedTasks(userid int, date utils.JsonDate) (*types.GetTasksResponse, error) {

	tasklist := &types.GetTasksResponse{}
	today, err := s.db.Query("select userid,id,isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done from tasks where userid=$1 and deadline=$2", userid, date.String())
	if err != nil {

		return nil, fmt.Errorf("error getting today tasks %v", err)
	}
	var t *types.Task
	tasklist.Today = make([]types.Task, 0)
	for today.Next() {
		t, err = scanRowsIntoTask(today)
		if err != nil {
			return nil, err
		}
		tasklist.Today = append(tasklist.Today, *t)
	}

	important, err := s.db.Query("select userid,id,isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done from tasks where userid=$1 and isImportant='true' and deadline>$2 and done='false'", userid, date.String())
	if err != nil {
		return nil, fmt.Errorf("error getting important tasks %v", err)
	}
	tasklist.Important = make([]types.Task, 0)
	for important.Next() {
		t, err = scanRowsIntoTask(important)
		if err != nil {
			return nil, err
		}
		tasklist.Important = append(tasklist.Important, *t)
	}

	unfinished, err := s.db.Query("select userid,id,isImportant,taskname,difficulty,sdescription,tasktype,taskstats,deadline,taskrepeat,subtask,fdescription,done from tasks where userid=$1 and done='false' and deadline<$2", userid, date.String())
	if err != nil {
		return nil, fmt.Errorf("error getting unfinished tasks %v", err)
	}
	tasklist.Unfinished = make([]types.Task, 0)
	for unfinished.Next() {
		t, err = scanRowsIntoTask(unfinished)
		if err != nil {
			return nil, err
		}
		tasklist.Unfinished = append(tasklist.Unfinished, *t)
	}

	return tasklist, nil

}
