package com.heiwawu.example.leason1;

public enum AnimalType {
    CAT(1,"猫"),
    DOG(2, "狗");
    private int value;
    private String displayName;

    private AnimalType(int value, String displayName){
        this.value = value;
        this.displayName = displayName;
    }

    /**
     * 获取枚举的 int 值,用于数据保存以及序列化
     *
     * @return 枚举的 int 值
     */
    public int value() {
        return this.value;
    }

    /**
     * 获取枚举的显示名称
     *
     * @return 枚举的显示名称
     */
    public String displayName() {
        return this.displayName;
    }

}
