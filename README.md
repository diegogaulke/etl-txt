# Golang TXT ETL example

This is an example on how to read a text file and extract information and send it to an processor.<br>
Information is extracted line by line using regex expression for each type of file defined by env vars (see below).<br>
There are two implemented entities, debt and client. You can implement your own adding new extractors and processors that supports your file.<br>
The default processor method is to print extracted information.<br>
You can send extracted data to elastic search, just set the below how to configure needed env variables below.

## Build
dep ensure<br>
go build<br>

## Run...
### Set env vars FILENAME and FILE_TYPE<br>

export FILE_TYPE=debt<br>
set FILENAME=debts<br>
./etl-txt<br>