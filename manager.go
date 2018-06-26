package main

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/pspaces/gospace"
)

//Insere um novo ambiente na Tuple Space
func createEnv(master *Space, envName string) {

	//Verifica se o ambiente já existe e se não existir cria o novo ambiente
	_, err := master.QueryP(envName)

	if err == nil {
		fmt.Printf("O ambiente %s já existe!\n", envName)
	} else {
		master.Put(envName)
		fmt.Printf("Ambiente %s criado com sucesso!", envName)
	}
}

//Lista todos os ambientes
func listEnv(master *Space) {

	var envName string
	var deviceName string
	var userName string
	var msg string
	var dest string
	var recv string

	t, _ := master.QueryAll(&envName)

	fmt.Printf("Ambientes disponíveis: \n")
	for i := 0; i < len(t); i++ {
		device, _ := master.QueryAll(t[i].GetFieldAt(0), &deviceName)
		user, _ := master.QueryAll(t[i].GetFieldAt(0), &userName, &msg, &dest, &recv)

		fmt.Printf("Ambiente: %s | qtdDevices: %d | qtdUsers: %d\n", t[i].GetFieldAt(0), len(device), len(user))
	}
	fmt.Printf("------------------------------------------\n")
}

//Insere um novo dispositivo em um ambiente
func createDevice(master *Space, envName string, deviceName string) {

	//Verifica se o ambiente existe
	_, err1 := master.QueryP(envName)

	if err1 == nil {
		var _envName string

		//Verifica se o dispositivo já existe em algum ambiente
		t, err2 := master.QueryP(&_envName, deviceName)

		if err2 == nil {
			fmt.Printf("O dispositivo %s já existe no ambiente %s, tente novamente!\n", deviceName, t.GetFieldAt(0))
		} else {
			master.Put(envName, deviceName)
			fmt.Printf("Dispositivo %s criado com sucesso no ambiente %s!", deviceName, envName)
		}
	} else {
		fmt.Printf("O ambiente %s não existe, tente novamente!\n", envName)
	}
}

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
			fmt.Printf("O usuário %s já existe no ambiente %s, tente novamente!\n", userName, t.GetFieldAt(0))
		} else {
			master.Put(envName, userName, msg, dest, recv)
			fmt.Printf("Usuário %s criado com sucesso no ambiente %s!", userName, envName)
		}
	} else {
		fmt.Printf("O ambiente %s não existe\n", envName)
	}
}

//Lista os dispositivos de um ambiente
func listDevice(master *Space, envName string) {
	var deviceName string

	t, _ := master.QueryAll(envName, &deviceName)

	fmt.Printf("\nDispositivos do ambiente %s:\n", envName)
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i].GetFieldAt(1))
	}
}

// Lista os usuários de um ambiente
func listUser(master *Space, envName string) {
	var userName string
	var msg string
	var dest string
	var recv string

	t, _ := master.QueryAll(envName, &userName, &msg, &dest, &recv)

	fmt.Printf("\nUsuários do ambiente %s:\n", envName)
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i].GetFieldAt(1))
	}
}

// Move um dispositivo de um ambiente
func moveDevice(master *Space, envSource string) {
	var envDest string
	var device string

	fmt.Printf("\nMover qual dispositivo? ")
	fmt.Scanf("%s", &device)

	fmt.Printf("Mover dispositivo %s para qual ambiente? ", device)
	fmt.Scanf("%s", &envDest)

	//Remove o dispositivo do ambiente origem
	master.GetP(envSource, device)

	//Insere o dispositivo no ambiente destino
	master.Put(envDest, device)

	fmt.Printf("Dispositivo %s movido com sucesso do ambiente %s para o ambiente %s", device, envSource, envDest)
}

// Move um usuário de um ambiente
func moveUser(master *Space, envSource string) {
	var envDest string
	var user string

	var msg string
	var dest string
	var recv string

	fmt.Printf("\nMover qual usuário? ")
	fmt.Scanf("%s", &user)

	fmt.Printf("Mover usuário %s para qual ambiente? ", user)
	fmt.Scanf("%s", &envDest)

	//Remove o usuário do ambiente origem
	master.GetP(envSource, user, &msg, &dest, &recv)

	//Insere o usuário no ambiente destino
	master.Put(envDest, user, "", "", "")

	fmt.Printf("Usuário %s movido com sucesso do ambiente %s para o ambiente %s", user, envSource, envDest)
}

