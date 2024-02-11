package initializers

import (
	"encoding/json"
	"github.com/DianaSun97/PaginationScratch/models"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

// Определяем структуру для разбора JSON
type WordsJSON struct {
	WordsEst []string `json:"words"`
}

func CreateWords() {
	// Открываем файл JSON
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
		return
	}

	// Теперь wordsData.WordsEng содержит массив английских слов
	wordsEng := wordsData.WordsEst
	wordsEst := []string{
		// Добавьте ваши эстонские слова здесь
	}

	for i := 0; i < 1000; i++ {
		estonianVerbIndex := rand.Intn(len(wordsEst))
		englishVerbIndex := rand.Intn(len(wordsEng))
		wordEst := wordsEst[estonianVerbIndex]
		wordEng := wordsEng[englishVerbIndex]

		word := models.Words{
			EstVerb: wordEst,
			EngVerb: wordEng,
		}

		if err := DB.Save(&word).Error; err != nil {
			log.Fatalln(err)
		}
	}
}
