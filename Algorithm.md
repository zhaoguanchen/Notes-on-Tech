# 1.最大假期天数
#### Description
您只能在1个城市中旅行，由0到N-1的索引表示。一开始，你周一在城市0。
这些城市都是通过航班连接起来的。这些航班被表示为N*N矩阵(非必要对称)，称为代表航空公司从城市i到j城市状态的flights矩阵。如果没有从城市i到城市j的航班，flights[i][j] = 0;否则,flights[i][j]= 1。还有，flights[i][i] = 0。
你总共有K周(每周有7天)旅行。你只能每天最多乘坐一次航班，而且只能在每周一早上乘坐航班。由于飞行时间太短，我们不考虑飞行时间的影响。
对于每个城市，你只能在不同的星期里限制休假日，给定一个命名为days的N*K矩阵表示这种关系。对于days[i][j]的值，它表示你可以在j+1周的城市i里休假的最长天数，你得到的是flights矩阵和days矩阵，你需要输出你在K周期间可以获得的最长假期。

#### Analyse
dp[i]表示最后到达i城市的最长休假时间。先枚举周数，然后枚举终点，然后是起点，判断是否前往(temp[j] = Math.max(temp[j], dp[k] + days[j][i]);)，即是否进行转移。每周更新dp数组，最后从dp数组选择最大值即可。

```java
import java.util.Arrays;

/**
 * @author Guanchen Zhao
 * @Description
 * @Date 2020/8/19 14:39
 **/

/**
 * 您只能在1个城市中旅行，由0到N-1的索引表示。一开始，你周一在城市0。
 * 这些城市都是通过航班连接起来的。这些航班被表示为N*N矩阵(非必要对称)，称为代表航空公司从城市i到j城市状态的flights矩阵。如果没有从城市i到城市j的航班，flights[i][j] = 0;否则,flights[i][j]= 1。还有，flights[i][i] = 0。
 * <p>
 * 你总共有K周(每周有7天)旅行。你只能每天最多乘坐一次航班，而且只能在每周一早上乘坐航班。由于飞行时间太短，我们不考虑飞行时间的影响。
 * <p>
 * 对于每个城市，你只能在不同的星期里限制休假日，给定一个命名为days的N*K矩阵表示这种关系。对于days[i][j]的值，它表示你可以在j+1周的城市i里休假的最长天数，你得到的是flights矩阵和days矩阵，你需要输出你在K周期间可以获得的最长假期。
 */

public class Main {

    public static void main(String[] args) {
//        int[][] flights = new int[][]{{0, 0, 1}, {1, 0, 1}, {0, 1, 0}};
//        int[][] days = new int[][]{{3, 1, 6}, {1, 2, 4}, {5, 1, 2}};
        int[][] flights = new int[][]{{0, 1}, {1, 0}};
        int[][] days = new int[][]{{3, 6}, {1, 4}};
        
        int res = maxVacationDays(flights, days);
        System.out.println(res);
    }

    public static int maxVacationDays(int[][] flights, int[][] days) {
        // Write your code here
        int N = flights.length;
        int K = days[0].length;
        int[] dp = new int[N];
        Arrays.fill(dp, Integer.MIN_VALUE);

        dp[0] = 0;
        for (int i = 0; i < K; i++) {            //逐渐扩大枚举周
            int[] temp = new int[N];
            Arrays.fill(temp, Integer.MIN_VALUE);
            for (int j = 0; j < N; j++) {            //枚举终点
                for (int k = 0; k < N; k++) {            //枚举起点
                    if (j == k || flights[k][j] == 1) {            //如果城市k到城市j存在航班
                        temp[j] = Math.max(temp[j], dp[k] + days[j][i]);        //则再对当前答案进行选择，即是否从k前往j
                    }
                }
            }
            dp = temp;
        }

        int ans = 0;
        for (int i = 0; i < N; i++) {            //最后对dp数组筛选最大值即可
            ans = Math.max(ans, dp[i]);
        }
        return ans;
    }
}
```
