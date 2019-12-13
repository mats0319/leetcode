package mario

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
        return 0
    }

    obstacleGrid[0][0] = 1-obstacleGrid[0][0]
    // handle 0th row separately
    for i := 1; i < len(obstacleGrid[0]); i++ {
        if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1 {
            obstacleGrid[0][i] = 1
        } else {
            obstacleGrid[0][i] = 0
        }
    }

    for i := 1; i < len(obstacleGrid); i++ {
        // handle 0th element in each row separately
        if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
            obstacleGrid[i][0] = 1
        } else {
            obstacleGrid[i][0] = 0
        }

        for j := 1; j < len(obstacleGrid[0]); j++ {
            if obstacleGrid[i][j] == 1{
                obstacleGrid[i][j] = 0
            } else {
                obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
            }
        }
    }

    return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
