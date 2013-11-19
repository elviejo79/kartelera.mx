package main

import(
	"io/ioutil"
	"fmt"
	"github.com/moovweb/gokogiri"
//	"github.com/moovweb/gokogiri/html"
	"github.com/moovweb/gokogiri/xml"
	"github.com/gregjones/httpcache"
	"net/http"
	"strconv"
	"encoding/json"
)

func main(){
	theaters := extractTheaters("http://cinemex.com.mx/")
	for _,t := range theaters {
		//fmt.Println("http://cinemex.com.mx/cines/"+t.id)
		if t.id == "4" || t.id=="54" || t.id=="199" {
			movies, _ := extractMovies("http://cinemex.com.mx/cines/"+t.id)
			for _,m := range movies {
				fmt.Printf("cinemex,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
					t.city, t.col,t.id,t.name,m[0],m[1],m[2],m[3],m[4],m[5])
			}
		}
	}
}

func extractTheaters(url string) (result []cine) {

	html,_ := getBody(url)
	doc, err := gokogiri.ParseHtml([]byte(html))
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

	for c,edo := range cities {
		if edo == "Zacatecas" || edo == "Aguascalientes" { //esto es para limitarlo sólo a 3 cines
			url := fmt.Sprintf("http://cinemex.com.mx/getddCines.php?ciudad=%d&movieId=",c)
			json_cines,_ := getBody(url)
			theaters, _ := theaters_json_decoder(json_cines,cities[c])
			result = append(result,theaters...)
		}

	}
	return result
}


func extractMovies(url string) (res [][]string, err error) {
	html,_ := getBody(url)
	doc, _ := gokogiri.ParseHtml([]byte(html))
	defer doc.Free()
	movies,_ := doc.Search("id('sch-cont')/div");
	for _, m := range movies{ 
		//fmt.Printf("--- %#v",m)
		row := []string{
			nodeContent("div[@class='cinema']",m),
			nodeContent("div[@style='width:35px;']",m),
			nodeContent("div[3]",m),
			nodeContent("div/img/@src",m),
			nodeContent("div[6]/div/p[1]|div[6]/p[1]",m),
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
	client := http.Client{Transport: httpcache.NewMemoryCacheTransport()}

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
