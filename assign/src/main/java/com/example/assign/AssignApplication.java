package com.example.assign;

import com.example.assign.User.User;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.kafka.annotation.EnableKafka;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@SpringBootApplication
@EnableKafka
public class AssignApplication {

	public static void main(String[] args) {
		SpringApplication.run(AssignApplication.class, args);
	}
//   @GetMapping
//    public String Hello(){
//        return "Hello World!";
//    }

}
