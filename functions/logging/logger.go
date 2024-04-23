package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Polilo-User/buildings/functions/errors"
)

// Конфигурация логгера
type loggerSettings struct {
	Debug     *bool
	traseFile *os.File
	errorFile *os.File
}

// Структура общего логгера, чтобы можно было легко заменить его
type Logger struct {
	*loggerSettings
}

var S = &loggerSettings{}

// Создаем конфигуратор логгера
func Init(d bool) {

	S.Debug = &d

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	if !*S.Debug {
		err := os.MkdirAll(dir+"/logs", 0777)
		if err != nil {
			panic(err)
		}
		currentTime := time.Now().Format("2006-01-02_15-04-05")

		S.errorFile, err = os.OpenFile(dir+"/logs/error_"+currentTime+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777) // создание файла
		if err != nil {
			panic(err)
		}
	}
}

// Функция получения логгера из любого места программы
func GetLogger() *Logger {
	return &Logger{S}
}

// Структура со всеми возможными полями для вывода логов
type logStruct struct {
	path  string
	msg   string
	time  string
	level string

	errorText string
	context   string
}

func (logger *Logger) Error(err error) {
	printErrorLog("error", err)
}

func (logger *Logger) Warning(err error) {
	printErrorLog("warning", err)
}

func (logger *Logger) Info(msg string, args ...interface{}) {
	printLog("info", msg, args...)
}

func (logger *Logger) Fatal(err error) {
	printErrorLog("fatal", err)
	os.Exit(1)
}

func (logger *Logger) Debug(msg string, args ...interface{}) {
	printLog("debug", msg, args...)
}

func printErrorLog(level string, err error) {
	//Приведение ошибки к нашему типу
	customErr, ok := err.(errors.CustomError)
	if !ok {
		fmt.Println("Failed to cast error type to customError")
	}
	//Получаем текст первоначальной ошибки
	customErr.DevelopText = customErr.Error()
	//Переносим данные от ошибки к структуре лога
	var values logStruct
	values.path = customErr.Path
	values.context = customErr.Context
	values.errorText = customErr.DevelopText
	values.level = level
	values.time = time.Now().Format("2006-01-02 15:04:05")
	//Записываем лог в консоль
	consoleLog := getConsoleLog(values)
	println(consoleLog)
	//Записываем лог в файл
	if !*S.Debug {
		fileLog := getFileLog(values)
		writeInFile(values.level, fileLog)
	}
}

func printLog(level, msg string, arrs ...interface{}) {

	_, file, line, _ := runtime.Caller(2)

	var values logStruct
	values.path = file + ":" + strconv.Itoa(line)
	values.msg = fmt.Sprintf(msg, arrs...)
	values.level = level
	values.time = time.Now().Format("2006-01-02 15:04:05")

	consoleLog := getConsoleLog(values)

	println(consoleLog)

	if !*S.Debug {
		fileLog := getFileLog(values)
		writeInFile(values.level, fileLog)
	}
}

func getConsoleLog(values logStruct) (log string) {

	if !*S.Debug {
		pathArr := strings.Split(values.path, "adminServer/")
		if len(pathArr) > 1 {
			values.path = pathArr[1]
		}
	}

	switch values.level {
	case "error":
		log = "\x1b[31m[ERROR] \x1b[0m"
	case "fatal":
		log = "\x1b[35m[FATAL] \x1b[0m"
	case "info":
		log = "\x1b[34m[INFO] \x1b[0m"
	case "debug":
		log = "\x1b[36m[DEBUG] \x1b[0m"
	case "warning":
		log = "\x1b[33m[WARN] \x1b[0m"
	}

	switch values.level {
	case "error", "fatal", "warning":
		log += fmt.Sprintf("%s %s", values.path, values.errorText)
		if values.context != "" {
			log += fmt.Sprintf(". context: %s", values.context)
		}
	default:
		log += fmt.Sprintf("%s %s", values.path, values.msg)
	}

	return log
}

func getFileLog(values logStruct) (log string) {

	pathArr := strings.Split(values.path, "adminServer/")
	if len(pathArr) > 1 {
		values.path = pathArr[1]
	}

	log += fmt.Sprintf("\ntime=\"%s\" level=\"%s\" msg=\"%s\" ", values.time, values.level, values.errorText)
	switch values.level {
	case "error", "fatal", "warning":
		if values.context != "" {
			log += fmt.Sprintf("context=\"%s\" ", values.context)
		}
		log += fmt.Sprintf("path=\"%s\" ", values.path)
	default:
		log += fmt.Sprintf("path=\"%s\" ", values.path)
	}

	return log

}

func writeInFile(level, log string) error {

	switch level {
	case "error", "fatal":
		_, err := S.errorFile.WriteString("\n" + log)
		if err != nil {
			return err
		}
	}

	return nil
}
