package downloader

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const WORKER_NUM = 10
const MAX_WAIT_TIME = 120 * time.Second

var (
	NO_CONTENT_LENGTH = errors.New("NO_CONTENT_LENGTH")
	CHUNK_RES_EMPTY   = errors.New("CHUNK_RES_EMPTY")
	WAIT_TOO_LONG     = errors.New("WAIT_TOO_LONG")
)

func master(filename, url string) error {
	meta, err := getContentMeta(url)
	if err != nil {
		return err
	}
	count := 0
	chunkList := divideContentChunk(meta.size, WORKER_NUM)
	fmt.Printf("chunkList: %v\n", chunkList)
	totalChunk := make([][]byte, WORKER_NUM)
	payloadChunk := make(chan *payload, WORKER_NUM/2)
	for index, _ := range chunkList {
		go downloadWorker(index, url, chunkList[index], payloadChunk)
	}

	done := time.After(MAX_WAIT_TIME)

Loop:
	for {
		select {
		case payload := <-payloadChunk:
			if payload.err != nil {
				return payload.err
			} else {
				index := payload.index
				totalChunk[index] = payload.data
				count += 1
				if count == WORKER_NUM {
					break Loop
				}
			}
		case <-done:
			return WAIT_TOO_LONG
		}
	}
	buffer := combineChunk(totalChunk)

	err = writeFile("", filename, meta.contentType, buffer.Bytes())

	return err
}

func getContentMeta(url string) (*contentMeta, error) {
	resp, err := client.SetRetryCount(3).SetRetryWaitTime(3 * time.Second).R().Head(url)
	if err != nil {
		return nil, err
	}
	meta := &contentMeta{}
	length := resp.Header().Get("Content-Length")
	if length == "" {
		return nil, NO_CONTENT_LENGTH
	}
	meta.size, err = strconv.Atoi(length)
	if err != nil {
		return nil, NO_CONTENT_LENGTH
	}
	contentType := resp.Header().Get("Content-Type")
	split := strings.Split(contentType, "/")
	if len(split) > 0 {
		meta.contentType = split[1]
	}
	return meta, nil
}

func divideContentChunk(size, workers int) [][]int {
	out := make([][]int, workers)
	chunk := size / workers

	for i := 0; i < workers; i++ {
		out[i] = make([]int, 2)
		if i == 0 {
			out[i][0] = 0
			out[i][1] = chunk
		} else {
			start := out[i-1][1] + 1
			out[i][0] = start
			next := start + chunk
			if next > size {
				out[i][1] = size
			} else {
				out[i][1] = next
			}
		}
	}
	return out
}

func downloadWorker(index int, url string, chunkRange []int, channel chan *payload) {
	payload := &payload{}
	Range := fmt.Sprintf("bytes=%d-%d", chunkRange[0], chunkRange[1])
	resp, err := client.SetRetryCount(3).SetRetryWaitTime(3*time.Second).R().SetHeader("Range", Range).Get(url)
	if err != nil {
		payload.err = err
		channel <- payload
		return
	}
	if resp == nil {
		payload.err = CHUNK_RES_EMPTY
		channel <- payload
		return
	}
	payload.index = index
	payload.data = resp.Body()
	fmt.Printf("chunk %d done, data: %v\n", index, len(payload.data))
	channel <- payload
}

func combineChunk(bytesList [][]byte) *bytes.Buffer {
	buffer := &bytes.Buffer{}

	for index, _ := range bytesList {
		buffer.Write(bytesList[index])
	}
	return buffer
}

func writeFile(path, filename, contentType string, data []byte) error {
	err := ioutil.WriteFile(path+filename+"."+contentType, data, 777)
	return err
}

type WriteCounter struct {
	Current int
	Total   int
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Current += n
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fc := float32(wc.Current)
	ft := float32(wc.Total)
	fmt.Printf("\rDownloading... %.2f%% complete\n", (fc/ft)*100)
}

func displayProgress(filename, url string) error {
	meta, err := getContentMeta(url)
	if err != nil {
		return err
	}
	out, err := os.Create(filename + "." + meta.contentType)
	if err != nil {
		return err
	}
	defer out.Close()
	wc := &WriteCounter{0, meta.size}
	// resp, err := resty.New().R().Get(url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, wc))
	if err != nil {
		return err
	}
	return nil
}
