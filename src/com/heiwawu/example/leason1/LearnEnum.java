package com.heiwawu.example.leason1;

public class LearnEnum {

    public static void main(String[] args){

        AnimalType animalType = AnimalType.CAT;
        if (animalType == AnimalType.CAT) {
            System.out.println(animalType.displayName());
            animalType = AnimalType.DOG;
        }
    }
}
