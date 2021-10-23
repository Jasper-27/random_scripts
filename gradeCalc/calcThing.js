firstScore = 66 // the score you got 
firstPart = 30 // percentage of first module 
targetScore = 70 // minimum to pass
    
otherPart = 100 - firstPart

overallFirstScore = firstScore * firstPart; 
overallFirstScore = overallFirstScore / 100
console.log("You have " + overallFirstScore + "% already down")

newNeeded = targetScore - overallFirstScore
console.log("Which means you need " + newNeeded + "% to go")

otherPartNeeded = newNeeded / otherPart
otherPartNeeded = otherPartNeeded * 100

console.log("Which means you need " + otherPartNeeded.toFixed(2) + "% on the other part to get " + targetScore + "% overall")



    
 


    