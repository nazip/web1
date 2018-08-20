package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Img ...
type Img struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Product ...
type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
	Img      []Img   `json:"img"`
	Quantity int     `json:"quantity"`
}

var products = []Product{
	Product{0, "Яблоко", 10.5, []Img{
		Img{"https://cdn.fishki.net/upload/post/2017/01/16/2192345/d9a7acec6ff0ab216f617e3179eb2db3.jpg", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb10edbb5093925c5e5", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb10edbb5853825c747", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb00edbb5ea3825c605", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb10edbb5d63825c653", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb20edbb57d3825c760", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb30edbb50d3925c5e3", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb20edbb56e3725ca2a", 200, 200},
		Img{"https://vsadu.ru/images/557e9cb00edbb5de3825c640", 200, 200}},
		10},
	Product{1, "Слива", 11.5, []Img{
		Img{"https://media.publika.md/ru/image/201410/full/050_93424200.jpg", 200, 200},
		Img{"http://koffkindom.ru/wp-content/uploads/2016/02/%D0%A1%D0%BB%D0%B8%D0%B2%D0%B0-%D0%B2%D0%B5%D0%BD%D0%B3%D0%B5%D1%80%D0%BA%D0%B0.jpeg", 200, 200},
		Img{"https://www.nastol.com.ua/download.php?img=201607/1920x1080/nastol.com.ua-182385.jpg", 200, 200},
		Img{"https://vsesorta.ru/upload/iblock/cbe/cbec025ce9eb2c1a22cca12a888affc5.jpg", 200, 200},
		Img{"http://pngimg.com/uploads/plum/plum_PNG8656.png", 200, 200},
		Img{"https://s1.1zoom.ru/b5050/122/Fruit_Plums_Closeup_445824_2048x1536.jpg", 200, 200},
		Img{"http://prodgid.ru/wp-content/uploads/2015/01/%D1%81%D0%BB%D0%B8%D0%B2%D0%B0.jpg", 200, 200}},
		10},
	Product{2, "Вишня", 12.5, []Img{
		Img{"https://previews.123rf.com/images/utima/utima1502/utima150200027/36245158-cherry-with-leaves-isolated-on-white-background.jpg", 200, 200},
		Img{"https://previews.123rf.com/images/photomaru/photomaru1009/photomaru100900047/7920700-sweet-cherry-isolated-on-white.jpg", 200, 200},
		Img{"https://previews.123rf.com/images/bhofack2/bhofack21602/bhofack2160200547/51972049-sweet-refreshing-cherry-cola-with-garnish-and-straw.jpg", 200, 200},
		Img{"https://cs1.livemaster.ru/storage/b5/ab/5fc656642ebd91904cff084d1ejx--materialy-dlya-tvorchestva-otdushka-vishnya-dlya-myla-svechej-.jpg", 200, 200},
		Img{"http://s1.1zoom.me/big7/504/Cherry_Closeup_Fruit_Red_422349.jpg", 200, 200},
		Img{"http://mobfanru.org/files/pics/original/eda-fon-frukty-vishnya-17149.jpg", 200, 200},
		Img{"https://previews.123rf.com/images/serg_v/serg_v1102/serg_v110200005/8796739-sweet-cherries-isolated-on-a-white-background.jpg", 200, 200}},
		10},
	Product{3, "Перцы", 13.5, []Img{
		Img{"https://cdn.fishki.net/upload/post/2017/01/16/2192345/34dcb5d4b3a0a2a331820177e08dc47c.jpg", 200, 200},
		Img{"https://get.pxhere.com/photo/plant-restaurant-food-produce-vegetable-kitchen-vegetables-peppers-bell-pepper-red-pepper-paprika-vegetable-garden-flowering-plant-green-chili-chili-pepper-yellow-pepper-land-plant-peperoncini-bell-peppers-and-chili-peppers-pimiento-italian-sweet-pepper-red-bell-pepper-habanero-chili-orange-pepper-885777.jpg", 200, 200},
		Img{"https://avatars.mds.yandex.net/get-marketpic/236284/market_0ENEvBPAMOn8lkH0N2UDiw/orig", 200, 200},
		Img{"http://pngimg.com/uploads/pepper/pepper_PNG3250.png", 200, 200},
		Img{"https://get.pxhere.com/photo/plant-fruit-food-produce-vegetable-yellow-tomato-peppers-bell-pepper-flowering-plant-chili-pepper-sweet-pepper-yellow-pepper-land-plant-potato-and-tomato-genus-bell-peppers-and-chili-peppers-pimiento-habanero-chili-958687.jpg", 200, 200},
		Img{"https://avatars.mds.yandex.net/get-marketpic/248830/market_UUdmqeE_tr1NaZ6XvF9dMg/orig", 200, 200},
		Img{"https://c.pxhere.com/photos/bb/4a/chilli_habanero_sharp_pods_vegetables_eat_red_food-715383.jpg!d", 200, 200},
		Img{"https://st.zakupka.tv/images/ffd569dc-35bc-30a8-8df8-f268523baf92.jpg", 200, 200}},
		10},
	Product{4, "Клубника", 14.5, []Img{
		Img{"https://galen-9.com/wp-content/uploads/2015/05/041.jpg", 200, 200},
		Img{"https://w-dog.ru/wallpapers/5/6/389697737201658.jpg", 200, 200},
		Img{"https://izobrazhenie.net/photo/1360-95-1/463_372087789.jpg", 200, 200},
		Img{"http://on-desktop.com/wps/Food_Berries__fruits__nuts_Ripe_strawberries_032304_.jpg", 200, 200},
		Img{"https://ru.best-wallpaper.net/wallpaper/1920x1080/1609/Red-strawberries-macro-photography-fruit-close-up_1920x1080.jpg", 200, 200},
		Img{"http://planetakartinok.net/photo/0-1/4287_anypicsru-.jpg", 200, 200},
		Img{"http://on-desktop.com/wps/Food___Berries_and_fruits_and_nuts_Ripe_strawberries_082303_.jpg", 200, 200}},
		10},
	Product{5, "Морковь", 15.5, []Img{
		Img{"http://getfaster.ru/upload/iblock/1bf/1bfa3b8fc993e332a6bac5363f22f9bc.jpg", 200, 200},
		Img{"https://orchardo.ru/wp-content/uploads/2012/12/kogda-luchshe-ubirat-morkov-i-kak-pravilno-provodit-uborku-urozhaya-2.jpg", 200, 200},
		Img{"https://avatars.mds.yandex.net/get-marketpic/230810/market_4BjaF8MJF71v4DSbj05oqQ/orig", 200, 200},
		Img{"http://mtdata.ru/u16/photo8EEA/20023600637-0/original.jpg", 200, 200},
		Img{"http://businesspskov.ru/pictures/160120100424.jpg", 200, 200},
		Img{"http://pishhevarenie.com/wp-content/uploads/2018/05/12.jpg", 200, 200},
		Img{"https://get.pxhere.com/photo/food-produce-vegetable-carrot-140344.jpg", 200, 200}},
		10},
	Product{6, "Помидор", 16.5, []Img{
		Img{"http://market-fmcg.ru/media/a/originals/sejw8_photo.jpg", 200, 200},
		Img{"http://s1.1zoom.me/b5050/471/377881-svetik_1280x800.jpg", 200, 200},
		Img{"http://mmedia.ozone.ru/multimedia/1021379124.JPG", 200, 200},
		Img{"https://get.pxhere.com/photo/plant-fruit-flower-food-red-produce-vegetable-healthy-tomato-vegetables-frisch-macro-photography-flowering-plant-nachtschattengew-chs-land-plant-potato-and-tomato-genus-535007.jpg", 200, 200},
		Img{"https://trickybox.ru/public/ttransd4800.jpg", 200, 200},
		Img{"http://img.pix.uz/u7786f145286.jpg", 200, 200},
		Img{"http://www.coollady.ru/pic/0002/055/070.jpg", 200, 200}},
		10},
}

func main() {
	router := httprouter.New()
	router.GET("/products", ShowProducts)
	router.GET("/product/:id", ShowProduct)
	router.GET("/error", ShowError)
	http.ListenAndServe(":4321", router)
}

// ShowError (w http.ResponseWriter, r *http.Request)
func ShowError(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	js, err := json.Marshal("NOT FOUND")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ShowProduct (w http.ResponseWriter, r *http.Request)
func ShowProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	index, error := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if error != nil || index > 6 {
		http.Redirect(w, r, "", http.StatusNotFound)
		return
	}

	product := products[index]

	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// SetAccessControl (w http.ResponseWriter, r *http.Request)
func SetAccessControl(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

// ShowProducts (w http.ResponseWriter, r *http.Request)
func ShowProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	SetAccessControl(w, r)
	// time.Sleep(10 * time.Second)
	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
