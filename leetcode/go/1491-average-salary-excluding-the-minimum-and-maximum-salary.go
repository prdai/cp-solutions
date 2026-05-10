package main;

import ("slices")

func average(salary []int) float64 {
       minSalary := slices.Min(salary);
       maxSalary := slices.Max(salary);
       var tot, count float64;
       for _, value := range salary {
        if (value != minSalary && value != maxSalary){
            tot += float64(value);
            count++; 
        }
       }
       return tot/count;
}

