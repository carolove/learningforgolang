package com.heiwawu.example.leason3;

import com.heiwawu.example.leason3.impl.Cat;
import com.heiwawu.example.leason3.impl.Dog;

public class WhoIAm {
    public static void  main(String[] args) {
        Animal dog = new Dog("beego");
        dog.eat();
        System.out.println(dog.name());

        Animal cat = new Cat("lily");
        cat.eat();
        System.out.println(cat.name());
    }
}
