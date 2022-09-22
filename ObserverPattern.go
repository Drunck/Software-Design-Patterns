package main

import (
	_ "errors"
	"fmt"
)

type Observer interface {
	handleEvent([]string)
}

type Observable interface {
	subscribe(o Observer) error
	unsubscribe(o Observer) error
	notifyAll()
}

type Person struct {
	name string
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Println("Hi dear " + p.name)
	fmt.Println("Vacancies updated: ")
	for id := range vacancies {
		fmt.Println(vacancies[id])
	}
}

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

/*
func (j *JobSite) subscribe(ob Observer) error {
	for _, observer := range j.subscribers {
		if ob == observer {
			return fmt.Errorf("observer already exists")
		}
	}
	j.subscribers = append(j.subscribers, ob)
	return nil
}

func (j *JobSite) unsubscribe(ob Observer) error {
	for i, observer := range j.subscribers {
		if ob == observer {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("observer not found")
}
*/

func (j *JobSite) subscribe(ob Observer) error {
	for _, observer := range j.subscribers {
		if observer == ob {
			fmt.Printf("Observer %v already exists!\n", observer)
			return nil
		}
	}
	j.subscribers = append(j.subscribers, ob)
	fmt.Printf("Observer %v subscribed.\n", ob)
	return nil
}

func (j *JobSite) unsubscribe(ob Observer) error {
	for i, observer := range j.subscribers {
		if ob == observer {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
			return nil
		}
	}
	fmt.Printf("Observer %v not found!\n", ob)
	return nil
}

func (j *JobSite) notifyAll() {
	for _, observer := range j.subscribers {
		observer.handleEvent(j.vacancies)
	}
}

func (j *JobSite) addVacancies(vacancy string) error {
	j.vacancies = append(j.vacancies, vacancy)
	j.notifyAll()
	return nil
}

func (j *JobSite) removeVacancies(vacancy string) error {
	length := len(j.vacancies)
	counter := 0
	for i, v := range j.vacancies {
		if v == vacancy {
			j.vacancies = append(j.vacancies[:i], j.vacancies[i+1:]...)
			j.notifyAll()
			return nil
		}
		counter++
	}
	if counter == length {
		fmt.Printf("Vacancy '%v' not found!\n", vacancy)
	}
	return nil
}

func main() {
	jobSite := &JobSite{
		subscribers: nil,
		vacancies:   nil,
	}

	personA := &Person{name: "Nurkuisa"}
	personB := &Person{name: "Bob"}

	jobSite.subscribe(personA)
	jobSite.subscribe(personB)
	jobSite.subscribe(personB)

	jobSite.notifyAll()
	fmt.Println("===================================")
	jobSite.addVacancies("Back-End Developer")

	fmt.Println("===================================")
	jobSite.unsubscribe(personB)
	jobSite.unsubscribe(personB)

	fmt.Println("===================================")
	jobSite.addVacancies("Front-End Developer")

	fmt.Println("===================================")
	jobSite.removeVacancies("Back Developer")

	fmt.Println("===================================")
	jobSite.removeVacancies("Back-End Developer")

	fmt.Println("===================================")
	jobSite.addVacancies("Game Designer")

}
