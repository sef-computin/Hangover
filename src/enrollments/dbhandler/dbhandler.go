package dbhandler

import (
	"database/sql"
	"fmt"
	_ "log"

	"github.com/google/uuid"
	"github.com/sef-comp/Hangover/enrollments/models"
)

type EnrollmentDB interface {
	GetAllEnrollments() ([]*models.Enrollment, error)
	GetEnrollmentsByEvent(event_id uuid.UUID) ([]*models.Enrollment, error)
	GetEnrollmentsByUser(user_id uuid.UUID) ([]*models.Enrollment, error)

	Enroll(user_id, event_id uuid.UUID) (uuid.UUID, error)
	CancelEnroll(enrollment_id uuid.UUID) error
}

type DBHandler struct {
	db *sql.DB
}

func InitDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{
		db: db,
	}
}

func (dbh *DBHandler) GetAllEnrollments() ([]*models.Enrollment, error) {
	var enrollments []*models.Enrollment
	rows, err := dbh.db.Query(`SELECT * FROM business.enrolments;`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	for rows.Next() {
		enroll := new(models.Enrollment)
		if err := rows.Scan(&enroll.EnrollmentID, &enroll.UserID, &enroll.EventID, &enroll.Status); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		enrollments = append(enrollments, enroll)
	}

	defer rows.Close()

	return enrollments, nil
}

func (dbh *DBHandler) GetEnrollmentsByEvent(event_id uuid.UUID) ([]*models.Enrollment, error){
	var enrollments []*models.Enrollment
	rows, err := dbh.db.Query(`SELECT * FROM business.enrolments WHERE event_id = $1;`, event_id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	for rows.Next() {
		enroll := new(models.Enrollment)
		if err := rows.Scan(&enroll.EnrollmentID, &enroll.UserID, &enroll.EventID, &enroll.Status); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		enrollments = append(enrollments, enroll)
	}

	defer rows.Close()

	return enrollments, nil
}

func (dbh *DBHandler) GetEnrollmentsByUser(user_id uuid.UUID) ([]*models.Enrollment, error){
	var enrollments []*models.Enrollment
	rows, err := dbh.db.Query(`SELECT * FROM business.enrolments WHERE user_id = $1;`, user_id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	for rows.Next() {
		enroll := new(models.Enrollment)
		if err := rows.Scan(&enroll.EnrollmentID, &enroll.UserID, &enroll.EventID, &enroll.Status); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		enrollments = append(enrollments, enroll)
	}

	defer rows.Close()

	return enrollments, nil
}

func (dbh *DBHandler) Enroll(user_id, event_id uuid.UUID) (uuid.UUID, error){
	enrollment_id, err := uuid.NewRandom()
	if err != nil{
		return uuid.UUID{}, err
	}
	_, err = dbh.db.Query(
		`INSERT INTO business.enrolments (enroll_id, event_id, user_id, status) VALUES ($1, $2, $3, $4);`,
		enrollment_id,
		event_id,
		user_id,
		models.ENROLLMENT_STATUS_INITIAL,
	)

	// log.Println(err)

	return enrollment_id, err
}

func (dbh *DBHandler) CancelEnroll(enrollment_id uuid.UUID) error{

	_, err := dbh.db.Query(
		`UPDATE business.enrolments SET status = $1 WHERE enroll_id = $2;`,
		models.ENROLLMENT_STATUS_CANCELLED,
		enrollment_id,
	)

	return err
}
