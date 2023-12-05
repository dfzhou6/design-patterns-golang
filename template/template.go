package template

import "fmt"

type FileProcessor interface {
	OpenFile()
	ReadFile()
	ProcessData()
	CloseFile()
}

type CsvFile struct {
}

func NewCsvFile() *CsvFile {
	return &CsvFile{}
}

func (impl *CsvFile) OpenFile() {
	fmt.Println("csv open file")
}

func (impl *CsvFile) ReadFile() {
	fmt.Println("csv read file")
}

func (impl *CsvFile) ProcessData() {
	fmt.Println("csv process data")
}

func (impl *CsvFile) CloseFile() {
	fmt.Println("csv close file")
}

type XmlFile struct {
}

func NewXmlFile() *XmlFile {
	return &XmlFile{}
}

func (impl *XmlFile) OpenFile() {
	fmt.Println("xml open file")
}

func (impl *XmlFile) ReadFile() {
	fmt.Println("xml read file")
}

func (impl *XmlFile) ProcessData() {
	fmt.Println("xml process data")
}

func (impl *XmlFile) CloseFile() {
	fmt.Println("xml close file")
}

type FileHandler struct {
	fileProcessor FileProcessor
}

func NewFileHandler(fileProcessor FileProcessor) *FileHandler {
	return &FileHandler{fileProcessor: fileProcessor}
}

func (impl *FileHandler) Handle() {
	impl.fileProcessor.OpenFile()
	impl.fileProcessor.ReadFile()
	impl.fileProcessor.ProcessData()
	impl.fileProcessor.CloseFile()
}

func (impl *FileHandler) SetProcessor(fileProcessor FileProcessor) {
	impl.fileProcessor = fileProcessor
}
