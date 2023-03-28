# quizgame


## Run Timed Quizzes via the Command Line  
A program that will read in a quiz provided via a CSV file and then give the quiz to a user    
keeping track of how many questions they get right and how many they get incorrect.   

The CSV file will be in a format like below, where the first column is a question   
and the second column in the same row is the answer to that question.
```
whats 5 + 5,10
whats the fastest animal on land,cheetah
whats the ASCII code for 'space' char,32
```

##  Usage
 - ## flags  
   -file => source csv file for questions, defaults to "problems.csv"  
   -timeout => time limit for a quiz game, default is 30s  
   -shuffle => shuffle questions on each run, default is false  
 - ## Install
   To install, run  
   ```git clone https://github.com/namikaze-dev/quizgame```  
   ```cd quizgame```  
   ```go build .```  
   finally  
   ```./quizgame```
