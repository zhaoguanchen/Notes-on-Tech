# 1. twoSum
my solution:
```java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
            for (int i = 0; i < nums.length; i++) {
                map.put(nums[i], i);
            }
            for (int i = 0; i < nums.length; i++) {
                int b = target - nums[i];
                if (map.containsKey(b) && i != map.get(b)) {
                    int index = map.get(b);
                    return new int[]{i, index};
                }
            }
            return new int[]{0, 1};
    }
}
```

better solution:
```java
public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (map.containsKey(complement)) {
            return new int[] { map.get(complement), i };
        }
        map.put(nums[i], i);
    }
    throw new IllegalArgumentException("No two sum solution");
}
```


# 2. removeDuplicates

```java
public class Solution {
    /*
     * @param nums: An ineger array
     * @return: An integer
     */
    public int removeDuplicates(int[] nums) {
        // write your code here
        int count = 0;
        int base = 0;
        for (int  i = base +1;i< nums.length ;i++ ){
            if (nums[i] == nums[base]){
                count++;
                continue;
            }             
            nums[++base] = nums[i];            
        }         
        return nums.length - count;
    }
}
```
