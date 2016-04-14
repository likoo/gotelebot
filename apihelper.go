package gotelebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"

	"github.com/likoo/gotelebot/types"
)

func sendGetRequest(method string, token string, params url.Values) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?%s",
		token, method, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return checkResult(resp)
}

func sendRequest(method, token, name, path string, params url.Values) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		part, err := writer.CreateFormFile(name, filepath.Base(path))
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, file); err != nil {
			return nil, err
		}
	}
	for field, values := range params {
		if len(values) > 0 {
			writer.WriteField(field, values[0])
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkResult(resp)

}

func checkResult(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != 200 {
		con, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("gotelebot error:%s-%s", resp.Status, con))
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(jsonStr, &result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("gotelebot The server returned an invalid JSON response. %s-%s",
			resp.Status, resp.Body))
	}
	if result["ok"] != true {
		return nil, errors.New(fmt.Sprintf("gotelebot: Error.ErrorCode: %s-Description%s",
			result["errorCode"], result["description"]))
	}
	str, errs := json.Marshal(result["result"])
	if errs != nil {
		fmt.Println("Error encoding JSON")
		return nil, errors.New(fmt.Sprintln("gotelebot"))
	}
	return []byte(str), nil
}

func makeRequest(method, token, name, filepath string, params url.Values) ([]byte, error) {
	return sendRequest(method, token, name, filepath, params)
}

