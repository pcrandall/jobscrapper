package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	Id       string
	Title    string
	Location string
	Salary   string
	Summary  string
}

var (
	config    SearchConfig
	baseURL   = "https://www.indeed.com/jobs?"
	baseQuery = "q=golang"
	baseLimit = "&limit=50"
	URL       = baseURL + baseQuery + baseLimit
	jobs      []extractedJob
	// urls      []string
)

func main() {
	GetConfig()
	var urls []string
	c := make(chan []extractedJob)

	for _, val := range config.Jobs {
		// fmt.Println("val here", val)
		for _, v := range val.Location {
			// fmt.Println(v)
			urls = append(urls, config.Baseurl+val.Keyword+"&"+v+config.Baselimit)
		}
	}

	for _, url := range urls {
		totalPages := getPages(url)
		for i := 0; i < totalPages; i++ {
			go getPage(url, i, c)
			// extractedJobs := getPage(i)
			// jobs = append(jobs, extractedJobs...)
		}

		for i := 0; i < totalPages; i++ {
			extractedJobs := <-c
			jobs = append(jobs, extractedJobs...)
		}
	}

	writeJobs(jobs)

	fmt.Println("Finished! Extracted", len(jobs), "jobs!")
	fmt.Println("Serving at http://localhost:8888")

	http.HandleFunc("/", genHTML)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status:", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0

	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

// c channel send only type slice extractedJob
func getPage(url string, page int, mainChannel chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*50)
	// fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		// jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	// return jobs
	mainChannel <- jobs
}

// channel send only type extractJob
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	c <- extractedJob{
		Id:       id,
		Title:    title,
		Location: location,
		Salary:   salary,
		Summary:  summary}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	err = w.Write(headers)
	checkErr(err)

	for _, job := range jobs {
		job.Id = "https://indeed.com/viewjob?jk=" + job.Id
		jobSlice := []string{job.Id, job.Title, job.Location, job.Salary, job.Summary}
		err = w.Write(jobSlice)
		checkErr(err)
	}
}

func genHTML(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Underlying Type: %T\n", jobs)
	// fmt.Printf("Underlying Value: %v\n", jobs)
	t, err := template.ParseFiles("layout.html")
	checkErr(err)
	err = t.Execute(w, jobs)
	checkErr(err)

}
