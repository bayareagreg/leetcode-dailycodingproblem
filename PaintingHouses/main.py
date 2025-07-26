def build_house(matrix):
    n = len(matrix)
    k = len(matrix[0])
    solution_matrix = [[0] * k]

    for r, row in enumerate(matrix):
        row_cost = []
        print(" r = ", r, ", row = ", row)
        for c, val in enumerate(row):
            row_cost.append(min(solution_matrix[r][i]
                                for i in range (k)
                                if i != c) + val)
        print(row_cost)
        
        solution_matrix.append(row_cost) 
        
    return min(solution_matrix[-1])        



matrix = [[1, 2, 10],
		  [6, 3, 9],
		  [4, 5, 6]]

build_house(matrix)

