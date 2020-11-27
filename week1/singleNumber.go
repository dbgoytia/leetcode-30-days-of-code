func singleNumber(nums []int) int {
    
    seen := map[int]bool{}    // A map in which to remember seen elements
    
    for i := 0; i < len(nums); i++ {    
        if seen[nums[i]] == false {
            seen [nums[i]] = true    // First time you see the number, append it to the map
        } else {
            delete(seen, nums[i])    // Second time you see the number, remove it from the map
        }
    }
    
    // Return the only value in the map
    var result int
    for key, _ := range seen {
        result = key
        break
    }
    return result
}
