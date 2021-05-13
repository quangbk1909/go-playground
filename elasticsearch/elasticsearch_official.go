package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9201",
		},
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}

	/*-----test get document-------*/

	//resp, err := es.Get("notification_logs", "6d68554f-0e59-4051-9872-c62fb68cbf95", es.Get.WithContext(context.Background()))
	//defer func() {
	//	_ = resp.Body.Close()
	//}()
	//bodyResp, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Status())
	//fmt.Println(string(bodyResp))
	//log.Println(strings.Repeat("-", 37))

	/*-----test search document by "doc"-------*/
	//var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"bool": map[string]interface{}{
	//			"must": []map[string]interface{}{
	//				{
	//					"term": map[string]interface{}{
	//						"user_id": map[string]interface{}{
	//							"value": "quang123",
	//						},
	//					},
	//				},
	//				{
	//					"term": map[string]interface{}{
	//						"request_id": map[string]interface{}{
	//							"value": "10",
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//if err := json.NewEncoder(&buf).Encode(query); err != nil {
	//	log.Fatalf("Error encoding query: %s", err)
	//}
	//
	//// Perform the search request.
	//resp, err := es.Search(
	//	es.Search.WithContext(context.Background()),
	//	es.Search.WithIndex("notification_logs"),
	//	es.Search.WithBody(&buf),
	//	es.Search.WithTrackTotalHits(true),
	//	es.Search.WithPretty(),
	//)
	//
	//defer func() {
	//	_ = resp.Body.Close()
	//}()
	//bodyResp, _ := ioutil.ReadAll(resp.Body)
	//var body map[string]interface{}
	//_ = json.Unmarshal(bodyResp, &body)
	//hits, ok := body["hits"].(map[string]interface{})["hits"].([]interface{})
	//fmt.Println(ok)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Status())
	//fmt.Println(string(bodyResp))
	//fmt.Println(len(hits))
	//log.Println(strings.Repeat("-", 37))
	//
	//notilog, err := getNotificationLogFromQueryResp(resp)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(notilog)

	/*-----test update document by "doc"-------*/
	//now := time.Now()
	//modelUpdate := NotificationLogESModel{
	//	TemplateID: "123321",
	//	Data: Data{
	//		Title: "hahaha",
	//		Body: "hehehe",
	//	},
	//	//Devices: []DeviceESModel{
	//	//	{
	//	//		DeviceToken:  "113322",
	//	//		Platform:     "android",
	//	//		FCMMessageID: "aaaaaa",
	//	//		Status:       "Success",
	//	//	},
	//	//	{
	//	//		DeviceToken:  "123123",
	//	//		Platform:     "ios",
	//	//		FCMMessageID: "bbbb",
	//	//		Status:       "Success",
	//	//	},
	//	//},
	//	UpdatedAt: &now,
	//}
	//
	//updatingModel := map[string]interface{}{
	//	"doc": modelUpdate,
	//}
	//
	//dataUpdate, _ := json.Marshal(updatingModel)
	//resp, err := es.Update("test_timestamp", "1", bytes.NewReader(dataUpdate),
	//	es.Update.WithContext(context.Background()),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//bodyResp, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Status())
	//fmt.Println(string(bodyResp))
	//fmt.Println(strings.Repeat("-", 37))

	/*-----test update document by "script"-------*/

	//updateByScript := UpdateByScript{"ctx._source.remove('device_tokens')"}
	//dataUpdate, _ := json.Marshal(updateByScript)
	//resp, err := es.Update("bank", "2", bytes.NewReader(dataUpdate), es.Update.WithContext(context.Background()))
	//if err != nil {
	//	panic(err)
	//}
	//bodyResp, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Status())
	//fmt.Println(string(bodyResp))
	//log.Println(strings.Repeat("-", 37))

	/*-------- test create document---------*/
	now := time.Now().UTC()
	log := NotificationLogESModel{
		UserID: "a123123",
		Devices: []DeviceESModel{
			{
				DeviceToken:  "113322",
				Platform:     "android",
				FCMMessageID: "aaaaaa",
				Status:       "Success",
			},
			{
				DeviceToken:  "123123",
				Platform:     "ios",
				FCMMessageID: "bbbb",
				Status:       "Success",
			},
		},
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	createdDoc, _ := json.Marshal(log)
	resp, err := es.Create("db-sp-notification-logs-local", "1", bytes.NewReader(createdDoc),
		es.Create.WithContext(context.Background()),
		es.Create.WithPipeline("timestamp"),
	)
	if err != nil {
		panic(err)
	}
	bodyResp, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Status())
	fmt.Println(string(bodyResp))
	fmt.Println(strings.Repeat("-", 37))
}

type UpdateByDoc struct {
	Doc NotificationLogESModel `json:"doc"`
}

type UpdateByScript struct {
	Script interface{} `json:"script"`
}

type NotificationLogESModel struct {
	RequestID  string          `json:"request_id,omitempty" mapstructure:"request_id"`
	UserID     string          `json:"user_id,omitempty" mapstructure:"user_id"`
	AppID      string          `json:"app_id,omitempty" mapstructure:"app_id"`
	GroupID    string          `json:"group_id,omitempty" mapstructure:"group_id"`
	TemplateID string          `json:"template_id,omitempty" mapstructure:"template_id"`
	Data       Data            `json:"data" mapstructure:"data"`
	Devices    []DeviceESModel `json:"devices,omitempty" mapstructure:"devices"`
	CreatedAt  *time.Time      `json:"created_at,omitempty" mapstructure:"created_at"`
	UpdatedAt  *time.Time      `json:"updated_at,omitempty" mapstructure:"updated_at"`
}

type DeviceESModel struct {
	DeviceToken  string `json:"device_token" mapstructure:"device_token"`
	Platform     string `json:"platform" mapstructure:"platform"`
	FCMMessageID string `json:"fcm_message_id" mapstructure:"fcm_message_id"`
	Status       string `json:"status" mapstructure:"status"`
}

type Data struct {
	CTADeepLink string `json:"cta_deeplink,omitempty" mapstructure:"cta_deeplink"`
	CTAText     string `json:"cta_text,omitempty" mapstructure:"cta_text"`
	Title       string `json:"title,omitempty" mapstructure:"title"`
	Body        string `json:"body,omitempty" mapstructure:"body"`
	DeepLink    string `json:"deeplink,omitempty" mapstructure:"deeplink"`
	BrandName   string `json:"brand_name,omitempty" mapstructure:"brand_name"`
	ItemImage   string `json:"item_image,omitempty" mapstructure:"item_image"`
	ImageUrl    string `json:"image_url,omitempty" mapstructure:"image_url"`
}

func getNotificationLogFromQueryResp(response *esapi.Response) (NotificationLogESModel, error) {
	bodyResp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return NotificationLogESModel{}, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyResp, &data)
	if err != nil {
		return NotificationLogESModel{}, err
	}

	hitsData, ok := data["hits"].(map[string]interface{})
	if !ok {
		return NotificationLogESModel{}, errors.New("error 1")
	}
	hits, ok := hitsData["hits"].([]interface{})
	if !ok {
		return NotificationLogESModel{}, errors.New("error 2")
	}
	if len(hits) == 0 {
		return NotificationLogESModel{}, errors.New("document not found")
	}
	doc, ok := hits[0].(map[string]interface{})
	if !ok {
		return NotificationLogESModel{}, errors.New("error 3")
	}

	var result NotificationLogESModel
	err = mapstructure.Decode(doc["_source"], &result)
	if err != nil {
		return NotificationLogESModel{}, err
	}
	return result, nil
}
