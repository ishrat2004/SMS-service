package com.example.assign.User;

import org.springframework.boot.data.redis.autoconfigure.DataRedisProperties;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.stereotype.Service;

@Service
public class UserCheck {
    private StringRedisTemplate stringRedisTemplate;
    // do constructor injection
    public UserCheck(StringRedisTemplate stringRedisTemplate) {
        this.stringRedisTemplate = stringRedisTemplate;
    }
    public void addUser(String phone_number){
        stringRedisTemplate.opsForSet().add("phone_number","123456789");

    }
    public boolean isblocked(String phone_number){
        return stringRedisTemplate.opsForSet().isMember("phone_number",phone_number);
    }

}
