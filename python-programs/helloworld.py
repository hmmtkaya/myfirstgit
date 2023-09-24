market_listesi = ['vanilla tesco', 'cikolata tesco', 'sut tesco', 'seker tesco', 'sogan tesco', 'limon lidl', 'bira lidl', 'su lidl', 'soda tesco', 'ekmek lidl',
                  'dondurma tesco', 'cips lidl', 'biber m&s', 'firin_kagidi lidl', 'poset lidl', 'pizza m&s', 'peynir lidl']

liste_sayisi = 0
for x in market_listesi:
    liste_sayisi=liste_sayisi +1
    esya, market = x.split()[0], x.split()[1]
    print('{0}. esyanin adi {1} ve marketi {2}'.format(liste_sayisi, esya, market))
          
    
