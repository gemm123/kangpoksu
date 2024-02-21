package shipping

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"kopoksu/internal/model"
	"log"
	"net/http"
	"strings"
)

type shippingHandler struct {
}

func NewShippingHandler() *shippingHandler {
	return &shippingHandler{}
}

func (h *shippingHandler) GetCity(ctx *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.rajaongkir.com/starter/city", nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("key", "20b7e685ae9b8e404a9a8203e7da5f0a")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var resp model.RespRajaOngkir

	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println(err)
		return
	}

	type dataCity struct {
		CityId   string `json:"city_id"`
		CityType string `json:"city_type"`
		CityName string `json:"city_name"`
	}

	var data []dataCity

	for _, r := range resp.RajaOngkir.Result {
		newData := dataCity{
			CityId:   r.CityID,
			CityType: r.Type,
			CityName: r.CityName,
		}

		data = append(data, newData)
	}

	ctx.JSON(response.StatusCode, gin.H{
		"data": data,
	})
}

func (h *shippingHandler) GetType(ctx *gin.Context) {
	destination := ctx.Query("destination")
	weight := ctx.Query("weight")

	client := &http.Client{}

	payload := strings.NewReader(fmt.Sprintf("origin=78&destination=%s&weight=%s&courier=jne", destination, weight))

	req, err := http.NewRequest("POST", "https://api.rajaongkir.com/starter/cost", payload)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("key", "20b7e685ae9b8e404a9a8203e7da5f0a")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var resp model.RespRajaOngkirCost

	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(response.StatusCode, resp.RajaOngkir.Result[0].Costs)
}
