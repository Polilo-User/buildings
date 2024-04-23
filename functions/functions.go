package functions

import (
	"bytes"
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	//uuid "github.com/nu7hatch/gouuid"
)

// Ф-ция предназначенная для запроса в БД (этой функции можно передавать args)
func Query2(db *sql.DB, str string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(str, args...)
	defer rows.Close()
	if err != nil {
		PrintLog("err", err)
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	return makeData(rows, columns), nil
}

func makeData(rows *sql.Rows, columns []string) []map[string]interface{} {
	values := make([]interface{}, len(columns))   // массив, в котором будут значения с ответа от SQL запроса
	scanArgs := make([]interface{}, len(columns)) // массив, элементы которого будут указателями на элементы массива со значениями... хоспадее...
	for i := range values {                       // определяем элементы второго массива как указатели на элементы первого
		scanArgs[i] = &values[i]
	}
	resultArray := make([]map[string]interface{}, 0)
	for rows.Next() {

		data := make(map[string]interface{})
		err := rows.Scan(scanArgs...) //тут передаем массив с указателями, через который мы получим значения в массив со значениями...
		// Если попытаться считать сразу в массив со значениями, то ругается, говорит дай мне указатели!!!
		if err != nil { // Если не удалось запихать значения в массив
			fmt.Println("error scan: ", err)
		}
		for index, v := range values { // Идем по массиву со значениями
			data[columns[index]] = v
		}
		resultArray = append(resultArray, data)
	}
	return resultArray
}

func CodeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

// Intf2str - convert interface to string
func Intf2str(i interface{}) string {
	iType := Typeof(i)
	if iType == "float64" || iType == "float32" {
		return Intf2jsonStr(i)
	} else if iType == "<nil>" {
		return ""
	}
	return fmt.Sprintf("%v", i)
}

func Intf2Int(i interface{}) (int, error) {
	return strconv.Atoi(Intf2str(i))
}

// ArrIntf2str - convert []interface{} to string
func ArrIntf2str(mark string, arr []interface{}) string {
	str := ""
	for _, v := range arr {
		str += mark + Intf2str(v)
	}
	return strings.TrimLeft(str, mark)
}

// Массив интов в строку
// mark - разделитель
func ArrInt2Str(mark string, arr []int) string {
	str := ""
	for _, v := range arr {
		str += mark + strconv.Itoa(v)
	}
	return strings.TrimLeft(str, mark)
}

func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Intf2jsonStr - переводим интерыейс в строку формата json
func Intf2jsonStr(m interface{}) string {
	if b, bErr := json.Marshal(m); bErr == nil {
		return string(b)
	} else {
		log.Println("Intf2jsonStr:err:", bErr)
	}
	return ""
}

func Str2f64(f64str string) float64 {
	f64, _ := strconv.ParseFloat(f64str, 64)
	return f64
}

// SIntf2str - convert []string to string
func SStr2str(mark string, arr []string) string {
	str := ""
	for _, v := range arr {
		str += mark + v
	}
	return strings.TrimLeft(str, mark)
}

func PrintLog(v ...interface{}) {
	log.Println(v...)
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.//use 3072
	buffer := make([]byte, 3072)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	if contentType == "application/zip" {
		contentType = mimetype.Detect(buffer).String()
	}

	return contentType, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type FileData struct {
	FileName string
	FileHash string
	FileUrl  string
	FileSize int
}

// Сохраняет файл в директорию /var/www/+path
func SaveFilesFromMultipart(fh *multipart.FileHeader, path string) (fileData *FileData, err error) {

	fileName := fh.Filename // Вытаскиваем имя файла
	//fileHeader := fh.Header
	fileSize := int(fh.Size) // размер файла

	f, err := fh.Open() // Открываем файл из мультипарта
	if err != nil {
		return nil, err
	}
	defer f.Close() // Закрываем соединение с файлом
	if err != nil {
		return fileData, err
	} else {
		buf := bytes.NewBuffer(nil)                // Создадим байт буффер
		if _, err := io.Copy(buf, f); err == nil { // Копируем побайтно файл в буффер

			bts := buf.Bytes() // Достаем байты из буффера

			sha := sha1.New()                       // Создаем новый sha1 хэшер
			sha.Write(bts)                          // Пихаем в него байты нашего файла
			hash := sha.Sum(nil)                    // Получаем хэш сумму
			hashStr := fmt.Sprintf("%x", hash)      // Переводим хэш в строку
			fileType := http.DetectContentType(bts) // Определим Контент тип файла

			typeSplit := strings.Split(fileType, "/")
			ext := typeSplit[len(typeSplit)-1] // Вытаскиваем последнее слово из content type файла. (Типо pdf, jpeg и тд.)
			if ext == "jpeg" {
				ext = "jpg" // Нам нужен jpg, а не jpeg
			}
			if ext == "zip" {
				ext = strings.TrimLeft(mimetype.Detect(bts).Extension(), ".")
			}
			if strings.Contains(ext, "html") {
				ext = "html" //Нам нужен html, а не "text/html; charset=utf-8"
			}
			newFile, err := os.OpenFile("/var/www/"+path+hashStr+"."+ext, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Создаем или открываем (если уже есть файл с таким хэшем) файл для записи
			defer newFile.Close()
			if err == nil {
				newFile.Truncate(0)
				newFile.Seek(0, 0)
				if _, err = newFile.Write(bts); err == nil { // Пишем байты в наш файл
					fileData := FileData{fileName, hashStr, path + hashStr + "." + ext, fileSize}
					return &fileData, nil
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
}

func RemoveNoNumsFromStr(str string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(str, "")
}

// Возвращает sql.NullString если пустая строка
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func GetFilesFromMultipart(MultipartForm *multipart.Form) []*multipart.FileHeader {
	var files []*multipart.FileHeader
	if MultipartForm != nil {
		files = MultipartForm.File["pictures"]
		if files == nil {
			files = MultipartForm.File["pictures[]"]
		}
		if files == nil {
			files = MultipartForm.File["files"]
		}
		if files == nil {
			files = MultipartForm.File["files[]"]
		}
		if files == nil {
			files = MultipartForm.File["file"]
		}
		if files == nil {
			files = MultipartForm.File["file[]"]
		}
	}
	return files
}

type TypeFileResponse struct {
	File TypeFile `json:"file"`
}

type TypeFile struct {
	FilePath    string `json:"filePath,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Size        int    `json:"size,omitempty"`
}

func GetNullString(str string) sql.NullString {
	var strNullable sql.NullString
	if len(str) > 0 {
		strNullable.Scan(str)
	} else {
		strNullable.Scan(nil)
	}
	return strNullable
}

// Для удобства работы с postgres. Заменяет все "?" на "$1", "$2"...
func AdaptReq(s string) string {
	i := 1
	for {
		s = strings.Replace(s, "?", "$"+strconv.Itoa(i), 1)
		i++
		if !strings.Contains(s, "?") {
			break
		}
	}
	return s
}
