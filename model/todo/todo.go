package todo

// Todo struct
type Todo struct {
	ID   int
	Name string
}

// All func
func All() []Todo {
	return []Todo{
		{ID: 1, Name: "勉強"},
		{ID: 2, Name: "家事"},
		{ID: 3, Name: "愛ちゃん"},
	}
}

// Get func
func Get(id int) Todo {
	for _, v := range All() {
		if v.ID == id {
			return v
		}
	}
	return Todo{}
}

// Create func
func Create(name string) []Todo {
	all := All()
	return append(all, Todo{
		ID:   nextID(all),
		Name: name,
	})
}

func nextID(all []Todo) int {
	var id int
	for _, v := range all {
		if id < v.ID {
			id = v.ID
		}
	}
	return id + 1
}
