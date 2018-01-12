package com.heiwawu.example.leason4.impl;

import com.heiwawu.example.leason4.Worker;
import com.heiwawu.example.leason4.WorkerProcess;

public class WhoAreYou {
    public static void main(String[] args) {
        WorkerProcess process = new WorkerProcess();
        Worker coder = new Coder();
        Worker banzhuaner = new BanZhuaner();
        process.process(coder);
        process.process(banzhuaner);
    }
}
