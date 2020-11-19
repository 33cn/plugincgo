package com.fuzamei.chain33.guess;

import com.fuzamei.chain33.Blockchain;
import com.fuzamei.chain33.LocalDB;
import com.google.gson.Gson;

import java.util.LinkedHashMap;

public class Record {
    private static final Record INSTANCE = new Record();
    private String address;
    //(round->(guessNumber->ticketNumber))
    private LinkedHashMap<Integer, LinkedHashMap<Integer, Integer>> guessRecord;
    //(round->bonus)
    private LinkedHashMap<Integer, Long> prizeRecord;

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public LinkedHashMap<Integer, LinkedHashMap<Integer, Integer>> getGuessRecord() {
        return guessRecord;
    }

    public void setGuessRecord(LinkedHashMap<Integer, LinkedHashMap<Integer, Integer>> guessRecord) {
        this.guessRecord = guessRecord;
    }

    public LinkedHashMap<Integer, Long> getPrizeRecord() {
        return prizeRecord;
    }

    public void setPrizeRecord(LinkedHashMap<Integer, Long> prizeRecord) {
        this.prizeRecord = prizeRecord;
    }

    public static final Record getInstance(String... args) {
        if (args.length == 0) {
            return INSTANCE.loadData();
        } else {
            return INSTANCE.loadData(args[0]);
        }
    }

    public Record loadData() {
        String values = LocalDB.getFromLocalInStr(Blockchain.getFrom());
//        System.out.println("loadData record=>:"+values);
        if (values == null) {
            Record record = new Record();
            record.setAddress(Blockchain.getFrom());
            record.guessRecord = new LinkedHashMap<Integer, LinkedHashMap<Integer, Integer>>();
            record.prizeRecord = new LinkedHashMap<Integer, Long>();
            return record;
        }
        Gson gson = new Gson();
        Record record = gson.fromJson(values, new Record().getClass());
        return record;
    }

    public Record loadData(String from) {
        String values = LocalDB.getFromLocalInStr(from);
//        System.out.println("loadData record,addr: "+from+", data: "+values);
        if (values == null) {
            Record record = new Record();
            record.setAddress(from);
            record.guessRecord = new LinkedHashMap<Integer, LinkedHashMap<Integer, Integer>>();
            record.prizeRecord = new LinkedHashMap<Integer, Long>();
            return record;
        }
        Gson gson = new Gson();
        Record record = gson.fromJson(values, new Record().getClass());
        return record;
    }

    public boolean saveData() {
        Gson gson = new Gson();
        String jsonStr = gson.toJson(this);
//        System.out.println("saveData record,addr:"+this.address+"data:"+jsonStr);
        return LocalDB.setLocalInStr(this.address, jsonStr);
    }

    /**
     * Get bet information for a specific round
     *
     * @param round
     * @return
     */
    public LinkedHashMap<Integer, Integer> getGuessRecordByRound(Integer round) {
        return this.getGuessRecord().get(round);
    }

    /**
     * Query the winning amount of a specified round
     *
     * @param round
     * @return
     */
    public Long getBonusByRound(Integer round) {
        if (this.getPrizeRecord() != null) {
            return this.getPrizeRecord().getOrDefault(round,0l);
        }
        return 0l;
    }
}
