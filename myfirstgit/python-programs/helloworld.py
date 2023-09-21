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

arket_listesi = ['vanilla tesco', 'cikolata tesco', 'sut tesco', 'seker tesco', 'sogan tesco', 'limon lidl', 'bira lidl', 'su lidl', 'soda tesco', 'ekmek lidl',
                  'dondurma tesco', 'cips lidl', 'biber m&s', 'firin_kagidi lidl', 'poset lidl', 'pizza m&s', 'peynir lidl']

liste_sayisi = 0
for x in market_listesi:
    liste_sayisi=liste_sayisi +1
    esya, market = x.split()[0], x.split()[1]
    print('{0}. esyanin adi {1} ve marketi {2}'.format(liste_sayisi, esya, market))
