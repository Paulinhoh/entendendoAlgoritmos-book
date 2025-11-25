def busca_binaria(arr, item):
    inicio = 0
    fim = len(arr) - 1

    while inicio <= fim:
        meio = (inicio + fim) // 2
        chute = arr[meio]
        if chute == item:
            return meio
        elif chute > item:
            fim = meio - 1
        else:
            inicio = meio + 1
    return None


def main():
    arr = [1, 3, 5, 7, 9]
    print(busca_binaria(arr, 3))
    print(busca_binaria(arr, -1))


if __name__ == "__main__":
    main()
