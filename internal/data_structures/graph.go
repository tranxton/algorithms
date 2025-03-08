package data_structures

type SocialNetwork struct {
	Users []*Person
}

type Person struct {
	Name    string
	Friends []*Person
}

type deque struct {
	items []*Person
}

func BreadthFirstSearch(person *Person, name string) bool {
	queue := deque{items: []*Person{person}}
	visited := map[*Person]bool{}

	for len(queue.items) > 0 {
		currentPerson := queue.items[0]
		queue.items = queue.items[1:]

		if visited[currentPerson] {
			continue
		}

		visited[currentPerson] = true

		if currentPerson.Name == name {
			return true
		}

		queue.items = append(queue.items, currentPerson.Friends...)
	}

	return false
}

func DepthFirstSearch(person *Person, name string, visited map[*Person]bool) bool {
	visited[person] = true

	if person.Name == name {
		return true
	}

	for _, personFriend := range person.Friends {
		if visited[personFriend] {
			continue
		}

		found := DepthFirstSearch(personFriend, name, visited)
		if found {
			return true
		}
	}

	return false
}

func TailDepthFirstSearch(person *Person, name string, visited map[*Person]bool, personFriends []*Person) bool {
	visited[person] = true

	if person.Name == name {
		return true
	}

	if len(personFriends) == 0 {
		if person.Friends == nil || len(person.Friends) == 0 {
			return false
		}

		return TailDepthFirstSearch(person.Friends[0], name, visited, person.Friends[1:])
	}

	return TailDepthFirstSearch(personFriends[0], name, visited, personFriends[1:])
}
