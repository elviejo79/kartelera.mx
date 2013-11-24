
package cinemex

import(
	"io/ioutil"
	"strings"
	"time"
	"fmt"
	"github.com/moovweb/gokogiri"
//	"github.com/moovweb/gokogiri/html"
	"github.com/moovweb/gokogiri/xml"
//	"github.com/gregjones/httpcache"
	"net/http"
	"strconv"
	"encoding/json"
)


func Screenings() (result []map[string]string){
	theaters := extractTheaters("http://cinemex.com.mx/")
	for _,t := range theaters {
		movies, _ := extractMovies("http://cinemex.com.mx/cines/"+t.id)
		for _,m := range movies {
			result = append(result, map[string]string{
				"cine":"cinemex" , 
				"edo": t.city, 
				"col":t.col , 
				"cineId": t.id, 
				"cineName":t.name, 
				"title": m[0], 
				"img": m[1], 
				"rating": m[2] , 
				"language": m[3], 
				"roomType": m[4], 
				"date": m[5], 
				"time":m[6],})
		}

	}
	return
}

func extractTheaters(url string) (result []cine) {

	html,_ := getBody(url)
	doc, err := gokogiri.ParseHtml(html)
	if err !=nil {
		fmt.Println(err)
	}
	defer doc.Free()

	options,_ := doc.Search("//option");
	cities := make(map[int]string)
	for _, o := range options{
		val,_ := strconv.Atoi(o.Attributes()["value"].Content())
		if val > 0 {
			cities[val] = o.Content()
		}
	}

	for c,_ := range cities {
		if c == 1 || c == 62 || c == 73 {
			url := fmt.Sprintf("http://cinemex.com.mx/getddCines.php?ciudad=%d&movieId=",c)
			json_cines,_ := getBody(url)
			theaters, _ := theaters_json_decoder(json_cines,cities[c])
			result = append(result,theaters...)
		}
	}
	return result
}


func extractMovies(url string) (res [][]string, err error) {
	html,err := getBody(url)
	if err != nil {
		fmt.Printf("%#v",err)
	}
	doc, _ := gokogiri.ParseHtml(html)
	defer doc.Free()
	movies,_ := doc.Search("id('sch-cont')/div");
	for _, m := range movies{ 
		t := time.Now().Format("20060102")
		title := strings.Replace(strings.ToUpper(nodeContent("div[@class='cinema']",m)),":","",-1) //title
		roomType := nodeContent("div/img/@src",m) //roomType
		roomType = strings.Replace(roomType,"/visual/imgs/icon-sch-cinemex.gif","",-1)
		roomType = strings.Replace(roomType,"/visual/imgs/icon-sch-platino.gif","PLA",-1)
		roomType = strings.Replace(roomType,"/visual/imgs/btn_premium.jpg","PREM",-1)
		roomType = "R"+roomType

		row := []string{
			title,
			nodeContent("a/img/@src",m), //img
			nodeContent("div[@style='width:35px;']",m), //rating
			nodeContent("div[3]",m), //language
			roomType,
			t, //nodeContent("div[6]/div/p[1]|div[6]/p[1]",m), //date
		}

		hours,_ := m.Search("div[6]/div/a")
		horas := []string{}
		for _,e := range hours {
			horas = append(horas,e.Content())
		}

		if row != nil {
			for _,h := range horas {
				res = append(res,append(row,h))
			}
		}

	}

	return
}

func getBody(url string) (body []byte, err error) {
	client := http.Client{}
	resp, err := client.Get(url)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()
        body, err = ioutil.ReadAll(resp.Body)
        if err != nil {
                return nil, err
        }
	return body,nil
}

func nodeContent(x_path string,m xml.Node) (result string){
	ts,_ := m.Search(x_path)
	for _,e := range ts{
		result = e.Content()
	}
	return 
}


type cine struct {
	city string
	col string
	id string
	name string
}

func theaters_json_decoder (json_data []byte, city string) (result []cine,err error){
	var m []map[string]map[string]map[string]map[string]string
	_ = json.Unmarshal(json_data, &m)
	
	for _,strCity := range m {
		for colName,strCines := range strCity {
			for _,strNumName := range strCines["cines"] {
				for cineId, cineName := range strNumName {
					result = append(result, cine{city,colName,cineId,cineName})
				}
			}
		}
	}

	return result,nil
}
