package pratica


import "fmt"

type person struct {
  id   int
  name string
  age  int
}

type people []*person

func (p *people) addPerson(person *person) {
  *p = append(*p, person)
}

func (p *people) editPerson(person *person, id int) {
  for _, personEdit := range *p {
    if personEdit.id == id {
      personEdit.age = person.age
      personEdit.name = person.name
    } else {
      fmt.Println("não foi possível achar essa pessoa")
    }
  }
}

func (p people) listPeople() {
  for _, person := range p {
    fmt.Println("Nome: ", person.name)
    fmt.Println("Idade: ", person.age)
  }
}

func playPraticaPonteiro() {
  p1 := people{}

  person1 := person{id: 1, name: "Pedro", age: 28}

  p1.addPerson(&person1)

  p1.listPeople()

  person2 := person{id: 3, name: "Pedro Silva", age: 16}

  p1.editPerson(&person2, 2)

  p1.listPeople()
}