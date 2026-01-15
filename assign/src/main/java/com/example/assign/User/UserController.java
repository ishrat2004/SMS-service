package com.example.assign.User;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(path="api/v1/send/sms")
public class UserController {

    UserProducer userProducer;
    UserCheck userCheck;
    @Autowired
    UserController(UserProducer userProducer, UserCheck userCheck) {
        this.userProducer = userProducer;
        this.userCheck = userCheck;
    }


//    @GetMapping
//    public List<User> getUser(){
//       return userService.getUser();
//    }

    @PostMapping
    public ResponseEntity<String> addUser(@RequestBody User user){
        System.out.println("came to post endpoint");
        userCheck.addUser("123456789");
        boolean temp=userCheck.isblocked(user.getPhone_number());
        if(temp==true) {
            System.out.println("user is blocked");
          return new ResponseEntity<>("sorry it is a blocked user", HttpStatus.BAD_REQUEST);
        }
        userProducer.send_message(user);
        return new ResponseEntity<>("success", HttpStatus.OK);
    }

}
