package sinonimo

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Test with multiple levels synonyms", func(t *testing.T) {
		wordToSearch := "deseo"
		wantedResults := []string{"aspiración", "ansia", "afán", "anhelo", "apetito", "pretensión", "capricho", "empeño", "antojo", "pasión", "ambición", "interés", "intención", "objetivo", "proyecto", "apetecer", "querer", "aspirar", "anhelar", "ambicionar", "ansiar", "codiciar", "suspirar", "pretender", "preferir", "aficionarse", "antojarse", "encapricharse", "prendarse", "pirrarse"}

		searchResults, err := WordReference{}.search(wordToSearch)

		if err != nil {
			t.Error("error while searching the world")
		}

		result := reflect.DeepEqual(searchResults, wantedResults)

		if result == false {
			t.Errorf("Words are not the same, received: %v, wanted: %v", searchResults, wantedResults)
		}
	})

	t.Run("Test synonyms that contains space", func(t *testing.T) {
		wordToSearch := "adios"
		wantedResults := []string{"chao", "abur", "agur", "despedida", "hasta la vista", "hasta más ver"}

		searchResults, err := WordReference{}.search(wordToSearch)

		if err != nil {
			t.Error("error while searching the world")
		}

		result := reflect.DeepEqual(searchResults, wantedResults)

		if result == false {
			t.Errorf("Words are not the same, received: %v, wanted: %v", searchResults, wantedResults)
		}
	})

	t.Run("Test word that does not contain synonyms", func(t *testing.T) {
		wordToSearch := "asdasdjk"
		wantedResults := []string{}

		searchResults, err := WordReference{}.search(wordToSearch)

		if err != nil {
			t.Error("error while searching the world")
		}

		result := reflect.DeepEqual(searchResults, wantedResults)

		if result == false {
			t.Errorf("Words are not the same, received: %v, wanted: %v", searchResults, wantedResults)
		}
	})

}
