package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `KMPV,medium_airport,Edward F Knapp State Airport,1166,NA,US,US-VT,Barre/Montpelier,KMPV,MPV,MPV,"-72.56230164, 44.20349884"
KMPZ,small_airport,Mount Pleasant Municipal Airport,730,NA,US,US-IA,Mount Pleasant,KMPZ,MPZ,MPZ,"-91.51110077, 40.94660187"
KMQB,small_airport,Macomb Municipal Airport,707,NA,US,US-IL,Macomb,KMQB,MQB,MQB,"-90.65239715579999, 40.520099639899996"
KMQI,small_airport,Dare County Regional Airport,13,NA,US,US-NC,Manteo,KMQI,MEO,MQI,"-75.69550323, 35.91899872"
KMQJ,small_airport,Mount Comfort Airport,862,NA,US,US-IN,Indianapolis,KMQJ,,MQJ,"-85.89710236, 39.84349823"
KMQS,small_airport,Chester County G O Carlson Airport,660,NA,US,US-PA,Coatesville,KMQS,CTH,MQS,"-75.8655014, 39.97900009"
KMQT,closed,Marquette Airport,1424,NA,US,US-MI,,KMQT,,,"-87.5614013671875, 46.53390121459961"
KMQW,small_airport,Telfair Wheeler Airport,202,NA,US,US-GA,Mc Rae,KMQW,,MQW,"-82.87999725341797, 32.09579849243164"
KMQY,medium_airport,Smyrna Airport,543,NA,US,US-TN,Smyrna,KMQY,MQY,MQY,"-86.5201034546, 36.0089988708"
KMRB,medium_airport,Eastern WV Regional Airport/Shepherd Field,565,NA,US,US-WV,Martinsburg,KMRB,MRB,MRB,"-77.98459625, 39.40190125"
KMRC,small_airport,Maury County Airport,681,NA,US,US-TN,Columbia/Mount Pleasant,KMRC,MRC,MRC,"-87.17890167239999, 35.5541000366"
KMRF,small_airport,Marfa Municipal Airport,4849,NA,US,US-TX,Marfa,KMRF,MRF,MRF,"-104.017997, 30.371099"
KMRH,small_airport,Michael J. Smith Field,11,NA,US,US-NC,Beaufort,KMRH,,MRH,"-76.66059875, 34.73360062"
KMRJ,small_airport,Iowa County Airport,1171,NA,US,US-WI,Mineral Point,KMRJ,,MRJ,"-90.236198, 42.886799"
KMRN,small_airport,Foothills Regional Airport,1270,NA,US,US-NC,Morganton,KMRN,,MRN,"-81.61139678955078, 35.8202018737793"
KMRT,small_airport,Union County Airport,1021,NA,US,US-OH,Marysville,KMRT,,MRT,"-83.35160064697266, 40.224700927734375"
KMRY,medium_airport,Monterey Peninsula Airport,257,NA,US,US-CA,Monterey,KMRY,MRY,MRY,"-121.84300231933594, 36.58700180053711"
KMSA,closed,Mt Pleasant Airport,400,NA,US,US-TX,,KMSA,,,"-94.97560119628906, 33.129398345947266"
KMSL,medium_airport,Northwest Alabama Regional Airport,551,NA,US,US-AL,Muscle Shoals,KMSL,MSL,MSL,"-87.61019897, 34.74530029"
KMSN,large_airport,Dane County Regional Truax Field,887,NA,US,US-WI,Madison,KMSN,MSN,MSN,"-89.3375015258789, 43.13990020751953"
KMSO,medium_airport,Missoula International Airport,3206,NA,US,US-MT,Missoula,KMSO,MSO,MSO,"-114.0910034, 46.91630173"
KMSP,large_airport,Minneapolis-St Paul International/Wold-Chamberlain Airport,841,NA,US,US-MN,Minneapolis,KMSP,MSP,MSP,"-93.221802, 44.882"
KMSS,medium_airport,Massena International Richards Field,215,NA,US,US-NY,Massena,KMSS,MSS,MSS,"-74.84559631347656, 44.93579864501953"
KMSV,small_airport,Sullivan County International Airport,1403,NA,US,US-NY,Monticello,KMSV,,MSV,"-74.79499817, 41.70159912"
KMSY,large_airport,Louis Armstrong New Orleans International Airport,4,NA,US,US-LA,New Orleans,KMSY,MSY,MSY,"-90.25800323486328, 29.99340057373047"
KMT,closed,Kampot Airport,44,AS,KH,KH-7,Kampot,,KMT,,"104.1617, 10.6343"
KMTC,medium_airport,Selfridge Air National Guard Base Airport,580,NA,US,US-MI,Mount Clemens,KMTC,MTC,MTC,"-82.836919, 42.613463"
KMTH,medium_airport,The Florida Keys Marathon Airport,5,NA,US,US-FL,Marathon,KMTH,MTH,MTH,"-81.051399, 24.726101"
KMTJ,medium_airport,Montrose Regional Airport,5759,NA,US,US-CO,Montrose,KMTJ,MTJ,MTJ,"-107.893997192, 38.509799957300004"
KMTN,medium_airport,Martin State Airport,21,NA,US,US-MD,Baltimore,KMTN,MTN,MTN,"-76.413803, 39.325699"
KMTO,small_airport,Coles County Memorial Airport,722,NA,US,US-IL,Mattoon/Charleston,KMTO,MTO,MTO,"-88.279198, 39.477901"
KMTP,small_airport,Montauk Airport,6,NA,US,US-NY,Montauk,KMTP,,MTP,"-71.9207992553711, 41.076499938964844"
KMTV,small_airport,Blue Ridge Airport,941,NA,US,US-VA,Martinsville,KMTV,,MTV,"-80.01830291748047, 36.630699157714844"
KMTW,small_airport,Manitowoc County Airport,651,NA,US,US-WI,Manitowoc,KMTW,,MTW,"-87.68060302734375, 44.12879943847656"
KMUF,closed,Pikeville-Bledsoe Airport,961,NA,US,US-KY,,KMUF,,,"-85.19080352783203, 35.619998931884766"
KMUI,medium_airport,Muir Army Air Field (Fort Indiantown Gap) Airport,488,NA,US,US-PA,Fort Indiantown Gap(Annville),KMUI,,MUI,"-76.56939697, 40.43479919"
KMUL,small_airport,Spence Airport,292,NA,US,US-GA,Moultrie,KMUL,,MUL,"-83.7041015625, 31.137699127197266"
KMUO,large_airport,Mountain Home Air Force Base,2996,NA,US,US-ID,Mountain Home,KMUO,MUO,MUO,"-115.872002, 43.043598"
KMUT,small_airport,Muscatine Municipal Airport,547,NA,US,US-IA,Muscatine,KMUT,,MUT,"-91.14820098876953, 41.367801666259766"
KMUU,closed,Huntingdon County Airport,,NA,US,US-PA,,KMUU,,,"-77.862478, 40.329182"
KMVC,small_airport,Monroe County Aeroplex Airport,419,NA,US,US-AL,Monroeville,KMVC,MVC,MVC,"-87.350996, 31.458"`
	r := csv.NewReader(strings.NewReader(in))
	airportMap := make(map[string]string)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		airportMap[record[10]] = record[7]
		fmt.Println(record)

	}

	fmt.Println("\nWhere do you want to go today?")
	var count int = 0
	var prettyItinerary string
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	itinerary := []string{s1, s2, s3}
	fmt.Println(itinerary)
	for _, destination := range itinerary {
		fmt.Println(destination)
		if count > 0 && destination != "" {
			prettyItinerary += " to "

		}
		count++
		for code, name := range airportMap {
			if destination == code {
				prettyItinerary += name
			}
		}
	}
	fmt.Println(prettyItinerary)
}
