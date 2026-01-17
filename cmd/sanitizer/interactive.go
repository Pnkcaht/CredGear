package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cred-sanitizer/internal/integration"
)

/*
ANSI colors / styles
*/
const (
	reset = "\033[0m"
	bold  = "\033[1m"

	green = "\033[32m"
	red   = "\033[31m"
	cyan  = "\033[36m"
	gray  = "\033[90m"
)

func runInteractive() {
	clearScreen()
	printBanner()
	menuLoop()
}

func menuLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()
		fmt.Print(bold + "› " + reset)

		opt, _ := reader.ReadString('\n')
		opt = strings.TrimSpace(opt)

		switch opt {

		case "1":
			path := ask(reader, "File path")
			fmt.Println(green + "\n[+] Cred Gear running in STRICT mode...\n" + reset)
			runFromFile(path, integration.Config{
				Strict:  true,
				Format:  "auto",
				Metrics: true,
			})

		case "2":
			path := ask(reader, "File path")
			fmt.Println(green + "\n[+] Cred Gear running in LOOSE mode...\n" + reset)
			runFromFile(path, integration.Config{
				Loose:   true,
				Format:  "auto",
				Metrics: true,
			})

		case "3":
			fmt.Println(cyan + "\n[+] Paste credentials below (Ctrl+D to finish):\n" + reset)
			runFromStdin(integration.Config{
				Format:  "auto",
				Metrics: true,
			})

		case "4":
			path := ask(reader, "File path")
			format := ask(reader, "Format (txt / csv / sql)")
			fmt.Println(green + "\n[+] Cred Gear processing with selected format...\n" + reset)
			runFromFile(path, integration.Config{
				Format:  format,
				Metrics: true,
			})

		case "5":
			showHelp()

		case "0":
			fmt.Println("\n" + gray + "Bye 👋" + reset)
			return

		default:
			fmt.Println(red + "\n[!] Invalid option\n" + reset)
		}
	}
}

func printMenu() {
	fmt.Println(`
────────────────────────────────
  [1] Process file (STRICT)
  [2] Process file (LOOSE)
  [3] Process from STDIN
  [4] Process with format
  [5] Help
  [0] Exit
────────────────────────────────`)
}

func ask(r *bufio.Reader, label string) string {
	fmt.Printf("%s%s:%s ", bold, label, reset)
	v, _ := r.ReadString('\n')
	return strings.TrimSpace(v)
}

func showHelp() {
	fmt.Println(`
Cred Gear CLI

Purpose:
  Heavy-duty credential processor.
  Extracts, normalizes and deduplicates credentials
  from messy dumps, logs and leaks.

Modes:
  Interactive → ./cred-gear
  CLI flags  → cred-gear [flags] <file>

Flags:
  --strict     Strong validation
  --loose      Permissive parsing
  --metrics    Show statistics
  --format     auto | txt | csv | sql
  --out        Save output to file

Examples:
  cred-gear --strict dump.txt
  cred-gear --format csv creds.csv
`)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
