package initializers

import (
	"encoding/json"
	"github.com/DianaSun97/PaginationScratch/models"
	"io/ioutil"
	"log"
	"os"
)

// Обновляем структуру для соответствия новому формату JSON
type VerbForm struct {
	ID       int    `json:"id"`
	Word     string `json:"word"`
	Ma       string `json:"ma"`
	Sa       string `json:"sa"`
	Ta       string `json:"ta"`
	Me       string `json:"me"`
	Te       string `json:"te"`
	Nad      string `json:"nad"`
	Negative string `json:"negative"`
}

type WordsJSON struct {
	Verbs []VerbForm `json:"verbs"`
}

func CreateWords() {
	file, err := os.Open("./assets/words.json")
	if err != nil {
		log.Fatalln("Error opening JSON file:", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	byteValue, _ := ioutil.ReadAll(file)

	// Разбираем JSON в структуру WordsJSON
	var wordsData WordsJSON
	err = json.Unmarshal(byteValue, &wordsData)
	if err != nil {
		log.Fatalln("Error unmarshaling JSON:", err)
	}

	// Используем данные из wordsData для работы
	for _, verb := range wordsData.Verbs {
		word := models.Words{
			EstVerb: verb.Word, //  EstVerb хранит основную форму слова
			EngVerb: "",        // EngVerb не указан в вашем JSON, возможно, потребуется другая логика для английских слов
		}

		if err := DB.Save(&word).Error; err != nil {
			log.Fatalln(err)
		}
	}
}
