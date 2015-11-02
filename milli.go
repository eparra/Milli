package main

import (
    "fmt"
    "os" 
    "time"
    "sync" 
    "encoding/csv"
    "github.com/parnurzeal/gorequest"
)

const workersCount = 24

func getUrlWorker(urlChan chan string) {
    for url := range urlChan {
        resp, body, errs := gorequest.New().Get(url).End()
        _ = resp
        _ = body
        _ = errs
    }
}

func main() {

    // Record start time.  Will be used for  
    // detemine elapsed time once completed  
    start := time.Now()

    // Open Alexa CSV file  
    fmt.Println("Opening top-1m.csv file")
    csvfile, err := os.Open("top-1m.csv")
    if err != nil {
        panic(err)
    } else {
        defer csvfile.Close()
    }
     
    // Read in CSV file
    reader := csv.NewReader(csvfile)
    reader.FieldsPerRecord = -1 
    fmt.Println("  Reading Entries (May take a few seconds)")
    rawCSVdata, err := reader.ReadAll()
    if err != nil {
        panic(err)
    } else {
        fmt.Println("  Completed!")
    }

    // Start concurrency management 
    var wg sync.WaitGroup
    urlChan := make(chan string)

    wg.Add(workersCount)

    for i := 0; i < workersCount; i++ {
        go func() {
            getUrlWorker(urlChan)
            wg.Done()
        }()
    }

    completed := 0
    for _, each := range rawCSVdata {
        url := fmt.Sprintf("http://%s", each[1]) 
        urlChan <- url
        completed++
        fmt.Printf("Completed %s (%d)\n", url, completed)
    }
    close(urlChan)

    wg.Wait()

    // End summary
    fmt.Println("Completed Tranactoins : ", completed)
    fmt.Println("Time Elapsed          : ", time.Since(start))

}









