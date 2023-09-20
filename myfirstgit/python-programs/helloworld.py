url = https://www.sinefil.com/otodidakt/all/watchlist

headers = {Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36
}


requests.get (url, headers=headers)


def calculator():
    num1 = float(input("3"))
    num2 = float(input("5"))

    print("işlemler")
    print("toplama")
    print("çıkarma")
    
    secim = input("5-3")
    if secim == '1':
        sonuc = num1 + num2