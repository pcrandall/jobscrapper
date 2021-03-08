package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
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
	Date     string
	FullDesc string
}

var (
	config  SearchConfig
	jobs    []extractedJob
	query   bool
	logfile *os.File
)

func main() {
	flag.BoolVar(&query, "q", false, "query indeed.com, if false build site from ./sites/jobs.csv -q=true")
	flag.Parse()
	if query == false {
		lines, err := ReadCsv("./site/jobs.csv")
		checkErr(err)
		// Loop through lines & turn into object
		for _, line := range lines {
			data := extractedJob{
				Id:       line[0],
				Title:    line[1],
				Location: line[2],
				Salary:   line[3],
				Summary:  line[4],
				Date:     line[5],
				FullDesc: line[6],
			}
			jobs = append(jobs, data)
		}
		serveJobs()
		os.Exit(0)
	}

	logfile, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)
	defer logfile.Close()

	GetConfig()
	var urls []string
	c := make(chan []extractedJob)

	for _, job := range config.Jobs {
		keyword := strings.TrimSpace(job.Keyword)
		keyword = "q=" + strings.ReplaceAll(keyword, " ", "+")
		if len(job.Location) == 0 {
			fmt.Println("Finding", job.Keyword)
			urls = append(urls, config.Baseurl+keyword+config.Baselimit)
		}
		for _, location := range job.Location {
			loc := strings.TrimSpace(location)
			loc = "l=" + strings.ReplaceAll(loc, " ", "%2C+")
			urls = append(urls, config.Baseurl+keyword+"&"+loc+config.Baselimit)
			fmt.Println("Finding", job.Keyword, "jobs in", location)
		}
	}

	//TODO make channgel and go func here
	for _, url := range urls {
		totalPages := getPages(url)
		// fmt.Println("totalPages here:", totalPages)
		pageLimit := config.Maxresults / 50
		// fmt.Println("pageLimit here:", pageLimit)
		if pageLimit > totalPages {
			pageLimit = totalPages
		}
		// fmt.Println("pageLimit here:", pageLimit)
		for i := 0; i < pageLimit; i++ {
			go getPage(url, i, c)
		}
		for i := 0; i < pageLimit; i++ {
			extractedJobs := <-c
			jobs = append(jobs, extractedJobs...)
		}
	}

	RemoveDuplicates(jobs)
	writeJobs(jobs)
	fmt.Println("Finished! Extracted", len(jobs), "jobs!")
	serveJobs()
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

func cleanFullDesc(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "####")
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

	// TODO make this better
	// doc.Find("#searchCountPages").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println("searchCountPages text here", s.Text())
	// })
	return pages
}

// c channel send only type slice extractedJob
func getPage(url string, page int, mainChannel chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)

	fmt.Println("Requesting", pageURL)
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

	mainChannel <- jobs // return jobs
}

// channel send only type extractJob
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	id = "https://www.indeed.com/viewjob?jk=" + id
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	date := cleanString(card.Find(".date").Text())

	fullDescription := make(chan string)
	go getFullDescription(id, fullDescription)
	fulldesc := <-fullDescription

	c <- extractedJob{
		Id:       id,
		Title:    title,
		Location: location,
		Salary:   salary,
		Summary:  summary,
		Date:     date,
		FullDesc: fulldesc,
	}

}

func getFullDescription(url string, description chan<- string) {
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// d := doc.Find("#jobDescriptionText")
	// checkType(d)

	des := doc.Find("#jobDescriptionText").Text()
	description <- des

	// card := doc.Find("#jobDescriptionText").Text()
	// // fmt.Println("Card HERE: ", card)
	// // fmt.Printf("%+v\n", card)
	// // d := cleanString(doc.Find(".jobsearch-JobDescriptionText").Text())
	// d := cleanString(card)

	// d := doc.Find("#jobDescriptionText")

	// d.Contents().Each(func(i int, s *goquery.Selection) {
	// 	// if goquery.NodeName(s) == "#text" {
	// 	// 	fmt.Println(s.Text())
	// 	// }
	// 	fmt.Println(s)
	// })
	// checkType(d)

	// des := doc.Find("#jobDescriptionText").Text()
}

func RemoveDuplicates(jobs []extractedJob) {
	found := make(map[string]bool)
	j := 0
	for i, x := range jobs {
		if !found[x.Summary] {
			found[x.Summary] = true
			(jobs)[j] = (jobs)[i]
			j++
		}
	}
	jobs = (jobs)[:j]
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("./site/jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()
	// headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	// err = w.Write(headers)
	// checkErr(err)
	for _, job := range jobs {
		jobSlice := []string{job.Id, job.Title, job.Location, job.Salary, job.Summary, job.Date, job.FullDesc}
		err = w.Write(jobSlice)
		checkErr(err)
	}
}

func serveJobs() {
	listener, err := net.Listen("tcp", ":0")
	checkErr(err)
	// handle `/static` route
	http.HandleFunc("/", serveTemplate)
	http.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("./site/static"))),
	)
	http.Handle("/static/images/",
		http.StripPrefix("/static/images", http.FileServer(http.Dir("./site/static/images"))),
	)
	fmt.Printf("Serving at http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)
	log.Fatal(http.Serve(listener, nil))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	t, err := template.ParseFiles("./site/layout.html")
	checkErr(err)
	err = t.Execute(w, jobs)
	checkErr(err)
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func checkType(t interface{}) {

	switch reflect.TypeOf(t).Kind() {

	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}

	case reflect.Array:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}

	case reflect.Chan:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}

	case reflect.Func:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}

	case reflect.Map:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}

	default:
		fmt.Printf("Underlying Type: %T\n", t)
		fmt.Printf("Underlying Value: %v\n", t)
	}

	// for k, _ := range s {
	// 	fmt.Printf("Underlying Type: %T\n  Underlying Value: %v\n", s[k], s[k])
	// }

	// fmt.Printf("Underlying Type: %T\n", jobs)
	// fmt.Printf("Underlying Value: %v\n", jobs)

	// fmt.Printf("%T %v\n", s, s)
}
