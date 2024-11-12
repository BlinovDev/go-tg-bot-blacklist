package tgblacklist

import (
	"encoding/json"
	"fmt"
	"os"
)

var strategy string
var list []string

func SetStrategy(s string) (string, error) {
	if s == "BL" || s == "WL" {
		strategy = s
		return strategy, nil
	} else {
		return "", fmt.Errorf("invalid strategy: %s", s)
	}
}

func IsBlocked(username string) (bool, error) {
	if strategy == "BL" {
		inBL, err := isInBlackList(username)
		if err != nil {
			return false, err
		}
		return inBL, nil
	} else if strategy == "WL" {
		inWL, err := isInWhiteList(username)
		if err != nil {
			return false, err
		}
		return !inWL, nil
	} else {
		return false, fmt.Errorf("strategy not set")
	}
}

func GetList() ([]string, error) {
	err := loadList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func AddToList(username string) error {
	err := loadList()
	if err != nil {
		return err
	}
	for _, u := range list {
		if u == username {
			// Username already in list
			return nil
		}
	}
	list = append(list, username)
	return saveList()
}

func DeleteFromList(username string) error {
	err := loadList()
	if err != nil {
		return err
	}
	for i, u := range list {
		if u == username {
			list = append(list[:i], list[i+1:]...)
			return saveList()
		}
	}
	// Username not found
	return nil
}

func isInBlackList(username string) (bool, error) {
	return isInList(username)
}

func isInWhiteList(username string) (bool, error) {
	return isInList(username)
}

func isInList(username string) (bool, error) {
	err := loadList()
	if err != nil {
		return false, err
	}
	for _, u := range list {
		if u == username {
			return true, nil
		}
	}
	return false, nil
}

func loadList() error {
	data, err := os.ReadFile("list.json")
	if err != nil {
		if os.IsNotExist(err) {
			list = []string{}
			return nil
		} else {
			return err
		}
	}
	return json.Unmarshal(data, &list)
}

func saveList() error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile("list.json", data, 0644)
}
