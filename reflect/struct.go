package reflect

import "fmt"

type User struct {
	Id    int
	Name  string
	uint8 // 匿名属性
}

func (s User) String() string {
	return fmt.Sprintf("id:%v, name:%v", s.Id, s.Name)
}

func (s User) GetId() int {
	return s.Id
}

func (s User) GetName() string {
	return s.Name
}

func (s *User) SetId(id int) {
	s.Id = id
}

func (s *User) SetName(name string) {
	s.Name = name
}
