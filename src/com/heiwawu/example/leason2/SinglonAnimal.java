package com.heiwawu.example.leason2;

public class SinglonAnimal {
    private static SinglonAnimal ourInstance = new SinglonAnimal();

    public static SinglonAnimal getInstance() {
        return ourInstance;
    }

    private SinglonAnimal() {
        System.out.println("animal is objected");
    }
}
