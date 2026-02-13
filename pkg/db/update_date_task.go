package db

func UpdateDateTask(nextDate string, id string) error {
	_, err := DB.Exec(
		"UPDATE scheduler SET date = ? WHERE id = ?",
		nextDate,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