func getMe(token string) (*types.User, error) {
	jsonStr, err := sendGetRequest("getMe", token, url.Values{})
	if err != nil {
		return nil, err
	}
	var user types.User
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUpdates(token, offset, limit, timeout string) ([]*types.Update, error) {
	payload := url.Values{}
	if offset != "" {
		payload.Add("offset", offset)
	}
	if limit != "" {
		payload.Add("limit", limit)
	}
	if timeout != "" {
		payload.Add("timeout", timeout)
	}
	jsonStr, err := sendGetRequest("getUpdates", token, payload)
	if err != nil {
		return nil, err
	}
	var result []*types.Update
	err = json.Unmarshal(jsonStr, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getUserProfilePhotos(token, userid string, opt *GetUserProfilePhotosOptional) (*types.UserProfilePhotos, error) {
	payload := url.Values{}
	payload.Add("user_id", userid)
	if opt != nil {
		opt.AppendPayload(&payload)
	}
	jsonStr, err := makeRequest("getUserProfilePhotos", token, "", "", payload)
	if err != nil {
		return nil, err
	}
	var ups types.UserProfilePhotos
	err = json.Unmarshal(jsonStr, &ups)
	if err != nil {
		return nil, err
	}
	return &ups, nil
}

func sendMessage(token, chat_id, text string, opt *SendMessageOptional) (*types.Message, error) {
	payload := url.Values{}
	payload.Add("chat_id", chat_id)
	payload.Add("text", text)
	if opt != nil {
		opt.AppendPayload(&payload)
	}
	jsonStr, err := makeRequest("sendMessage", token, "", "", payload)
	if err != nil {
		return nil, err
	}
	return transformToMessage(jsonStr)
}

func forwardMessage(token, chat_id, from_chat_id, message_id string) (*types.Message, error) {
	payload := url.Values{}
	payload.Add("chat_id", chat_id)
	payload.Add("from_chat_id", from_chat_id)
	payload.Add("message_id", message_id)
	jsonStr, err := makeRequest("forwardMessage", token, "", "", payload)
	var msg types.Message
	err = json.Unmarshal(jsonStr, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func sendPhoto(token, chat_id, photo string, opt *SendPhotoOptional) (*types.Message, error) {
	return sendFile(token, chat_id, "sendPhoto", "photo", photo, opt)
}

func sendAudio(token, chat_id, audio string, opt *SendAudioOptional) (*types.Message, error) {
	return sendFile(token, chat_id, "sendAudio", "audio", audio, opt)
}

func sendDocument(token, chat_id, document string, opt *SendDocumentOptional) (*types.Message, error) {
	return sendFile(token, chat_id, "sendDocument", "document", document, opt)
}

func sendSticker(token, chat_id, sticker string, opt *SendStickerOptional) (*types.Message, error) {
	return sendFile(token, chat_id, "sendSticker", "sticker", sticker, opt)
}

func sendVideo(token, chat_id, video string, opt *SendVideoOptional) (*types.Message, error) {
	return sendFile(token, chat_id, "sendVideo", "video", video, opt)
}

func sendVoice(token, chat_id, voice string, opt *SendVoiceOptional) (*types.Message, error) {
	return sendFile(token, chat_id, "sendVoice", "voice", voice, opt)
}

func sendLocation(token, chat_id, latitude, longitude string, opt *SendLocationOptional) (*types.Message, error) {
	payload := url.Values{}
	payload.Add("chat_id", chat_id)
	payload.Add("latitude", latitude)
	payload.Add("longitude", longitude)
	if opt != nil {
		opt.AppendPayload(&payload)
	}
	jsonStr, err := makeRequest("sendLocation", token, "", "", payload)
	if err != nil {
		return nil, err
	}
	return transformToMessage(jsonStr)
}

func sendChatAction(token, chat_id, action string) (string, error) {
	payload := url.Values{}
	payload.Add("chat_id", chat_id)
	payload.Add("action", action)
	jsonStr, err := makeRequest("sendChatAction", token, "", "", payload)
	if err != nil {
		return "", err
	}
	ret := string(jsonStr[:])
	return ret, nil
}

func getFile(token, fileId string) (*types.File, error) {
	payload := url.Values{}
	payload.Add("file_id", fileId)
	jsonStr, err := makeRequest("getFile", token, "", "", payload)
	if err != nil {
		return nil, err
	}
	var file types.File
	err = json.Unmarshal(jsonStr, &file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func downloadFile(token, filePath string) (*[]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s",
		token, filePath)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("downloadFile error.statue code :%d", resp.StatusCode))
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func sendFile(token, chat_id, methodname, typename, file string, opt Optional) (*types.Message, error) {
	payload := url.Values{}
	filepath := ""
	formname := ""
	payload.Add("chat_id", chat_id)
	if _, err := os.Stat(file); err == nil {
		filepath = file
		formname = typename
	}
	if filepath == "" { // Use telegram fileid
		payload.Add(typename, file)
	}
	if !reflect.ValueOf(opt).IsNil() { // Check interface conatain nil
		opt.AppendPayload(&payload)
	}
	jsonStr, err := makeRequest(methodname, token, formname, filepath, payload)
	if err != nil {
		return nil, err
	}
	return transformToMessage(jsonStr)
}

func answerInlineQuery(token, inlineQueryId string, results []interface{}, opt *AnswerInlineQueryOptional) (bool, error) {
	payload := url.Values{}
	payload.Add("inline_query_id", inlineQueryId)
	resultJosn, err := getInlineQueryResultJsonString(results)
	if err != nil {
		return false, err
	}
	payload.Add("results", resultJosn)
	if !reflect.ValueOf(opt).IsNil() { // Check interface conatain nil
		opt.AppendPayload(&payload)
	}
	jsonStr, err := makeRequest("answerInlineQuery", token, "", "", payload)
	if err != nil {
		return false, err
	}
	var ret bool
	err = json.Unmarshal(jsonStr, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func getInlineQueryResultJsonString(results []interface{}) (string, error) {
	ret := ""
	for _, r := range results {
		jsonStr, err := json.Marshal(r)
		if err != nil {
			return "", errors.New(fmt.Sprint("InlineQueryResult json encode error. %s. %#v", err, r))
		}
		ret = ret + string(jsonStr) + ","
	}
	if len(ret) > 0 {
		ret = string(ret[0 : len(ret)-1])
	}
	return "[" + ret + "]", nil
}

func transformToMessage(jsonStr []byte) (*types.Message, error) {
	var msg types.Message
	err := json.Unmarshal(jsonStr, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
