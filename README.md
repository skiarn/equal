# equal
compares two values json | xml | string

##Example
go run main.go ´[ {name: 'Sweden', code: 'SE'}, {name: 'Trump', code: 'WALL'}, {name: 'Mexico', code: 'MX'} ]´ ´[ {name: 'Sweden', code: 'SE'}, {name: 'Trump', code: 'WALL'}, {name: 'Mexico', code: 'MX'} ]´


###TODO: Handle unordered array
Goal:
Usage: ./equal ´[ {name: 'Sweden', code: 'SE'}, {name: 'Trump', code: 'WALL'}, {name: 'Mexico', code: 'MX'} ]´ ´[ {name: 'Trump', code: 'WALL'}, {name: 'Sweden', code: 'SE'}, {name: 'Mexico', code: 'MX'} ]´
