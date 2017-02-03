package vrmp

import (
	"time"

	"log"

	"appengine"
	"appengine/datastore"
)

type Employee struct {
	Name        string
	Bio         string
	Email       string
	LastUpdated time.Time
}

func updateEmployee(ctx appengine.Context, name, bio, email string) {
	employee := &Employee{
		Name:        name,
		Bio:         bio,
		Email:       email,
		LastUpdated: time.Now(),
	}
	// use email as StringKey
	key := datastore.NewKey(ctx, "Employee", email, 0, nil)
	_, err := datastore.Put(ctx, key, employee)
	if err != nil {
		// handle err
		log.Println(err)
		return
	}
}

func getEmployee(ctx appengine.Context, email string) (Employee, error) {
	key := datastore.NewKey(ctx, "Employee", email, 0, nil)
	var employee Employee
	err := datastore.Get(ctx, key, &employee)
	if err != nil {
		log.Println(err)
	}
	return employee, err
}
