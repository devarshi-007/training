package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var userMap = struct {
	sync.RWMutex
	m map[int]User
}{m: make(map[int]User)}

func init() {
	userMapLoc, err := loadUser()
	if err != nil {
		log.Fatal(err)
	}
	userMap.m = userMapLoc
}

func loadUser() (map[int]User, error) {
	filename := "data/userD.json"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, errors.New("FIle not exists")
	}
	file, _ := ioutil.ReadFile(filename)
	userList := make([]User, 0)
	err = json.Unmarshal([]byte(file), &userList)
	if err != nil {
		log.Fatal(err)
	}
	userMapLoc := make(map[int]User)
	for i := 0; i < len(userList); i++ {
		userMapLoc[userList[i].Id] = userList[i]
	}
	return userMapLoc, nil
}

func addOrUpdateUser(u User) (int, error) {
	aId := -1
	if u.Id > 0 {
		oldUser := getUser(u.Id)
		if oldUser == nil {
			return 0, fmt.Errorf("User does not exist")
		}
		aId = oldUser.Id
	} else {
		aId = getNextId()
		u.Id = aId
	}
	userMap.Lock()
	userMap.m[aId] = u
	userMap.Unlock()
	return aId, nil
}

func getNextId() int {
	userLs := getUserIds()
	return userLs[len(userLs)-1] + 1
}

func getUserList() []User {
	userMap.RLock()
	users := make([]User, 0, len(userMap.m))
	fmt.Println(users)
	for _, value := range userMap.m {
		users = append(users, value)
	}
	userMap.RUnlock()
	return users
}

func getUser(id int) *User {
	userMap.RLock()
	defer userMap.RUnlock()
	if u, ok := userMap.m[id]; ok {
		return &u
	}
	return nil
}

func getUserIds() []int {
	userMap.RLock()
	userIds := []int{}
	for key := range userMap.m {
		userIds = append(userIds, key)
	}
	userMap.RUnlock()
	sort.Ints(userIds)
	return userIds
}

func removeUser(userId int) {
	userMap.Lock()
	defer userMap.Unlock()
	delete(userMap.m, userId)
}
