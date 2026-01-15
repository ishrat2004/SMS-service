package com.example.assign.User;

public class User {

    private String phone_number;
    private String message;
    private String name;

    public User(String phone_number, String message, String name){
        this.phone_number = phone_number;
        this.message = message;
        this.name = name;
        System.out.println(message);
    }
    public String getPhone_number() {
        return phone_number;
    }
    public void setPhone_number(String phone_number) {
        this.phone_number = phone_number;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String getName() {
        return name;
    }

    public String getMessage() {
        return message;
    }
    public void printUser(){
        System.out.println(name);
        System.out.println(message);
        System.out.println(phone_number);
    }
}
