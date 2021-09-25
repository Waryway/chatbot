package handlers

import (
	"bufio"
	"fmt"
	"os"
)

type History struct {
	Location string
	Current  []string
	file     *os.File
}

func (h *History) Initialize(location string) *History {
	h.Location = location
	return h
}

func (h *History) LoadHistory() error {

	err := h.open()
	if err != nil {
		return err
	}

	defer func() { _ = h.close() }()
	// Slice for storing the History of the world.
	scanner := bufio.NewScanner(h.file) // use the buffer input output package to start a file scanner
	for scanner.Scan() {                // loop through the file scanner
		h.addRecord(scanner.Text() + "\r\n") // add the scanned line to the lines slice.
	}

	return nil
}

func (h *History) Record(record string, isPrintable bool) {
	if isPrintable {
		fmt.Print(record + "\r\n")
	}
	h.addRecord(record + "\r\n")
	h.writeRecord(record + "\r\n")
}

func (h *History) addRecord(record string) {
	h.Current = append(h.Current, record)
}

func (h *History) writeRecord(record string) {
	_ = h.open()
	defer func() { _ = h.close() }()
	_, _ = h.file.WriteString(record)
}

func (h History) SaveAll() {
	for _, record := range h.Current {
		h.writeRecord(record)
	}
}

/**
 * This will perform the three error producing steps to erase the History file.
 * 1. Close the writer
 * 2. Truncate the History file and the local History
 * 3. Re-open the file
 */
func (h *History) EraseHistory() error {
	err := h.close()
	if err == nil {
		err = os.Truncate(h.Location, 0)
	}
	if err == nil {
		h.Current = []string{}
		err = h.open()
	}

	return err
}

func (h History) close() error {
	return h.file.Close()
}

func (h *History) open() error {
	f, err := os.OpenFile(h.Location, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	h.file = f
	// check for the error.
	if err != nil { // a 'nil' error means no error occurred.
		fmt.Println(err) // print the error
		return err       // stop processing the context.
	}

	return nil
}

func (h History) Length() int {
	return len(h.Current)
}

func (h History) Print() {
	// iterate (loop) through the history
	for k, v := range h.Current {
		// print out the Key and the Value
		fmt.Print(k, v)
	}
}
