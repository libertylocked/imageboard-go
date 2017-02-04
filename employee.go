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

func getAllEmployees(ctx appengine.Context) []Employee {
	employees := []Employee{}
	q := datastore.NewQuery("Employee")
	t := q.Run(ctx)
	for {
		var e Employee
		_, err := t.Next(&e)
		if err == datastore.Done {
			break // done
		}
		if err != nil {
			log.Printf("error fetching next person %v", err)
			break
		}
		// push to arr
		employees = append(employees, e)
	}
	return employees
}
