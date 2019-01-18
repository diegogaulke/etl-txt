# Golang TXT ETL example

This is an example on how to extract information from a text file, transform and load it into some datasource.<br>
Information is extracted line by line using regex expression for each type of file defined by env vars (see below).<br>
There are two implemented entities, debt and client. You can implement your own adding new extractors, transformers and loaders  that supports your file.<br>
The default loader method is to print extracted and transformed information.<br>

## Build
dep ensure<br>
go build<br>

## Run...
### Set env vars FILENAME and FILE_TYPE<br>

export FILE_TYPE=debt<br>
export FILENAME=debts<br>
./etl-txt<br>