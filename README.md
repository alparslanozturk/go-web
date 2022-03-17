# Go Web Sunucusu ( Postgresql Primary/Standby )


haproxy, replication kurulmuş postgresql sunucularından hangisinin primary veya standby olduğunu anlamsı için yazılmış bir programdır. 
Header'da http status 200 döndüğünde body kısmında primary veya standby döner.
500 döndüğünde bir sorun vardır. 


bağlantı kontrolü
```
psql "user=postgres dbname=postgres host=localhost sslmode=disable"
```


programin çalıştırilmasi 
```
./pgcheck  &
```




