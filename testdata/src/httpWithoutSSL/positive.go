package httpWithoutSSL

const _ string = "http://delivery-club.ru" // want `HTTP links are not secure`
const _ = "http://delivery-club.ru"        //TODO: false-negative, type - untyped string

func warn() {
    _ = "http://delivery-club.ru" // want `HTTP links are not secure`

    print(123, "http://delivery-club.ru", 123123) // TODO: false-negative, i guess that now we break on first arg and doesn't continue
    print("http://delivery-club.ru", 123123) // want `HTTP links are not secure`
}
