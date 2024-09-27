package services

import (
	"trocup-user/repository"
)

func DeleteUser(id string) error {
	return repository.DeleteUser(id)  // Appel au repository avec l'ID en tant que chaîne
}
