# Rename File
This is a very simple rename file program written in go.
I use it to change mp3 file names that could be stored in a Sd Card for the MP3-YX5300 project (https://github.com/aaaasmile/MP3-YX5300) 

## How to use it
Change in the file main.go the dirToScan to point to your mp3 file folder that you want to rename.
Then run it with:

    go run main.go

## Example result
In folder:

    'my best song.mp3'
    'your best song.mp3'

will be:

    001myb.mp3
    002you.mp3
