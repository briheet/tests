package maps

import "errors"

func Search(mpp map[string]any, object string) (any, error) {
	for key, value := range mpp {
		if object == key {
			return value, nil
		}
	}

	return nil, errors.New("could not find the word you are looking for")
}
