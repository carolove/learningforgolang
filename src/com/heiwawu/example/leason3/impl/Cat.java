package com.heiwawu.example.leason3.impl;

import com.heiwawu.example.leason3.Animal;

public class Cat implements Animal{
    private String name;

    public Cat(String name) {
        this.name = name;
    }

    @Override
    public void eat() {
        System.out.println("fish");
    }

    @Override
    public String name() {
        return this.name;
    }
}
