package main

import (
	"fmt"

	. "github.com/pspaces/gospace"
)

// Insere um novo usuário em um ambiente
func createUser(master *Space, envName string, userName string, msg string, dest string, recv string) {
	//Verifica se o ambiente existe
	_, err1 := master.QueryP(envName)

	if err1 == nil {
		var _envName string
		var _msg string
		var _dest string
		var _recv string

		//Verifica se o usuário já existe
		t, err2 := master.QueryP(&_envName, userName, &_msg, &_dest, &_recv)

		if err2 == nil {
			fmt.Printf("O usuário %s já existe no ambiente %s!\n", userName, t.GetFieldAt(0))
		} else {
			master.Put(envName, userName, msg, dest, recv)
		}
	} else {
		fmt.Printf("O ambiente %s não existe\n", envName)
	}
}

// Lista os usuários de um ambiente
func listUser(master *Space) string {
	var envName string
	var userName string
	var msg string
	var dest string
	var recv string

	fmt.Printf("Ambientes disponíveis: \n")
	listEnv(master)
	fmt.Printf("Listar usuários de qual ambiente?\n")
	fmt.Scanf("%s", &envName)

	t, _ := master.QueryAll(envName, &userName, &msg, &dest, &recv)

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i].GetFieldAt(0), t[i].GetFieldAt(1))
	}
	return envName
}

// Move um usuário de um ambiente
func moveUser(master *Space) {
	var envSource string
	var envDest string
	var user string

	var msg string
	var dest string
	var recv string

	envSource = listUser(master)

	fmt.Printf("Mover qual usuário?\n")
	fmt.Scanf("%s", &user)

	fmt.Printf("Mover usuário %s para qual ambiente?\n", user)
	fmt.Scanf("%s", &envDest)

	//Remove o usuário do ambiente origem
	master.GetP(envSource, user, &msg, &dest, &recv)

	//Insere o usuário no ambiente destino
	master.Put(envDest, user, "", "", "")
}

func main() {
	client := NewRemoteSpace("tcp://localhost:12765/master")

	client.Put("amb1", "user1")
}
