# go ile db bağlantısı

Bu proje go dilinde veritabanı bağlantısı kurmak ve bağlanılan veritabanında işlemler yapmak için yardımcı döküman amacıyla yazılmıştır.
dosya ve değişken isimleri, modeller, sorgular değiştirilip gerekli özelleştirmeler yapılarak her yerde kullanılabilir. :+1:

### içerik

- :gear: [db bağlantısı](https://github.com/murattarslan/go_db_connect#go-ile-veritaban%C4%B1na-ba%C4%9Flanma)
- :triangular_ruler: [tablo oluşturma](https://github.com/murattarslan/go_db_connect#yeni-bir-tablo-olu%C5%9Fturma)
- :heavy_plus_sign: [tabloya veri ekleme](https://github.com/murattarslan/go_db_connect#tabloya-veri-ekleme)
- :mag: [tablodan veri alma](https://github.com/murattarslan/go_db_connect#tablodan-veri-alma)

### çok yakında...

- :heavy_minus_sign: tablodan veri silme
- :wrench: tablodaki veriyi güncelleme

## go ile veritabanına bağlanma

dao paketinin içerisinde `conf.go` dosyasında db için konfigürasyon bilgileri bulunacak

```
package dao

const (
	host     = "localhost"
	port     = "5432"
	user     = "user"
	password = "password"
	dbname   = "dbname"
)
```

konfigürasyon bilgileri tanımlandıktan sonra geriye sadece bağlanmak kalıyor

```
	conf :=
		"host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable"

	db, err := sql.Open("postgres", conf)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println("connect")
	
	// başarılı şekilde bağlandığımızda fonksiyon bitiminde bağlantıyı kesmemiz gerekecek
	defer db.Close()

```
## yeni bir tablo oluşturma

Öncelikle bir model oluşturmamız gerekiyor bu örnekte modelimiz ```type Desk struct``` olacak

```
type Desk struct {
	name   string
	id     int64
	active bool
}
```

Model oluştuğuna göre bu modele uygun bir sql sorgusu hazırlayabiliriz

```
	createQuery :=
		"CREATE TABLE IF NOT EXISTS " + tableName + " (" +
			" ID SERIAL PRIMARY KEY," +
			" NAME TEXT NOT NULL," +
			" ACTIVE BOOL NOT NULL" +
			" ); "
```
Burada standart bir sql sorgusu oluşturuyoruz. Sorguda id değerini 'Serial' tanımladık, bu yüzden veri eklerken id değeri vermemeliyiz. Bu tabloya eklenen her verinin id değeri otomatik olarak postgreSQL tarafından veriliyor.

Herşey hazır şimdi son olarak sorguyu çalıştıracağız
```
r, err := db.Exec(createQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
	
```

ve bitti. Eğer konsolda 'complete' yazısı görüldüyse tablomuz hazır demektir. :tada:

Şimdi sıra tabloya öge eklemekte...

## Tabloya veri ekleme

Oluşturulan tabloya öge eklerken dikkat edeceğimiz durum sorgu. Bu işlemde yazdığımız sorgu bize eklediği verinin id değerini dönecek ve tablo oluştururken yaptığımız konfigürasyon sebebiyle bu sorguda id değeri vermiyoruz. [^1]

```
insertQuery := fmt.Sprintf("insert into %s (name, active) values ('%s', %v) returning id;", tableName, item.name, item.active)

```

Sorgunun sonundaki 'returning id' eki bize eklenen verinin aldığı id değerini dönecek.

Sorgu hazır olduğuna göre şimdi çalıştırma zamanı...

```
	id := 0
	err = db.QueryRow(insertQuery).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("add item... id:%d", id)
```

Konsolda id değerini gördüyseniz tebrikler. :tada:

Sıradaki madde bu eklediğimiz verileri tekrar çekme üzerine.

## Tablodaki verileri alma

Sorgu standart select sorgusu. komut yazılır, istenen parametreler girilir, sonra tablo gösterilir

```
selectQuery := fmt.Sprintf("select id, name, active from %s", tableName)

```

Sorgu hazır olduğuna göre şimdi çalıştırma zamanı...

```

	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
	}
	// satırlarla işimiz bittiğinde kapatılması için
	defer rows.Close()

	var result []Desk

	// satırları tek tek dolaşıp her satırda nesnemizi oluşturarak dönüş listemize ekliyoruz
	for rows.Next() {

		var id int64
		var name string
		var active string

		err = rows.Scan(&id, &name, &active)
		if err != nil {
			panic(err)
		}

		item := Desk{name, id, active}
		result = append(result, item)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return result
```

Ve son :tada:

fonksiyonumuz tüm saatırları tek tek dolaşıp nnesneleri oluşturdu. Ve bize tabloyu liste olarak döndü.

Verileri yazıp yazdıklarımızı gördüğümüze göre şimdi hoşumuza gitmeyenleri silme zamanı


[^1]: :warning: sorguda string değer verirken tırnak işareti(') kullanmayı unutmayın
