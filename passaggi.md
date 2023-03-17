# Rename-File
Questo piccolo programma mi serve per rinominare i files mp3 per poi metterli 
su una sd card e lanciarli usando il player Yx5300 e ESP8266.

## Formato dei files
Il formato del file è mp3 con un nome a 6 caratteri, dove i primi 3 sono un numero
compreso tra 001 e 999.

## Creazione della Lista
Il programma https://github.com/aaaasmile/MP3-YX5300 esegue tutti i files messi all'interno
di un folder, per esempio il folder 01, in una lista casuale senza ripetizioni.
Il folder che viene processato è codificato hard mel file main.go. Un solo folder
viene processato per volta.
Quindi è importante avere i files mp3 con il nome giusto per esere lanciati.
Per creare una lista copio i files mp3 presi da itunes manualmente nella directory ./folder_01
Qui i nomi dei files sono i più disparati. 
Ora lancio questo programma Rename-file con 

    go run main.go 

e la lista dei files dovrebbe essere
pronta. 
Ora bisogna copiare il contenuto di folder_01 nella sd-card del player e modificare il codice di 
MP3-YX5300. 
Come mai bisogna modificare il codice? Perché il comando seriale AT del player, che teoriacamnte
dovrebbe fornire il numero dei files nel folder 01 della SD-card, non sono riuscito a farlo andare.
Il valore che legge è sempre zero.
Siccome la lista dei files mp3 non è che viene attualizzata spesso, quando succede, basta ricompilare
il software con il valore corretto.
L'ultima volta che ho letto la carta sd sul computer non ci sono riuscito con il san disk usb reader.
Ho usato, invece, l'adattatore della carta nello slot del portatile.

Nota che non tengo gli originali mp3 prima di rinominarli. Perlomeno non in questo progetto.
Semplicemente con file explorer scrivo il titolo nei meta dati, così riesco a vedere i titoli che ho incluso.

Il file da cambiare nel progetto ESP è synch-card-data.inc

## Splittare files mp3 grandi
Ho scaricato dei files da youtube e sono troppo grandi per sentirli in macchina 
dove una pausa mi rimanda il play all'inizio del file. Così ho fatto uno split.
Con WLC 20.04 ho lanciato il seguente comando per spezzettare il file dei sintetizzatori in
blocchi da 8 minuti e rotti:

    igors@Laptop-Toni:/mnt/d/scratch/go-lang/rename-file/tmp$ ffmpeg -i 001xBa.mp3 -f segment -segment_time 500 -c copy synt-mix-%03d.mp3