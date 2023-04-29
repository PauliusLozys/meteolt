# weather (Meteo-CLI)
Atvaizduojama valandinė dienos temperatūra pasinaudojant Meteo.lt API

# Kodo kompiliavimas
Reikia bent GO 1.20 versijos.
```sh
go build .
```

# Komandos
 - -h komandų sąrašas
 - -r <Numeris1|. Numeris2|.> intervalo parinkimas [Numeris1, Numeris2]. Galima naudoti "." norint imti visus įeinančius skaičius.
 - -lv rodyti turinį sąrašo formatu (Numatytas yra stulpelio formatu)
 - -lvi rodyti turinį sąrašo formatu su papildoma informacija
 - -c <Miesto pavadinimas> pakeisit rodomą miestą
 - -d <Numeris> nurodyti kurią dieną rodyti (Numatytas 0). 
 - -n rodyti kitos dienos prognozę

# Stulpelio formatas
![Stulpelio formatas](https://github.com/RainbowDog98/Meteo-CLI/blob/master/cmd/weather/weather%20cloumn.png)
# Sąrašo formatas
![Stulpelio formatas](https://github.com/RainbowDog98/Meteo-CLI/blob/master/cmd/weather/weather%20list.png)

