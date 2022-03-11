
def test(a):
    for i in a:
        if i == 1:
            continue
        else:
            print(i)




if __name__ == '__main__':
    print("The is a Test Python File.")
    a = [1, 2, 3, 4, 5]
    test(a)