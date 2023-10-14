def toplama(x,y):
    return x+y
def cikarma(x,y):
    return x-y
def carpma(x,y):
    return x*y
def bolme (x,y):
    return x/y

while True:
    print("islem sec")
    print("1. toplama")
    print("2. cikarma")
    print("3. carpma")
    print("4. bolme")

    secim = input ("select (1/2/3/4):")

    sayi1 = float(input("enter first number: "))
    sayi2 = float(input("enter second number: "))
    
    if secim == '1':
        print("result:", toplama(sayi1, sayi2))
    elif secim == '2':
        print("result:", cikarma(sayi1, sayi2))
    elif secim == '3':
        print("result:", carpma(sayi1, sayi2))
    elif secim == '4':
        print("result:", bolme(sayi1, sayi2))
print


def