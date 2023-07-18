package templatedemo

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

type BlogPost struct {
	Title   string
	Content string
}

func SimpleTemplate() {
	post := BlogPost{"this is title", "content"}
	temp, err := template.New("blog.temp1").Parse("<h1>{{.Title}}</h1><br/><p>{{.Content}}</p>")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(os.Stdout, post)

	if err != nil {
		panic(err)
	}
}

func handleProductReport(w http.ResponseWriter, r *http.Request) {

	var randSlice = struct {
		List     []int
		BlogPost BlogPost
	}{List: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, BlogPost: BlogPost{}}
	rand.Shuffle(len(randSlice.List), func(i, j int) {
		randSlice.List[i], randSlice.List[j] = randSlice.List[j], randSlice.List[i]
	})

	switch r.Method {
	case http.MethodGet:
		//var blogPost BlogPost
		err := json.NewDecoder(r.Body).Decode(&randSlice.BlogPost)

		if err != nil {
			log.Println(err, "*")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fm := template.FuncMap{"mod": func(i, j int) bool {
			return i%j == 0
		}}
		t := template.New("output.gotmpl").Funcs(fm)

		t, err = t.ParseFiles(path.Join("templateDemo", "template", "output.gotmpl"))
		if err != nil {
			log.Println(err, "**")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var tmpl bytes.Buffer

		err = t.Execute(&tmpl, randSlice)

		if err != nil {
			log.Println("err", err)
			return
		}

		rdr := bytes.NewReader(tmpl.Bytes())

		w.Header().Set("Content-Disposition", "Attachment")

		http.ServeContent(w, r, "report.html", time.Now(), rdr)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// listeners
func StartListening() {
	demotemplate := http.HandlerFunc(handleProductReport)
	http.Handle("/temdemo", demotemplate)
	http.ListenAndServe(":5000", nil)
}
