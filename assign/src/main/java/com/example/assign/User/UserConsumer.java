package com.example.assign.User;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class UserConsumer {

    @KafkaListener(
            topics = "letmeeshoyou2k26",
            groupId = "firstassigngroup-v99"
    )
    public void consume(String json_type) {
        System.out.println("📥 Consumed: "+json_type);
//        user.printUser();
    }
}
