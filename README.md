# go db connect

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
Burada standart bir sql sorgusu oluşturuyoruz. Herşey hazır şimdi son olarak sorguyu çalıştıracağız
```
r, err := db.Exec(createQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
	
```

ve bitti. Eğer konsolda 'complete' yazısı görüldüyse tablomuz hazır demektir.

Şimdi sıra tabloya öge eklemekte...
