package com.heiwawu.example.leason2;

public class LearnSinglon {

    public static void main(String[] args) {
        SinglonAnimal singlonAnimal1 = SinglonAnimal.getInstance();
        SinglonAnimal singlonAnimal2 = SinglonAnimal.getInstance();
    }
}
