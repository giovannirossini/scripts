def matrix(x,y,z,n):
    list = [[i, j, k] for i in range(x + 1) for j in range(y + 1) for k in range(z + 1) if i + j + k != n]
    return print(list)

if __name__ == '__main__':
    matrix(1,2,1,3)