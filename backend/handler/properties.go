package handler
import(	
	"net/http"
	"crim/models"
	"encoding/json"
	"io"
	"net/url"
)
func GetProperty(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		var prop models.Property
		err := json.NewDecoder(body).Decode(&prop)
		if err != nil{
			http.Error(w, http.StatusText(400), 400)
			return
		}
		where := prop.CreateUrl()
		if where == ""{
			http.Error(w, http.StatusText(400), 400)
			return
		}
		url := "https://catastro.crimpr.net/proxy/proxy.ashx?https://catastro.crimpr.net/server/rest/services/Parcelario/Parcelas/MapServer/654/query?f=json&where=" + url.PathEscape(where) + "&returnGeometry=true&spatialRel=esriSpatialRelIntersects&outFields=OLDPID,NUM_CATASTRO,TIPO,CATASTRO,MUNICIPIO,CONTACT,DIRECCION_FISICA,DIRECCION_POSTAL,CABIDA,LAND,STRUCTURE,MACHINERY,TOTALVAL,EXEMP,EXON,TAXABLE,DEEDBOOK,DEEDPAGE,ESTATE,DEEDNUM,SALESAMT,SALESDTTM,SELLERNAME,BYERNAME,Shape.STArea(),Shape.STLength(),OBJECTID_1,INSIDE_X,INSIDE_Y,OBJECTID&outSR=102100&resultOffset=0&resultRecordCount=2000"
		res, err := http.Get(url)
		if err != nil{
			http.Error(w, http.StatusText(500), 500)
			return
		}
		response, err := io.ReadAll(res.Body)
		if err != nil{
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(response)
}
