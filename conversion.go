package ps

import (
	"encoding/json"
	"fmt"

	"github.com/fmagana-fhps/powerschool-api-go/models"
)

func convertBody[T any](resp models.Response[any], newResponse T) models.Response[T] {
	body, ok := resp.Body.(*T)
	if !ok {
		panic("unexpected type during conversion")
	}

	return models.Response[T]{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       *body,
	}
}

func CreatePowerqueryModels[T any](fields []map[string]any) (models []T) {
	model := new(T)
	for _, record := range fields {
		obj := []byte{'{'}
		for key, value := range record {
			obj = fmt.Appendf(obj, `"%s":"%v",`, key, value)
		}

		obj = fmt.Append(obj[:len(obj)-1], "}")
		err := json.Unmarshal(obj, model)
		if err != nil {
			panic(err)
		}
		models = append(models, *model)
	}

	return models
}
