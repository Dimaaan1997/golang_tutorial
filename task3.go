package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type Graph struct {
	users       map[int]*User
	connections map[int]map[int]bool
}

func NewGraph() *Graph {
	return &Graph{
		users:       make(map[int]*User),
		connections: make(map[int]map[int]bool),
	}
}

func (g *Graph) NewGraph() *Graph {
	graph := &Graph{users: make(map[int]*User),
		connections: make(map[int]map[int]bool)}
	return graph
}

func (g *Graph) AddUser(id int, name string) {
	g.users[id] = &User{ID: id, Name: name}
	if _, ok := g.connections[id]; !ok {
		g.connections[id] = make(map[int]bool)
	}
}

func (g *Graph) GetUser(id int) (*User, bool) {
	user, ok := g.users[id]
	return user, ok
}

func (g *Graph) AddConnection(fromID, toID int) bool {
	if _, ok := g.users[fromID]; !ok {
		return false
	}
	if _, ok := g.users[toID]; !ok {
		return false
	}
	g.connections[fromID][toID] = true
	g.connections[toID][fromID] = true
	return true

}

func (g *Graph) GetConnections(userID int) []*User {

	userConnections := []*User{}

	if _, ok := g.users[userID]; !ok {
		return userConnections
	}

	for userID := range g.connections[userID] {
		if user, ok := g.users[userID]; ok {
			userConnections = append(userConnections, user)
		}
	}
	return userConnections
}

func (g *Graph) HasConnection(fromID, toID int) bool {
	if _, ok := g.connections[fromID]; !ok {
		return false
	}

	return g.connections[fromID][toID]
}

func (g *Graph) UserCount() int {
	return len(g.users)
}

func (g *Graph) RemoveConnection(fromID, toID int) bool {
	if _, ok := g.users[fromID]; !ok {
		return false
	}
	if _, ok := g.users[toID]; !ok {
		return false
	}
	if !g.connections[fromID][toID] {
		return false
	}
	delete(g.connections[fromID], toID)
	delete(g.connections[toID], fromID)
	return true
}

func (g *Graph) RemoveUser(id int) bool {
	if _, ok := g.users[id]; !ok {
		return false
	}
	delete(g.users, id)
	delete(g.connections, id)
	for _, connectedUsers := range g.connections {
		delete(connectedUsers, id)
	}
	return true
}

func (g *Graph) IsMutual(id1, id2 int) bool {
	isMute := g.HasConnection(id1, id2) &&
		g.HasConnection(id2, id1)
	return isMute
}

func (g *Graph) ConnectionCount(userID int) int {
	if _, ok := g.connections[userID]; !ok {
		return 0
	}
	return len(g.connections[userID])
}

func (g *Graph) CommonConnections(id1, id2 int) []*User {
	var result []*User
	if _, ok := g.users[id1]; !ok {
		return result
	}
	if _, ok := g.users[id2]; !ok {
		return result
	}
	for friendID := range g.connections[id1] {
		if g.connections[id2][friendID] {
			result = append(result, g.users[friendID])
		}
	}
	return result
}

func (g *Graph) SuggestConnections(userID int) []*User {
	var result []*User
	if _, ok := g.users[userID]; !ok {
		return result
	}
	suggestions := make(map[int]bool)
	for friendID := range g.connections[userID] {
		for candidateID := range g.connections[friendID] {
			if candidateID == userID {
				continue
			}
			if g.HasConnection(userID, candidateID) {
				continue
			}
			suggestions[candidateID] = true
		}
	}
	for id := range suggestions {
		result = append(result, g.users[id])
	}
	return result
}

func (g *Graph) GetAllUsers() []*User {
	result := make([]*User, 0, len(g.users))
	for _, user := range g.users {
		result = append(result, user)
	}
	return result
}

func main_task3() {
	graph := NewGraph()

	graph.AddUser(1, "Alice")
	graph.AddUser(2, "Bob")
	graph.AddUser(3, "Charlie")

	graph.AddConnection(1, 2) // Alice -> Bob
	graph.AddConnection(1, 3) // Alice -> Charlie
	graph.AddConnection(2, 3) // Bob -> Charlie

	if user, ok := graph.GetUser(1); ok {
		fmt.Printf("User: %s\n", user.Name)

		friends := graph.GetConnections(1)
		fmt.Printf("Friends: %d\n", len(friends))

		for _, friend := range friends {
			fmt.Printf("  - %s\n", friend.Name)
		}
	}

	fmt.Printf("Alice and Bob connected: %v\n",
		graph.HasConnection(1, 2))

	fmt.Printf("Alice and Bob remove connection: %v\n",
		graph.RemoveConnection(1, 2))

	fmt.Printf("Alice and Bob connected: %v\n",
		graph.HasConnection(1, 2))

	// fmt.Printf("Remove User Alice: %v\n",
	// 	graph.RemoveUser(1))

	fmt.Printf("Is mutual: %v\n",
		graph.IsMutual(2, 3))

	fmt.Printf("Charlie connections count: %v\n",
		graph.ConnectionCount(3))

	fmt.Printf("Common connections Alice and Bob : %v\n",
		graph.CommonConnections(1, 2))

	fmt.Printf("Suggested Connections for Alice  : %v\n",
		graph.SuggestConnections(1))

	fmt.Printf("All users : %v\n",
		graph.GetAllUsers())

}
