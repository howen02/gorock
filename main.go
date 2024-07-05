package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Account struct {
	account   string
	authToken string
	domain    string
}

func main() {
	printLogo()
	accounts := extractAccounts()
	chosenAccount := getChoice(accounts)
	port := getPort()
	setupAuth(chosenAccount)
	updateEnv(chosenAccount)
	runNgrok(chosenAccount, port)
}

func printLogo() {
	gorock := `%s
                                    __    
   ____   ___________  ____   ____ |  | __
  / ___\ /  _ \_  __ \/  _ \_/ ___\|  |/ /
 / /_/  >  <_> )  | \(  <_> )  \___|    < 
 \___  / \____/|__|   \____/ \___  >__|_ \
/_____/                          \/     \/
%s`

	blue := "\033[34m"
	fmt.Println(fmt.Sprintf(gorock, blue, "\033[0m"))
	fmt.Print("Welcome to GoRock! ngrok account management made easier 🗿\n\n")
}

func extractAccounts() []Account {
	// Open accounts file
	file, err := os.Open("accounts.txt")
	accounts := make([]Account, 0)

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Read accounts from file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		if len(fields) != 3 {
			log.Fatalf("Invalid line format: " + line + "\nUse format: account, authToken, domain")
			continue
		}

		account := Account{
			account:   strings.TrimSpace(fields[0]),
			authToken: strings.TrimSpace(fields[1]),
			domain:    strings.TrimSpace(fields[2]),
		}

		accounts = append(accounts, account)
	}

	return accounts
}

func getChoice(accounts []Account) Account {
	fmt.Println("Choose an account:")
	fmt.Println("0. Exit")

	for i, account := range accounts {
		fmt.Printf("%d. %s: %s\n", i+1, account.account, account.domain)
	}

	var choice int

	// Get user choice
	for {
		fmt.Print("\nEnter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Please enter a number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if choice == 0 {
			os.Exit(0)
		}

		if choice < 0 || choice > len(accounts) {
			fmt.Println("Invalid index")
			continue
		} else {
			break
		}
	}

	return accounts[choice-1]
}

func getPort() string {
	fmt.Print("Enter port number (default 3000): ")
	var portNumber string
	fmt.Scanln(&portNumber)

	if portNumber == "" {
		portNumber = "3000"
	}

	return portNumber
}

func setupAuth(a Account) {
	// Set up auth token
	fmt.Println("Setting up auth token...")

	addTokenCommand := "ngrok"
	addTokenArgs := []string{"config", "add-authtoken", a.authToken}

	cmd := exec.Command(addTokenCommand, addTokenArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to set up auth token: %v", err)
	}
}

func updateEnv(a Account) {
	// Update .env file
	envFile, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Failed to open .env file: %v", err)
		panic(err)
	}
	defer envFile.Close()

	tempFile, err := os.Create("temp.env")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
		panic(err)
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(envFile)

	// Replace domains in .env
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`[a-z0-9-]+\.ngrok-free.app`)

		if r.MatchString(line) {
			// Replace domain while keeping the rest of the line
			start, end := r.FindStringIndex(line)[0], r.FindStringIndex(line)[1]
			line = line[:start] + a.domain + line[end:]
		}

		_, err := tempFile.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Failed to write to stdout: %v", err)
		}
	}

	envFile.Close()
	tempFile.Close()
	if err := os.Rename("temp.env", ".env"); err != nil {
		log.Fatalf("Failed to rename temp file: %v", err)
	}
}

func runNgrok(a Account, p string) {
	// Run ngrok
	runDomainCommand := "ngrok"
	runDomainArgs := []string{"http", "--domain=" + a.domain, p}

	fmt.Println("Running ngrok...")

	cmd := exec.Command(runDomainCommand, runDomainArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run ngrok: %v", err)
	}
}