//Destrói um ambiente
func destroyEnv(master *Space) {

	var envName string
	t, _ := master.QueryAll(&envName)

	//Variáveis utilizadas para construir os templates
	var deviceName string
	var userName string
	var msg string
	var dest string
	var recv string

	var emptyEnv = make([]string, len(t))

	for i := 0; i < len(t); i++ {

		//Verifica se existem dispositivos ou usuários no ambiente
		_, err1 := master.QueryP(t[i].GetFieldAt(0), &deviceName)
		_, err2 := master.QueryP(t[i].GetFieldAt(0), &userName, &msg, &dest, &recv)

		if err1 != nil && err2 != nil {
			emptyEnv[i] = t[i].GetFieldAt(0).(string)
		}
	}

	var willBeDestroyed string
	//Percorre o array
	for _, v := range emptyEnv {
		if len(v) > 0 {
			fmt.Printf("O ambiente %s está vazio e pode ser destruído\n", v)
		}
	}

	if len(emptyEnv) > 0 {
		fmt.Printf("\nQual ambiente você deseja destruir? ")
		fmt.Scanf("%s", &willBeDestroyed)

		if contains(emptyEnv, willBeDestroyed) {
			master.GetP(willBeDestroyed)
		} else {
			fmt.Printf("O ambiente %s não está vazio, portanto não pode ser destruído!\n", willBeDestroyed)
		}

	} else {
		fmt.Printf("\nNão existe nenhum ambiente vazio que possa ser destruído ")
	}
}

//Não existe a função contains nos built-in do Go
func contains(slice []string, search string) bool {
	for _, value := range slice {
		if value == search {
			return true
		}
	}
	return false
}

func main() {

	//Cria uma Tuple Space
	master := NewSpace("master")

	var option string

	var envName string
	var dvcName string
	var usrName string

	for {
		fmt.Printf("1 - Criar ambiente\n")
		fmt.Printf("2 - Criar dispositivo\n")
		fmt.Printf("3 - Criar usuário\n\n")

		fmt.Printf("4 - Destruir ambiente\n\n")

		fmt.Printf("5 - Listar dispositivos\n")
		fmt.Printf("6 - Listar usuários\n\n")
		fmt.Printf("7 - Mover dispositivo\n")
		fmt.Printf("8 - Mover usuário\n")

		fmt.Printf("Opção: ")
		fmt.Scanf("%s", &option)

		if option == "1" {
			clearScreen()
			fmt.Printf("Nome do ambiente: ")
			fmt.Scanf("%s", &envName)

			createEnv(&master, envName)

		} else if option == "2" {
			clearScreen()
			listEnv(&master)
			fmt.Printf("Em qual ambiente criar o dispositivo? ")
			fmt.Scanf("%s", &envName)
			fmt.Printf("Nome do dispositivo: ")
			fmt.Scanf("%s", &dvcName)

			createDevice(&master, envName, dvcName)

		} else if option == "3" {
			clearScreen()
			listEnv(&master)
			fmt.Printf("Em qual ambiente criar o usuário? ")
			fmt.Scanf("%s", &envName)
			fmt.Printf("Nome do usuário: ")
			fmt.Scanf("%s", &usrName)

			createUser(&master, envName, usrName, "", "", "")

		} else if option == "4" {
			clearScreen()
			listEnv(&master)
			destroyEnv(&master)

		} else if option == "5" {
			clearScreen()
			listEnv(&master)
			fmt.Printf("Listar dispositivos de qual ambiente? ")
			fmt.Scanf("%s", &envName)
			listDevice(&master, envName)

		} else if option == "6" {
			clearScreen()
			listEnv(&master)
			fmt.Printf("Listar usuários de qual ambiente? ")
			fmt.Scanf("%s", &envName)
			listUser(&master, envName)

		} else if option == "7" {
			clearScreen()
			listEnv(&master)
			fmt.Printf("Mover dispositivo de qual ambiente? ")
			fmt.Scanf("%s", &envName)
			listDevice(&master, envName)
			moveDevice(&master, envName)

		} else if option == "8" {
			clearScreen()
			listEnv(&master)
			fmt.Printf("Mover usuário de qual ambiente? ")
			fmt.Scanf("%s", &envName)
			listUser(&master, envName)
			moveUser(&master, envName)

		} else {
			fmt.Printf("---------------------------------\n")
			fmt.Printf("Opção %s inválida, tente novamente!\n", option)
			fmt.Printf("---------------------------------\n")
		}

		fmt.Printf("\n\n")
		// clearScreen()
	}

}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
