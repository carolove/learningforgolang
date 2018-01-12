package com.heiwawu.example.leason3.impl;

import com.heiwawu.example.leason3.Animal;

public class Dog implements Animal{
    private String name;
    public Dog(String name){
        this.name = name;
    }
    @Override
    public void eat() {
        System.out.println("bone");
    }

    @Override
    public String name() {
        return this.name;
    }
}
