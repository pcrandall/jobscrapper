package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gobuffalo/packr/v2"
	input "github.com/pcrandall/jobScrapper/input"
	"github.com/theckman/yacspin"
	"gopkg.in/yaml.v2"
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

type SearchConfig struct {
	Baseurl    string `yaml:"baseurl"`
	Baselimit  string `yaml:"baselimit"`
	Maxresults int    `yaml:"maxresults"`
	Jobs       []struct {
		Job      interface{} `yaml:"job"`
		Keyword  string      `yaml:"keyword"`
		Location []string    `yaml:"location"`
	} `yaml:"jobs"`
}

var (
	clear       map[string]func()
	openBrowser map[string]func()
	config      SearchConfig
	jobs        []extractedJob

	urlSlice   []string
	query      bool
	configFile bool

	baseurl    = "https://www.indeed.com/jobs?"
	baselimit  = "&limit=50"
	maxresults = 50
	pageLimit  int
)

func main() {
	CallClear()
	logfile, err := os.OpenFile("./logs/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)
	defer logfile.Close()

	flag.BoolVar(&query, "q", true, "query indeed.com, if false build site from ./sites/jobs.csv -q=false")
	flag.BoolVar(&configFile, "c", false, "use config file for query parameters ./config/config.yml -c=true")
	flag.Parse()

	// No new query, use jobs.csv to build site.
	if !query {
		buildFromJobs() // if query is false this doesn't return.
	}

	c := make(chan []extractedJob)

	if configFile {

		GetConfig() // get the config
		baseurl = config.Baseurl
		baselimit = config.Baselimit
		maxresults = config.Maxresults

		for _, job := range config.Jobs {
			keyword := strings.TrimSpace(job.Keyword)
			keyword = "q=" + strings.ReplaceAll(keyword, " ", "+")
			if len(job.Location) == 0 {
				fmt.Println("Finding", job.Keyword)
				urlSlice = append(urlSlice, config.Baseurl+keyword+config.Baselimit)
			}
			for _, location := range job.Location {
				loc := strings.TrimSpace(location)
				loc = "l=" + strings.ReplaceAll(loc, " ", "%2C")
				urlSlice = append(urlSlice, config.Baseurl+keyword+"&"+loc+config.Baselimit)
				fmt.Println("Finding", job.Keyword, "jobs in", location)
			}
		}
	} else {
		urlSlice = input.UserInput(urlSlice) // get user input for query
	}

	CallClear() // clear screen print things.

	if !configFile {
		pageLimit = maxresults / 50
	} else {
		pageLimit = config.Maxresults / 50
	}

	//TODO make channgel and go func here

	cfg := yacspin.Config{
		Frequency:       200 * time.Millisecond,
		CharSet:         yacspin.CharSets[54],
		Suffix:          "Searching Indeed...",
		SuffixAutoColon: false,
		Message:         "",
		StopCharacter:   "√",
		StopMessage:     "Completed!",
		StopColors:      []string{"fgGreen"},
		Colors:          []string{"fgYellow"},
	}

	spinner, err := yacspin.New(cfg) // handle the error
	if err != nil {
		panic(err)
	}
	spinner.Start() // Start the spinner
	for _, url := range urlSlice {
		totalPages := getPages(url)
		if pageLimit < totalPages {
			totalPages = pageLimit
		}
		for i := 0; i < totalPages; i++ {
			go getPage(url, i, c)
		}
		for i := 0; i < totalPages; i++ {
			extractedJobs := <-c
			jobs = append(jobs, extractedJobs...)
		}
	}

	spinner.Stop() // connected stop spinner

	CallClear() // clear screen print things.
	RemoveDuplicates(jobs)
	writeJobs(jobs)
	fmt.Println("Finished! Extracted", len(jobs), "jobs!")
	serveJobs()
}

func buildFromJobs() {
	lines, err := ReadCsv("./jobs/jobs.csv")
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

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Println("Request failed with status:", res.StatusCode)
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

	// TODO make this better
	doc.Find("#searchCountPages").Each(func(i int, s *goquery.Selection) {
		p := strings.Split(s.Text(), " ")
		jobCount := strings.ReplaceAll(p[23], ",", "")
		pages, _ = strconv.Atoi(jobCount)
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

	// d, err := doc.Find("#jobDescriptionText").Children().Html()
	// checkErr(err)

	d := ""

	// doc.Find("#jobDescriptionText").Children().Each(func(i int, s *goquery.Selection) {
	// 	// fmt.Printf("nodeName() = %+v\n", goquery.NodeName(s))
	// 	// fmt.Printf("s.Text() = %+v\n", s.Text())
	// 	// html, _ := goquery.OuterHtml(s)
	// 	// html = strings.ReplaceAll(html, "<br/>", "\n")
	// 	// html = strings.ReplaceAll(html, "<li>", "\u2022    ")
	// 	// html = strings.ReplaceAll(html, "<ul>", "\u2022    ")
	// 	// html = strings.ReplaceAll(html, "</li>", "")
	// 	// html = strings.ReplaceAll(html, "</ul>", "")

	// 	// if html != "" {
	// 	// 	d += html + "\n"
	// 	// }

	// 	switch goquery.NodeName(s) {
	// 	case "br":
	// 		// fmt.Printf("BR = %+v\n", goquery.NodeName(s))
	// 		d += "\n"
	// 	}
	// 	if s.Text() != "" {
	// 		switch goquery.NodeName(s) {

	// 		case "b":
	// 			d += s.Text()
	// 		case "br":
	// 			d += "\n"
	// 		case "p":
	// 			d += s.Text() + "\n\n"
	// 		case "#text":
	// 			d += "\n" + s.Text() + "\n"
	// 		case "div":
	// 			s.Children().Each(func(i int, s *goquery.Selection) {
	// 				if goquery.NodeName(s) == "li" || goquery.NodeName(s) == "ul" {
	// 				fmt.Printf(" = %+v\n", s)
	// 					str := strings.ReplaceAll(s.Text(), "\n", "\n\u2022    ")
	// 					d += str
	// 				} else if goquery.NodeName(s) == "b" {
	// 					fmt.Printf("goquery.NodeName(s) = %+v\t%+v\n", goquery.NodeName(s), s.Text())
	// 					str := strings.ReplaceAll(s.Text(), "\n", "\n\n")
	// 					d += "\n" + str
	// 				} else {
	// 					fmt.Printf("div = %+v\t%+v\n", goquery.NodeName(s), s.Text())
	// 					str := strings.ReplaceAll(s.Text(), "\n", "\n\n")
	// 					d += str
	// 				}
	// 			})

	// 		}

	// 		d = strings.ReplaceAll(d, "\n\n\n", "\n")
	// 		d = strings.ReplaceAll(d, "\n\n\n\n", "\n")
	// 		d = strings.ReplaceAll(d, "\n\n\n\n\n", "\n")
	// 		// d = strings.ReplaceAll(d, "\n\u2022    \n\u2022", "\n\u2022\n")
	// 	}
	// 	// fmt.Printf("html = %+v\n", html)
	// 	// childrenHtml = append(childrenHtml, html)
	// })

	card := doc.Find("#jobDescriptionText")
	card.Contents().Each(func(i int, s *goquery.Selection) {
		// d += s.Find("div").Text() + "\n\n"
		// d += s.Find("p").Text() + "\n\n"
		// d += "<b>" + s.Find("b").Text() + "</b>\n" // these need to be bold
		// 	d += "\u2022" + s.Find("ul").Text() + "\n"
		// 	d += "\u2022" + s.Find("li").Text() + "\n"
		// d += s.Text() + "\n\n"

		s.Contents().Each(func(i int, s *goquery.Selection) {
			switch goquery.NodeName(s) {
			case "br":
				d = "\n" + d
			}
			if s.Text() != "" {
				switch goquery.NodeName(s) {

				case "b", "br":
					d += "\n\n" + s.Text() + "\n"

				case "p", "li", "div":
					d += "\n\n" + s.Text() + "\n"

				case "#text":
					d += "\n\n" + s.Text() + "\n"

				case "ul":
					s.Each(func(i int, s *goquery.Selection) {
						if goquery.NodeName(s) != "b" {
							str := strings.ReplaceAll(s.Text(), "\n", "\n\u2022    ")
							d += "\n\n" + str + "\n"
						} else {
							// fmt.Printf("goquery.NodeName(s) = %+v\t%+v", goquery.NodeName(s), s.Text())
							d += "\n\n" + s.Text() + "\n"
						}
					})
				}
			}
		})
	})

	d = strings.ReplaceAll(d, "\n\n\n", "\n")
	d = strings.ReplaceAll(d, "\n\n\n\n", "\n")
	d = strings.ReplaceAll(d, "\n\n\n\n\n", "\n")
	d = strings.ReplaceAll(d, "\n\n\n\n\n\n", "\n")
	d = strings.ReplaceAll(d, "\n\u2022    \n\u2022", "\n\u2022\n\n")

	description <- d
}

func RemoveDuplicates(jobs []extractedJob) {
	found := make(map[string]bool)
	j := 0
	for i, x := range jobs {
		if !found[x.Id] {
			found[x.Id] = true
			(jobs)[j] = (jobs)[i]
			j++
		}
	}
	jobs = (jobs)[:j]
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("./jobs/jobs.csv")
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

	sitebox := packr.New("staticBox", "./site")

	http.HandleFunc("/", serveTemplate)

	http.Handle("/static/", // handle `/static` route
		http.StripPrefix("/static", http.FileServer(http.Dir(sitebox.Path)+"/static")),
	)

	http.Handle("/static/images/", // handle `/static/images` route
		http.StripPrefix("/static/images", http.FileServer(http.Dir(sitebox.Path)+"/static/images")),
	)

	fmt.Printf("Serving at http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)

	url := "http://localhost:" + fmt.Sprintf("%d", listener.Addr().(*net.TCPAddr).Port)

	openbrowser(url)

	log.Println(http.Serve(listener, nil))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
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
		fmt.Printf("Underlying Value: %+v\n", t)
	}
}

func GetConfig() {
	if _, err := os.Stat("./config/config.yml"); err == nil { // check if config file exists
		yamlFile, err := ioutil.ReadFile("./config/config.yml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			panic(err)
		}
	} else if os.IsNotExist(err) { // config file not included, use embedded config
		yamlFile := packr.New("configBox", "./config")

		configFile, err := yamlFile.Find("config.yml")
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(configFile, &config)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Schrodinger: file may or may not exist. See err for details.")
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	checkErr(err)
}
