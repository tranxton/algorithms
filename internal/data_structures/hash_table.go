package data_structures

var votedList map[string]bool

func CreatePhoneBook() map[string]string {
	return make(map[string]string)
}

func AddNumberToPhoneBook(phoneBook map[string]string, name string, phone string) map[string]string {
	phoneBook[name] = phone

	return phoneBook
}

func GetNumberFromPhoneBook(phoneBook map[string]string, name string) string {
	return phoneBook[name]
}

func GetNameFromPhoneBook(phoneBook map[string]string, searchablePhone string) string {
	for name, phone := range phoneBook {
		if phone == searchablePhone {
			return name
		}
	}

	return ""
}

func CanVote(name string) bool {
	if !votedList[name] {
		votedList[name] = true

		return true
	}

	return false
}
