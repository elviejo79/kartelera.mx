package cinepolis

import(
	"io/ioutil"
	"time"
	"fmt"
	"strings"
	"github.com/moovweb/gokogiri"
//	"github.com/moovweb/gokogiri/html"
	"github.com/moovweb/gokogiri/xml"
//	"github.com/gregjones/httpcache"
	"net/http"
	"strconv"
	"encoding/json"
)


func Screenings() (result []map[string]string){
	cities := extractTheaters("http://cinepolis.com.mx/")
	for ic,c := range cities {
		if strings.Contains(c,"D.F.") {
			c = "D.F. y Área Metropolitana"
		}
		if strings.Contains(c,"Monterrey") {
			c = "Mty. y Área Metropolitana"
		}


		if ic == 3 || ic==124 || ic==32 {
			movies, _ := extractMovies(fmt.Sprintf("http://cinepolis.com/_CARTELERA/cartelera.aspx?ic=%d", ic))
			for _,m := range movies {
				result = append(result, map[string]string{
					"cine":"cinepolis" , "edo": c, "col":c , 
					"cineId": m[0], 
					"cineName":m[1], 
					"title": m[2], 
					"rating": m[3] , 
					"language": m[4], 
					"roomType": m[5], 
					"date": m[6], 
					"time":m[7]})
			}

		}
	}

	return
}

func extractTheaters(url string) map[int]string {

	html,_ := getBody(url)
	doc, err := gokogiri.ParseHtml(html)
	if err !=nil {
		fmt.Println(err)
	}
	defer doc.Free()

	options,_ := doc.Search("id('ctl00_ddlCiudad')/option");
	cities := make(map[int]string)
	for _, o := range options{
		val,_ := strconv.Atoi(o.Attributes()["value"].Content())
		if val > 0 {
			cities[val] = o.Content()
		}
	}
	
	return cities

}


func extractMovies(url string) (res [][]string, err error) {
	html,err := getBody(url)
	if err != nil {
		fmt.Printf("%#v",err)
	}
	doc, _ := gokogiri.ParseHtml(html)
	defer doc.Free()
	//theaters,_ := doc.Search("//a[ends-with(@id,'306')]");
	movies,_ := doc.Search("//a[contains(@id, 'idPelCine')]")
	
	for _, m := range movies{ 
		cineId := nodeContent("@id",m)[14:]
		titulo := nodeContent("parent::*//a[@class='peliculaCartelera']",m)
		subtitulos := titulo[len(titulo)-3:]
		if subtitulos == "Sub" {
			subtitulos = "SUBTITULADA"
                } else {
			subtitulos = "ESPAÑOL"
		}

		sala := titulo[len(titulo)-7:len(titulo)-4]
		if strings.Contains(titulo," 4D") {
			titulo=titulo[:strings.Index(titulo," 4D")]
			sala = "4D"
		} else if strings.Contains(titulo," 3D ") {
			titulo=titulo[:strings.Index(titulo," 3D ")]
			sala = "3D"
		} else if strings.Contains(titulo," IMAX") {
			titulo=titulo[:strings.Index(titulo," IMAX")]
			sala = "IMAX"
		} else if strings.Contains(titulo," XE") {
			titulo=titulo[:strings.Index(titulo," XE ")]
			sala = "XE"
		} else {
			titulo=titulo[:strings.Index(titulo," Dig ")]
			sala = "Dig"
		}
		titulo = strings.ToUpper(titulo)
		t := time.Now().Format("20060102")
		
		row := []string{
			cineId, //cineID
			nodeContent("//select[@name='cartelera"+cineId+"']/parent::*/parent::*/parent::*//span[@class='TitulosBlanco']",m),
			titulo,
			nodeContent("parent::*//span[@class='textoPequeñoNegro']",m),
			subtitulos,
			sala,
			t,
			/*nodeContent("div/img/@src",m),
			nodeContent("div[6]/div/p[1]|div[6]/p[1]",m),
*/		
		}

		hours,err := m.Search("parent::*/parent::*//*[contains(@class,'horariosCartelera')]")
		if err != nil {
			fmt.Println(err)
		}

		horas := []string{}
		for _,e := range hours {
			t,_:= time.Parse("3:04pm",e.Content())
			horas = append(horas,t.Format("15:04"))
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

