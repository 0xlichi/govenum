// Package output
package output

func Success(text string) string {
	return BoldGreenColor("[+] " + text)
}

func Error(text string) string {
	return BoldRedColor("[-] " + text)
}

func Warning(text string) string {
	return BoldYellowColor("[!] " + text)
}

func Info(text string) string {
	return BlueColor("[*] " + text)
}
