package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Checkchars(s string) bool {
	for _, c := range s {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

func MapBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)
	charIndex := 32
	for i := 0; i < len(lines); i += 9 {
		if i+8 < len(lines) {
			banner[rune(charIndex)] = lines[i+1 : i+9]
			charIndex++
		}
	}
	return banner, nil
}

func Checknewline(inpultsplit []string) bool {
	for _, line := range inpultsplit {
		if len(line) != 0 {
			return false
		}
	}
	return true
}

func Draw(banner map[rune][]string, inpultsplit []string) {
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			fmt.Println()
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			fmt.Println()
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					fmt.Print(banner[rune(v[j])][i])
				}
				fmt.Println()
			}
		}
	}
}

func DrawInFile(banner map[rune][]string, inputsplit []string, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for idx, v := range inputsplit {
		if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
			_, err := file.WriteString("\n")
			if err != nil {
				return err
			}
			continue
		} else if len(v) == 0 && !Checknewline(inputsplit) {
			_, err := file.WriteString("\n")
			if err != nil {
				return err
			}
		} else if len(v) != 0 && !Checknewline(inputsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					_, err := file.WriteString(banner[rune(v[j])][i])
					if err != nil {
						return err
					}
				}
				_, err := file.WriteString("\n")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}