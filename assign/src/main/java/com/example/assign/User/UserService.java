package com.example.assign.User;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

import java.util.List;

//@Component
@Service
public class UserService {
    private UserProducer userProducer;
    @Autowired
    public UserService(UserProducer userProducer) {
        this.userProducer = userProducer;
    }
    public UserService(){

    }
    public List<User> getUser(){
        return List.of(
                new User("912919191L","hello","iak")
        );
    }
    public void addMessage(User user){
        System.out.println("came to post endpoint and to user service");
        user.printUser();
        ///  Checking the redis
        /// Pushing to kafka
        /// Using 3P vendor
        userProducer.send_message(user);

    }
}
