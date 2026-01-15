package com.example.assign.User;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

@Service
public class UserProducer  {
    @Autowired
    private KafkaTemplate<String,String> kafkaTemplate;
//
//    public void push_message(){
//        System.out.println("Push Message to kafka");
//    }
    public UserProducer(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }
    public void send_message(User user){
        System.out.println("Sending User: ");
        user.printUser();
//        kafkaTemplate.send()
        String no= user.getPhone_number();
        String name = user.getName();
        String message=user.getMessage();
        String json_object = String.format(
                "{\"phone_number\":\"%s\",\"name\":\"%s\",\"message\":\"%s\"}",
                no, name, message
        );
        kafkaTemplate.send("letmeeshoyou2k26", json_object)
                .whenComplete((result, ex) -> {
                    if (ex != null) {
                        System.out.println("Failed to send: " + ex.getMessage());
                    } else {
                        System.out.println("Sent to partition "
                                + result.getRecordMetadata().partition()
                                + " offset "
                                + result.getRecordMetadata().offset());
                    }
                });

    }
}
