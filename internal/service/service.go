package service

import (
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TextOrMorse(data []byte) string {
	isMorse := true
	count := 0
	if data == nil {
		log.Fatal("Отсутствует data из file")
	}
	dataToString := string(data)

	for _, v := range dataToString {
		if v == '.' || v == '-' || v == ' ' {
			count++
		} else {
			isMorse = false
			break
		}
	}

	if isMorse && count == len(dataToString) {
		dataText := morse.ToText(dataToString)
		fmt.Println(dataText)
		return dataText
	}

	dataMorse := morse.ToMorse(dataToString)
	fmt.Println(dataMorse)
	return dataMorse
}
